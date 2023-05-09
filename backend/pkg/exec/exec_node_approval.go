package exec

import (
	"github.com/gookit/slog"
	"github.com/wegoteam/weflow/pkg/common/entity"
	"github.com/wegoteam/wepkg/snowflake"
)

// ExecApprovalNode 审批节点
type ExecApprovalNode struct {
}

/**
执行审批节点
生成实例节点任务
执行任务
下节点
*/
func (receiver *ExecApprovalNode) ExecCurrNode(node *entity.NodeModelBO, exec *entity.Execution) ExecResult {
	processDefModel := exec.ProcessDefModel
	nodeTaskId := snowflake.GetSnowflakeId()

	//生成执行节点任务
	var execNodeTask = &entity.ExecNodeTaskBO{
		NodeTaskID: nodeTaskId,
		NodeModel:  node.NodeModel,
		NodeID:     node.NodeId,
		Status:     2,
	}
	exec.ExecNodeTaskMap[node.NodeId] = *execNodeTask

	//生成实例节点任务
	var instNodeTask = entity.InstNodeTaskBO{}
	instNodeTasks := *exec.InstNodeTasks
	instNodeTasks = append(instNodeTasks, instNodeTask)

	//生成用户任务

	//执行任务

	nextNodes := receiver.NextNodes(node, processDefModel.NodeModelMap)
	return ExecResult{
		NextNodes: nextNodes,
	}
}

func (receiver *ExecApprovalNode) PreNodes(node *entity.NodeModelBO, nodeModelMap map[string]entity.NodeModelBO) *[]entity.NodeModelBO {
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

func (receiver *ExecApprovalNode) NextNodes(node *entity.NodeModelBO, nodeModelMap map[string]entity.NodeModelBO) *[]entity.NodeModelBO {
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
