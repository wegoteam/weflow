package exec

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/wegoteam/weflow/pkg/common/constant"
	"github.com/wegoteam/weflow/pkg/common/entity"
)

type IExecNode interface {
	// execCurrNodeModel 执行当前节点
	execCurrNodeModel(execution *Execution) ExecResult
	// execNextNodeModels 获取下一节点
	execNextNodeModels(nodeModelMap map[string]entity.NodeModelBO) *[]entity.NodeModelBO
	// execPreNodeModels 获取上一节点
	execPreNodeModels(nodeModelMap map[string]entity.NodeModelBO) *[]entity.NodeModelBO
}

type ExecResult struct {
	NextNodes *[]entity.NodeModelBO
}

/**
获取执行节点
*/
func getExecNode(node *entity.NodeModelBO) IExecNode {
	var exec IExecNode
	switch node.NodeModel {
	case constant.StartNodeModel:
		exec = NewStartNode(node)
	case constant.ApprovalNodeModel:
		exec = NewApprovalNode(node)
	case constant.TransactNodeModel:
		exec = NewTransactNode(node)
	case constant.NotifyNodeModel:
		exec = NewNotifyNode(node)
	case constant.CustomNodeModel:
		exec = NewCustomNode(node)
	case constant.ConditionNodeModel:
		exec = NewConditionNode(node)
	case constant.BranchNodeModel:
		exec = NewBranchNode(node)
	case constant.ConvergenceNodeModel:
		exec = NewConvergenceNode(node)
	case constant.EndNodeModel:
		exec = NewEndNode(node)
	default:
		hlog.Error("未知节点类型，节点模型数据为%+v", node)
	}
	return exec
}

/**
执行流转节点
*/
func execNode(currNode *entity.NodeModelBO, execution *Execution) {
	iexec := getExecNode(currNode)
	if iexec == nil {
		return
	}
	handleNodes := iexec.execCurrNodeModel(execution)

	nodes := handleNodes.NextNodes
	if nodes == nil || len(*nodes) == 0 {
		return
	}
	for _, next := range *nodes {
		execNode(&next, execution)
	}
}

/**
判断是否为父节点
*/
func isParent(parentId string) bool {

	if parentId == "" || len(parentId) == 0 {
		return true
	}

	return false
}
