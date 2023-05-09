package exec

import (
	"github.com/gookit/slog"
	"github.com/wegoteam/weflow/pkg/common/entity"
	"github.com/wegoteam/weflow/pkg/expr"
)

// ExecConditionNode 条件节点
type ExecConditionNode struct {
}

/**
执行条件节点，验证条件表达式
*/
func (receiver *ExecConditionNode) ExecCurrNode(node *entity.NodeModelBO, exec *entity.Execution) ExecResult {
	slog.Infof("ExecConditionNode 执行条件节点")

	//条件
	conditions := node.Conditions
	//参数
	paramMap := exec.InstTaskParamMap

	//执行条件
	flag := expr.ExecExpr(conditions, paramMap)
	if !flag {
		slog.Infof("节点[%v]的条件不成立", node.NodeId)
		return ExecResult{}
	}
	processDefModel := exec.ProcessDefModel
	nextNodes := receiver.NextNodes(node, processDefModel.NodeModelMap)
	return ExecResult{
		NextNodes: nextNodes,
	}
}

func (receiver *ExecConditionNode) PreNodes(node *entity.NodeModelBO, nodeModelMap map[string]entity.NodeModelBO) *[]entity.NodeModelBO {
	var preNodes = make([]entity.NodeModelBO, 0)
	if node.PreNodes == nil {
		return &preNodes
	}
	for _, val := range node.PreNodes {
		pre, ok := nodeModelMap[val]
		if !ok {
			slog.Infof("节点[%v]的上节点不存在", node.NodeId)
		}
		preNodes = append(preNodes, pre)
	}
	return &preNodes
}

func (receiver *ExecConditionNode) NextNodes(node *entity.NodeModelBO, nodeModelMap map[string]entity.NodeModelBO) *[]entity.NodeModelBO {
	var nextNodes = make([]entity.NodeModelBO, 0)
	if node.NextNodes == nil {
		return &nextNodes
	}
	for _, val := range node.NextNodes {
		next, ok := nodeModelMap[val]
		if !ok {
			slog.Infof("节点[%v]的下节点不存在", node.NodeId)
		}
		nextNodes = append(nextNodes, next)
	}
	return &nextNodes
}
