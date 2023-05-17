package parser

import (
	"context"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/redis/go-redis/v9"
	"github.com/wegoteam/weflow/pkg/common/constant"
	"github.com/wegoteam/weflow/pkg/common/entity"
	"github.com/wegoteam/weflow/pkg/common/utils"
	"github.com/wegoteam/weflow/pkg/model"
	"time"
)

/*
*
将字符串解析为节点实体
*/
func Parser(data string) *[]entity.NodeModelBO {

	var nodes = make([]entity.NodeModelEntity, 0)
	var datas = make([]entity.NodeModelBO, 0)
	if utils.IsStrBlank(data) {
		hlog.Info("节点json数据为空，无需解析")
		return &datas
	}
	err := sonic.Unmarshal([]byte(data), &nodes)
	if err != nil {
		hlog.Warnf("解析流程图失败，错误信息：%s", err.Error())
	}

	//遍历解析节点
	nodeLen := len(nodes)
	for nodeInd, node := range nodes {
		var nodeBO = parserNodeModel(&node)
		//设置分支节点下标
		nodeBO.Index = nodeInd
		//设置上节点，下节点
		if nodeInd > 0 {
			var preIds = make([]string, 0)
			preIds = append(preIds, nodes[nodeInd-1].NodeID)
			nodeBO.PreNodes = preIds
		}
		if nodeInd+1 < nodeLen {
			var nextIds = make([]string, 0)
			nextIds = append(nextIds, nodes[nodeInd+1].NodeID)
			nodeBO.NextNodes = nextIds
		}
		if node.NodeModel == constant.BranchNodeModel {
			//解析分支节点
			parserBranchNodeModel(nodeBO, node.Children, &datas)
			continue
		}
		datas = append(datas, *nodeBO)
	}
	hlog.Info("解析流程图成功")
	return &datas
}

/**
解析节点：开始节点、审批节点、知会节点、条件节点
节点：上节点、下节点、节点下标、尾节点等基础信息
*/
func parserNodeModel(node *entity.NodeModelEntity) *entity.NodeModelBO {
	//var bo = &entity.NodeModelBO{}
	//err := utils.BeanCopy(bo, node)
	//if err != nil {
	//	hlog.Errorf("节点属性转换失败%v\n", err)
	//}
	//return bo
	return &entity.NodeModelBO{
		NodeID:         node.NodeID,
		NodeName:       node.NodeName,
		NodeModel:      node.NodeModel,
		ParentID:       node.ParentID,
		ApproveType:    node.ApproveType,
		FormPer:        node.FormPer,
		NodeSetting:    node.NodeSetting,
		NodeHandler:    node.NodeHandler,
		NoneHandler:    node.NoneHandler,
		AppointHandler: node.AppointHandler,
		HandleMode:     node.HandleMode,
		FinishMode:     node.FinishMode,
		Level:          node.Level,
		ConditionGroup: node.ConditionGroup,
		ConditionExpr:  node.ConditionExpr,
		BranchMode:     node.BranchMode,
		DefaultBranch:  node.DefaultBranch,
	}
}

/**
解析节点：分支节点
节点：上节点、下节点、节点下标、尾节点等基础信息
*/
func parserBranchNodeModel(nodeBO *entity.NodeModelBO, childs [][]entity.NodeModelEntity, datas *[]entity.NodeModelBO) {
	if nodeBO.NodeModel != constant.BranchNodeModel {
		return
	}
	var branchIds = make([][]string, len(childs))
	var branchLastIds = make([]string, len(childs))
	nodeBO.ChildrenIDs = branchIds
	nodeBO.LastNodes = branchLastIds
	*datas = append(*datas, *nodeBO)
	if childs == nil {
		return
	}
	//分支节点
	for branch, branchChilds := range childs {
		if branchChilds == nil {
			continue
		}
		var branchChildIds = make([]string, len(branchChilds))
		var branchChildLen = len(branchChilds)
		for ind, child := range branchChilds {
			branchChildIds[ind] = child.NodeID
			var childNode = parserNodeModel(&child)
			childNode.Index = ind
			childNode.BranchIndex = branch
			//设置上节点，下节点，尾结点
			if ind > 0 {
				var preIds = make([]string, 0)
				preIds = append(preIds, branchChilds[ind-1].NodeID)
				childNode.PreNodes = preIds
			}
			if ind+1 < branchChildLen {
				var nextIds = make([]string, 0)
				nextIds = append(nextIds, branchChilds[ind+1].NodeID)
				childNode.NextNodes = nextIds
			}
			if ind == branchChildLen-1 {
				branchLastIds[branch] = child.NodeID
			}
			if child.NodeModel == constant.BranchNodeModel {
				parserBranchNodeModel(childNode, child.Children, datas)
				continue
			}
			*datas = append(*datas, *childNode)
		}
		branchIds[branch] = branchChildIds
	}

}

