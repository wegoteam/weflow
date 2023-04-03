package parser

import (
	"encoding/json"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/wego2023/weflow/pkg/common/constant"
	"github.com/wego2023/weflow/pkg/common/entity"
	"github.com/wego2023/weflow/pkg/common/utils"
)

var FlowParserServiceImpl = new(FlowParser)

type FlowParser struct {
}

type FlowParserService interface {
	Parser(data string) []entity.NodeModelEntity
}

func (receiver *FlowParser) Test() {
	hlog.Errorf("parse yaml error - %v", "ces")
}

/*
*
将字符串解析为节点实体
*/
func (receiver *FlowParser) Parser(data string) []entity.NodeModelBO {

	var nodes = make([]entity.NodeModelEntity, 0)
	var datas = make([]entity.NodeModelBO, 0)
	if utils.IsStrBlank(data) {
		hlog.Info("节点json数据为空，无需解析")
		return datas
	}
	err := json.Unmarshal([]byte(data), &nodes)
	if err != nil {
		hlog.Warnf("解析流程图失败，错误信息：%s", err.Error())
	}

	//遍历解析节点
	var preIds = make([]string, 0)
	nodeLen := len(nodes)
	for nodeInd, node := range nodes {
		var nodeBO = ParserNode(&node)
		//设置分支节点下标
		nodeBO.Index = nodeInd
		datas = append(datas, *nodeBO)

		//设置上节点，下节点
		if nodeInd > 0 {
			preIds = append(preIds, node.NodeId)
			nodeBO.PreNodes = preIds
		}
		if nodeInd+1 < nodeLen {
			var nextIds = make([]string, 0)
			nextIds = append(nextIds, nodes[nodeInd+1].NodeId)
			nodeBO.NextNodes = nextIds
		}
		if node.NodeModel != constant.BranchNodeModel ||
			node.Children == nil {
			continue
		}
		//分支节点
		for branch, branchChilds := range node.Children {

			for ind, child := range branchChilds {
				var childNode = ParserNode(&child)
				childNode.Index = ind
				childNode.BranchIndex = branch
				datas = append(datas, *childNode)
			}
		}
	}
	hlog.Info("解析流程图成功")
	return datas
}

/**
解析节点
节点：上节点、下节点、节点下标、尾节点等基础信息
*/
func ParserNode(node *entity.NodeModelEntity) *entity.NodeModelBO {
	var bo = &entity.NodeModelBO{}
	err := utils.BeanCopy(bo, node)
	if err != nil {
		hlog.Errorf("节点属性转换失败%v\n", err)
	}

	return bo
}
