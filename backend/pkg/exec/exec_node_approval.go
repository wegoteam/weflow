package exec

import (
	"fmt"
	"github.com/wegoteam/weflow/pkg/common/entity"
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
	fmt.Println("ExecApprovalNode 执行审批节点")
	processDefModel := exec.ProcessDefModel
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
			fmt.Println("上节点不存在")
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
			fmt.Println("下节点不存在")
		}
		nextNodes = append(nextNodes, next)
	}
	return &nextNodes
}
