package exec

import (
	"fmt"
	"github.com/wegoteam/weflow/pkg/common/entity"
	"github.com/wegoteam/wepkg/snowflake"
)

// ExecStartNode 开始节点
type ExecStartNode struct {
}

/**
执行开始节点
生成实例节点任务
执行任务
下节点
*/
func (receiver *ExecStartNode) ExecCurrNode(node *entity.NodeModelBO, exec *entity.Execution) ExecResult {
	fmt.Println("ExecStartNode 执行开始节点")
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
	var instNodeTask = entity.InstNodeTask{}
	instNodeTasks := *exec.InstNodeTasks
	instNodeTasks = append(instNodeTasks, instNodeTask)

	processDefModel := exec.ProcessDefModel
	nextNodes := receiver.NextNodes(node, processDefModel.NodeModelMap)
	return ExecResult{
		NextNodes: nextNodes,
	}
}

func (receiver *ExecStartNode) PreNodes(node *entity.NodeModelBO, nodeModelMap map[string]entity.NodeModelBO) *[]entity.NodeModelBO {
	var preNodes = make([]entity.NodeModelBO, 0)
	if node.PreNodes == nil {
		return &preNodes
	}
	for _, val := range node.PreNodes {
		pre, ok := nodeModelMap[val]
		if !ok {
			fmt.Println("上节点不存在")
		}
		preNodes = append(preNodes, pre)
	}
	return &preNodes
}

func (receiver *ExecStartNode) NextNodes(node *entity.NodeModelBO, nodeModelMap map[string]entity.NodeModelBO) *[]entity.NodeModelBO {
	var nextNodes = make([]entity.NodeModelBO, 0)
	if node.NextNodes == nil {
		return &nextNodes
	}
	for _, val := range node.NextNodes {
		next, ok := nodeModelMap[val]
		if !ok {
			fmt.Println("下节点不存在")
		}
		nextNodes = append(nextNodes, next)
	}
	return &nextNodes
}
