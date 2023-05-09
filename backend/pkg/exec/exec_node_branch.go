package exec

import (
	"github.com/gookit/slog"
	"github.com/wegoteam/weflow/pkg/common/entity"
)

// ExecBranchNode 分支节点
type ExecBranchNode struct {
}

/**
执行分支节点
分支节点三个状态：1：分支节点未完成；2：分支节点完成且存在出口；3：分支节点完成无分支出口
生成实例节点任务
执行任务
下节点
*/
func (receiver *ExecBranchNode) ExecCurrNode(node *entity.NodeModelBO, exec *entity.Execution) ExecResult {
	slog.Infof("ExecBranchNode 执行分支节点")
	var branchNodes = make([]entity.NodeModelBO, 0)

	execNodeTaskMap := exec.ExecNodeTaskMap
	processDefModel := exec.ProcessDefModel
	nodeModelMap := processDefModel.NodeModelMap
	_, ok := execNodeTaskMap[node.NodeId]
	if !ok {
		slog.Infof("节点[%s]的分支节点未执行", node.NodeId)
		for _, childBranchs := range node.ChildrenIds {
			if childBranchs == nil {
				continue
			}
			bo, hasNode := nodeModelMap[childBranchs[0]]
			if !hasNode {
				slog.Infof("节点[%v]的分支节点不存在", node.NodeId)
			}
			branchNodes = append(branchNodes, bo)
		}

		return ExecResult{
			NextNodes: &branchNodes,
		}
	}

	//判断内存信息分支节点未执行

	return ExecResult{
		NextNodes: &branchNodes,
	}
}

func (receiver *ExecBranchNode) PreNodes(node *entity.NodeModelBO, nodeModelMap map[string]entity.NodeModelBO) *[]entity.NodeModelBO {
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

func (receiver *ExecBranchNode) NextNodes(node *entity.NodeModelBO, nodeModelMap map[string]entity.NodeModelBO) *[]entity.NodeModelBO {
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
