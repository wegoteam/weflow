package exec

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/wegoteam/weflow/pkg/common/constant"
	"github.com/wegoteam/weflow/pkg/common/entity"
)

type IExecNode interface {
	//执行当前节点
	ExecCurrNode(node *entity.NodeModelBO, exec *entity.Execution) ExecResult
	//执行当前节点
	ExecCurrNodeModel(exec *entity.Execution) ExecResult
	//获取下一节点
	ExecNextNodes(node *entity.NodeModelBO, nodeModelMap map[string]entity.NodeModelBO) *[]entity.NodeModelBO
	//获取下一节点
	ExecNextNodeModels(nodeModelMap map[string]entity.NodeModelBO) *[]entity.NodeModelBO
	//获取上一节点
	ExecPreNodes(node *entity.NodeModelBO, nodeModelMap map[string]entity.NodeModelBO) *[]entity.NodeModelBO
	//获取上一节点
	ExecPreNodeModels(nodeModelMap map[string]entity.NodeModelBO) *[]entity.NodeModelBO
}

type ExecResult struct {
	NextNodes *[]entity.NodeModelBO
}

/**
获取执行节点
*/
func GetExecNode(node *entity.NodeModelBO) IExecNode {
	var exec IExecNode
	switch node.NodeModel {
	case constant.START_NODE_MODEL:
		exec = NewStartNode(node)
	case constant.APPROVAL_NODE_MODEL:
		exec = NewApprovalNode(node)
	case constant.TRANSACT_NODE_MODEL:
		exec = NewTransactNode(node)
	case constant.NOTIFY_NODE_MODEL:
		exec = NewNotifyNode(node)
	case constant.CUSTOM_NODE_MODEL:
		exec = NewCustomNode(node)
	case constant.CONDITION_NODE_MODEL:
		exec = NewConditionNode(node)
	case constant.BRANCH_NODE_MODEL:
		exec = NewBranchNode(node)
	case constant.CONVERGENCE_NODE_MODEL:
		exec = NewConvergenceNode(node)
	case constant.END_NODE_MODEL:
		exec = NewEndNode(node)
	default:
		hlog.Error("未知节点类型")
		panic("未知节点类型")
	}

	return exec
}

/**
执行流转
*/
func Exec(currnode *entity.NodeModelBO, execution *entity.Execution) {
	execNode(currnode, execution)
}

/**
执行节点
*/
func execNode(currnode *entity.NodeModelBO, execution *entity.Execution) {
	iexec := GetExecNode(currnode)
	handleNodes := iexec.ExecCurrNodeModel(execution)

	nodes := handleNodes.NextNodes
	if nodes == nil || len(*nodes) == 0 {
		return
	}
	for _, next := range *nodes {
		execNode(&next, execution)
	}
}
