package parser

import (
	"encoding/json"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/wegoteam/weflow/pkg/common/constant"
	"github.com/wegoteam/weflow/pkg/common/entity"
	"github.com/wegoteam/weflow/pkg/common/utils"
)

var FlowParserServiceImpl = new(FlowParser)

type FlowParser struct {
}

type FlowParserService interface {
	Parser(data string) []entity.NodeModelEntity
}

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
	err := json.Unmarshal([]byte(data), &nodes)
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
			preIds = append(preIds, nodes[nodeInd-1].NodeId)
			nodeBO.PreNodes = preIds
		}
		if nodeInd+1 < nodeLen {
			var nextIds = make([]string, 0)
			nextIds = append(nextIds, nodes[nodeInd+1].NodeId)
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
	var bo = &entity.NodeModelBO{}
	err := utils.BeanCopy(bo, node)
	if err != nil {
		hlog.Errorf("节点属性转换失败%v\n", err)
	}
	fmt.Printf("解析节点%v\n", bo.NodeName)
	return bo
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
	nodeBO.ChildrenIds = branchIds
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
			branchChildIds[ind] = child.NodeId
			var childNode = parserNodeModel(&child)
			childNode.Index = ind
			childNode.BranchIndex = branch
			//设置上节点，下节点
			if ind > 0 {
				var preIds = make([]string, 0)
				preIds = append(preIds, branchChilds[ind-1].NodeId)
				childNode.PreNodes = preIds
			}
			if ind+1 < branchChildLen {
				var nextIds = make([]string, 0)
				nextIds = append(nextIds, branchChilds[ind+1].NodeId)
				childNode.NextNodes = nextIds
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
