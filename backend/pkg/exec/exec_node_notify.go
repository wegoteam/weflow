package exec

import (
	"github.com/gookit/slog"
	"github.com/wegoteam/weflow/pkg/common/entity"
)

// ExecNotifyNode 知会节点
type ExecNotifyNode struct {
}

/**
执行知会节点
生成实例节点任务
执行任务
生成知会用户任务
下节点
*/
func (receiver *ExecNotifyNode) ExecCurrNode(node *entity.NodeModelBO, exec *entity.Execution) ExecResult {
	slog.Infof("ExecNotifyNode 执行知会节点")
	processDefModel := exec.ProcessDefModel
	nextNodes := receiver.NextNodes(node, processDefModel.NodeModelMap)
	return ExecResult{
		NextNodes: nextNodes,
	}
}

func (receiver *ExecNotifyNode) PreNodes(node *entity.NodeModelBO, nodeModelMap map[string]entity.NodeModelBO) *[]entity.NodeModelBO {
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

func (receiver *ExecNotifyNode) NextNodes(node *entity.NodeModelBO, nodeModelMap map[string]entity.NodeModelBO) *[]entity.NodeModelBO {
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
