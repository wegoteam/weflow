package exec

import (
	"fmt"
	"github.com/wegoteam/weflow/pkg/common/entity"
)

// ExecEndNode 结束节点
type ExecEndNode struct {
}

func (receiver *ExecEndNode) ExecCurrNode(node *entity.NodeModelBO, exec *entity.Execution) ExecResult {
	fmt.Println("ExecEndNode 执行结束节点")

	return ExecResult{
		NextNodes: &[]entity.NodeModelBO{},
	}
}

func (receiver *ExecEndNode) PreNodes(node *entity.NodeModelBO, nodeModelMap map[string]entity.NodeModelBO) *[]entity.NodeModelBO {
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

func (receiver *ExecEndNode) NextNodes(node *entity.NodeModelBO, nodeModelMap map[string]entity.NodeModelBO) *[]entity.NodeModelBO {
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
