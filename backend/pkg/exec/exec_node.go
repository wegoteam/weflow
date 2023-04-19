package exec

import (
	"fmt"
	"github.com/wegoteam/weflow/pkg/common/constant"
	"github.com/wegoteam/weflow/pkg/common/entity"
)

type IExecNode interface {
	HandleNode(node *entity.NodeModelBO, exec *Execution) ExecResult
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

func Exec(currnode *entity.NodeModelBO, execution *Execution) {
	iexec := GetExecNode(currnode)
	execNode := NewExecNode(currnode, iexec)
	handleNodes := execNode.Exec.HandleNode(currnode, nil)

	nodes := handleNodes.NextNodes
	if nodes == nil || len(*nodes) == 0 {
		return
	}
	for _, next := range *nodes {
		Exec(&next, execution)
	}
}

// ExecStartNode 开始节点
type ExecStartNode struct {
}

// ExecApprovalNode 审批节点
type ExecApprovalNode struct {
}

// ExecNotifyNode 知会节点
type ExecNotifyNode struct {
}

// ExecCustomNode 自定义节点
type ExecCustomNode struct {
}

// ExecConditionNode 条件节点
type ExecConditionNode struct {
}

// ExecBranchNode 分支节点
type ExecBranchNode struct {
}

// ExecConvergenceNode 汇聚节点
type ExecConvergenceNode struct {
}

// ExecEndNode 结束节点
type ExecEndNode struct {
}

func (receiver *ExecStartNode) HandleNode(node *entity.NodeModelBO, exec *Execution) ExecResult {
	fmt.Println("ExecStartNode 执行开始节点")
	return ExecResult{
		NextNodes: &[]entity.NodeModelBO{},
	}
}

func (receiver *ExecApprovalNode) HandleNode(node *entity.NodeModelBO, exec *Execution) ExecResult {
	fmt.Println("ExecApprovalNode 执行审批节点")
	return ExecResult{
		NextNodes: &[]entity.NodeModelBO{},
	}
}

func (receiver *ExecNotifyNode) HandleNode(node *entity.NodeModelBO, exec *Execution) ExecResult {
	fmt.Println("ExecNotifyNode 执行知会节点")
	return ExecResult{
		NextNodes: &[]entity.NodeModelBO{},
	}
}

func (receiver *ExecCustomNode) HandleNode(node *entity.NodeModelBO, exec *Execution) ExecResult {
	fmt.Println("ExecCustomNode 执行自定义节点")
	return ExecResult{
		NextNodes: &[]entity.NodeModelBO{},
	}
}

func (receiver *ExecConditionNode) HandleNode(node *entity.NodeModelBO, exec *Execution) ExecResult {
	fmt.Println("ExecConditionNode 执行条件节点")
	return ExecResult{
		NextNodes: &[]entity.NodeModelBO{},
	}
}

func (receiver *ExecBranchNode) HandleNode(node *entity.NodeModelBO, exec *Execution) ExecResult {
	fmt.Println("ExecBranchNode 执行分支节点")
	return ExecResult{
		NextNodes: &[]entity.NodeModelBO{},
	}
}

func (receiver *ExecConvergenceNode) HandleNode(node *entity.NodeModelBO, exec *Execution) ExecResult {
	fmt.Println("ExecConvergenceNode 执行汇聚节点")
	return ExecResult{
		NextNodes: &[]entity.NodeModelBO{},
	}
}

func (receiver *ExecEndNode) HandleNode(node *entity.NodeModelBO, exec *Execution) ExecResult {
	fmt.Println("ExecEndNode 执行结束节点")

	return ExecResult{
		NextNodes: &[]entity.NodeModelBO{},
	}
}
