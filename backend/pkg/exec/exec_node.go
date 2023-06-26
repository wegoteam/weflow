package exec

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/wegoteam/weflow/pkg/common/constant"
	"github.com/wegoteam/weflow/pkg/common/entity"
)

// IExecNode
// @Description: 执行节点接口
type IExecNode interface {
	// execCurrNodeModel 执行当前节点
	execCurrNodeModel(execution *Execution) ExecResult
	// execNextNodeModels 获取下一节点
	execNextNodeModels(nodeModelMap map[string]entity.NodeModelBO) *[]entity.NodeModelBO
	// execPreNodeModels 获取上一节点
	execPreNodeModels(nodeModelMap map[string]entity.NodeModelBO) *[]entity.NodeModelBO
}

// ExecResult
// @Description: 执行结果
type ExecResult struct {
	NextNodes *[]entity.NodeModelBO
}

// getExecNode
// @Description: 获取执行节点
// @param: node
// @return IExecNode
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

// execNode
// @Description: 执行流转节点
// @param: currNode
// @param: execution
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

// execNextNode
// @Description: 执行下节点
// @param: currNode
// @param: execution
func execNextNode(currNode *entity.NodeModelBO, execution *Execution) {
	iexec := getExecNode(currNode)
	if iexec == nil {
		return
	}
	processDefModel := execution.ProcessDefModel
	nextNodes := iexec.execNextNodeModels(processDefModel.NodeModelMap)
	if nextNodes == nil || len(*nextNodes) == 0 {
		return
	}
	for _, nextNode := range *nextNodes {
		execNode(&nextNode, execution)
	}
}

// execNextTask
// @Description: 执行流转串行节点任务
// @param: currNode
// @param: execution
func execNextTask(currNode *entity.NodeModelBO, userTaskExecution *UserTaskExecution) {
	addUserTasks := ExecNextUserTask(userTaskExecution, currNode.NodeHandler)
	execution := userTaskExecution.Execution
	userTasks := execution.UserTasks
	*userTasks = append(*userTasks, addUserTasks...)
}

// isParent
// @Description: 判断是否为父节点
// @param: parentId
// @return bool
func isParent(parentId string) bool {
	if parentId == "" || len(parentId) == 0 {
		return true
	}
	return false
}