/**
在数据库中获取流程定义节点信息，部署到Redis中
*/
func buildProcessDefOnDB(processDefId string) *entity.ProcessDefModel {
	var processDefKey = constant.RedisProcessDefModel + processDefId
	ctx := context.Background()
	var processDefInfo = &model.ProcessDefInfo{}
	dbErr := MysqlDB.WithContext(ctx).Where(&model.ProcessDefInfo{ProcessDefID: processDefId}).First(processDefInfo).Error
	if dbErr != nil {
		hlog.Warnf("获取流程定义模型失败，错误信息：%s", dbErr.Error())
	}
	if processDefInfo == nil {
		return nil
	}
	//流程定义模型
	var processDefModel = &entity.ProcessDefModel{}
	nodes := Parser(processDefInfo.StructData)
	processDefModel.NodeModels = nodes
	var nodeModelMap = make(map[string]entity.NodeModelBO)
	_, err := RedisCliet.Pipelined(ctx, func(pipeliner redis.Pipeliner) error {
		for _, node := range *nodes {
			if node.NodeModel == constant.StartNodeModel {
				processDefModel.StartNodeId = node.NodeID
			}
			nodeModelMap[node.NodeID] = node
			nodeStr, _ := sonic.Marshal(&node)
			//设置key
			RedisCliet.HSet(ctx, processDefKey, node.NodeID, string(nodeStr))
		}
		return nil
	})
	if err != nil {
		hlog.Warnf("获取流程定义模型失败，错误信息：%s", err.Error())
	}
	//设置过期时间
	err = RedisCliet.Expire(ctx, processDefKey, time.Hour*72).Err()
	if err != nil {
		hlog.Warnf("获取流程定义模型失败，错误信息：%s", err.Error())
	}
	processDefModel.ProcessDefId = processDefId
	processDefModel.NodeModelMap = nodeModelMap
	return processDefModel
}

/**
从Redis中获取流程定义的节点信息
*/
func buildProcessDefOnRedis(processDefId string) *entity.ProcessDefModel {
	var processDefKey = constant.RedisProcessDefModel + processDefId
	ctx := context.Background()
	var nodeStrMap map[string]string
	nodeStrMap, err := RedisCliet.HGetAll(ctx, processDefKey).Result()
	if err != nil {
		hlog.Warnf("获取流程定义模型失败，错误信息：%s", err.Error())
	}
	var processDefModel = &entity.ProcessDefModel{}
	var nodeModelMap = make(map[string]entity.NodeModelBO)
	var nodes = make([]entity.NodeModelBO, 0)
	for key, val := range nodeStrMap {
		var node = &entity.NodeModelBO{}
		err := sonic.Unmarshal([]byte(val), node)
		if err != nil {
			hlog.Warnf("获取流程定义模型失败，错误信息：%s", err.Error())
		}
		if node.NodeModel == constant.StartNodeModel {
			processDefModel.StartNodeId = node.NodeID
		}
		nodeModelMap[key] = *node
		nodes = append(nodes, *node)
	}
	processDefModel.NodeModels = &nodes
	processDefModel.NodeModelMap = nodeModelMap
	processDefModel.ProcessDefId = processDefId
	return processDefModel
}
