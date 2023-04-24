package exec

import (
	"fmt"
	"github.com/wegoteam/weflow/pkg/common/constant"
	"github.com/wegoteam/weflow/pkg/common/entity"
)

type IExecNode interface {
	//执行当前节点
	ExecCurrNode(node *entity.NodeModelBO, exec *entity.Execution) ExecResult
	//获取下一节点
	NextNodes(node *entity.NodeModelBO, nodeModelMap map[string]entity.NodeModelBO) *[]entity.NodeModelBO
	//获取上一节点
	PreNodes(node *entity.NodeModelBO, nodeModelMap map[string]entity.NodeModelBO) *[]entity.NodeModelBO
}

type ExecNode struct {
	Exec IExecNode
	Node *entity.NodeModelBO
}
type ExecResult struct {
	NextNodes *[]entity.NodeModelBO
}

func NewExecNode(node *entity.NodeModelBO, exec IExecNode) *ExecNode {

	return &ExecNode{
		Exec: exec,
		Node: node,
	}
}

func GetExecNode(node *entity.NodeModelBO) IExecNode {
	var exec IExecNode
	switch node.NodeModel {
	case constant.START_NODE_MODEL:
		exec = &ExecStartNode{}
	case constant.APPROVAL_NODE_MODEL:
		exec = &ExecApprovalNode{}
	case constant.NOTIFY_NODE_MODEL:
		exec = &ExecNotifyNode{}
	case constant.CUSTOM_NODE_MODEL:
		exec = &ExecCustomNode{}
	case constant.CONDITION_NODE_MODEL:
		exec = &ExecConditionNode{}
	case constant.BRANCH_NODE_MODEL:
		exec = &ExecBranchNode{}
	case constant.CONVERGENCE_NODE_MODEL:
		exec = &ExecConvergenceNode{}
	case constant.END_NODE_MODEL:
		exec = &ExecEndNode{}
	default:
		fmt.Println("未知节点类型")
	}

	return exec
}

/**
执行流转
*/
func Exec(currnode *entity.NodeModelBO, execution *entity.Execution) {
	iexec := GetExecNode(currnode)
	execNode := NewExecNode(currnode, iexec)
	handleNodes := execNode.Exec.ExecCurrNode(currnode, execution)

	nodes := handleNodes.NextNodes
	if nodes == nil || len(*nodes) == 0 {
		return
	}
	for _, next := range *nodes {
		Exec(&next, execution)
	}
}
