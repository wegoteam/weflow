package exec

import (
	"github.com/gookit/slog"
	"github.com/wegoteam/weflow/pkg/common/entity"
)

// ExecConvergenceNode 汇聚节点
type ExecConvergenceNode struct {
	NodeModel int8   `json:"nodeModel"` // 节点类型
	NodeName  string `json:"nodeName"`  // 节点名称
	NodeID    string `json:"nodeId"`    // 节点ID
	ParentID  string `json:"parentId"`  // 父节点ID
	//实例节点位置信息
	PreNodes    []string `json:"preNodes,omitempty"`    //上节点ID
	NextNodes   []string `json:"nextNodes,omitempty"`   //下节点ID
	LastNodes   []string `json:"lastNodes,omitempty"`   //分支节点尾节点ID
	Index       int      `json:"index,omitempty"`       // 下标
	BranchIndex int      `json:"branchIndex,omitempty"` // 分支下标
}

/**
实例化执行节点对象
*/
func NewConvergenceNode(node *entity.NodeModelBO) *ExecConvergenceNode {

	return &ExecConvergenceNode{
		NodeModel:   node.NodeModel,
		NodeName:    node.NodeName,
		NodeID:      node.NodeID,
		ParentID:    node.ParentID,
		PreNodes:    node.PreNodes,
		NextNodes:   node.NextNodes,
		LastNodes:   node.LastNodes,
		Index:       node.Index,
		BranchIndex: node.BranchIndex,
	}
}

/**
执行汇聚节点
生成实例节点任务
执行任务
下节点
*/
func (execConvergenceNode *ExecConvergenceNode) ExecCurrNodeModel(exec *entity.Execution) ExecResult {
	slog.Infof("ExecConvergenceNode 执行汇聚节点")
	processDefModel := exec.ProcessDefModel
	nextNodes := execConvergenceNode.ExecNextNodeModels(processDefModel.NodeModelMap)
	return ExecResult{
		NextNodes: nextNodes,
	}
}

func (execConvergenceNode *ExecConvergenceNode) ExecPreNodeModels(nodeModelMap map[string]entity.NodeModelBO) *[]entity.NodeModelBO {
	var preNodes = make([]entity.NodeModelBO, 0)
	if execConvergenceNode.PreNodes == nil {
		return &preNodes
	}
	for _, val := range execConvergenceNode.PreNodes {
		pre, ok := nodeModelMap[val]
		if !ok {
			slog.Infof("节点[%v]的上节点不存在", execConvergenceNode.NodeID)
		}
		preNodes = append(preNodes, pre)
	}
	return &preNodes
}

func (execConvergenceNode *ExecConvergenceNode) ExecNextNodeModels(nodeModelMap map[string]entity.NodeModelBO) *[]entity.NodeModelBO {
	var nextNodes = make([]entity.NodeModelBO, 0)
	if execConvergenceNode.NextNodes == nil {
		return &nextNodes
	}
	for _, val := range execConvergenceNode.NextNodes {
		next, ok := nodeModelMap[val]
		if !ok {
			slog.Infof("节点[%v]的下节点不存在", execConvergenceNode.NodeID)
		}
		nextNodes = append(nextNodes, next)
	}
	return &nextNodes
}
