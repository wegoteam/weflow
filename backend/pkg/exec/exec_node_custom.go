package exec

import (
	"github.com/gookit/slog"
	"github.com/wegoteam/weflow/pkg/common/entity"
)

//ExecCustomNode 自定义节点
type ExecCustomNode struct {
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
func NewCustomNode(node *entity.NodeModelBO) *ExecCustomNode {

	return &ExecCustomNode{
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
func (execCustomNode *ExecCustomNode) ExecCurrNodeModel(exec *entity.Execution) ExecResult {
	slog.Infof("ExecCustomNode 执行自定义节点")
	processDefModel := exec.ProcessDefModel
	nextNodes := execCustomNode.ExecNextNodeModels(processDefModel.NodeModelMap)
	return ExecResult{
		NextNodes: nextNodes,
	}
}

/**
执行自定义节点
生成实例节点任务
执行任务
下节点
*/
func (execCustomNode *ExecCustomNode) ExecCurrNode(node *entity.NodeModelBO, exec *entity.Execution) ExecResult {
	slog.Infof("ExecCustomNode 执行自定义节点")
	processDefModel := exec.ProcessDefModel
	nextNodes := execCustomNode.ExecNextNodes(node, processDefModel.NodeModelMap)
	return ExecResult{
		NextNodes: nextNodes,
	}
}
func (execCustomNode *ExecCustomNode) ExecPreNodeModels(nodeModelMap map[string]entity.NodeModelBO) *[]entity.NodeModelBO {
	var preNodes = make([]entity.NodeModelBO, 0)
	if execCustomNode.PreNodes == nil {
		return &preNodes
	}
	for _, val := range execCustomNode.PreNodes {
		pre, ok := nodeModelMap[val]
		if !ok {
			slog.Infof("节点[%v]的上节点不存在", execCustomNode.NodeID)
		}
		preNodes = append(preNodes, pre)
	}
	return &preNodes
}

func (execCustomNode *ExecCustomNode) ExecPreNodes(node *entity.NodeModelBO, nodeModelMap map[string]entity.NodeModelBO) *[]entity.NodeModelBO {
	var preNodes = make([]entity.NodeModelBO, 0)
	if node.PreNodes == nil {
		return &preNodes
	}
	for _, val := range node.PreNodes {
		pre, ok := nodeModelMap[val]
		if !ok {
			slog.Infof("节点[%v]的上节点不存在", node.NodeID)
		}
		preNodes = append(preNodes, pre)
	}
	return &preNodes
}

func (execCustomNode *ExecCustomNode) ExecNextNodeModels(nodeModelMap map[string]entity.NodeModelBO) *[]entity.NodeModelBO {
	var nextNodes = make([]entity.NodeModelBO, 0)
	if execCustomNode.NextNodes == nil {
		return &nextNodes
	}
	for _, val := range execCustomNode.NextNodes {
		next, ok := nodeModelMap[val]
		if !ok {
			slog.Infof("节点[%v]的下节点不存在", execCustomNode.NodeID)
		}
		nextNodes = append(nextNodes, next)
	}
	return &nextNodes
}

func (execCustomNode *ExecCustomNode) ExecNextNodes(node *entity.NodeModelBO, nodeModelMap map[string]entity.NodeModelBO) *[]entity.NodeModelBO {
	var nextNodes = make([]entity.NodeModelBO, 0)
	if node.NextNodes == nil {
		return &nextNodes
	}
	for _, val := range node.NextNodes {
		next, ok := nodeModelMap[val]
		if !ok {
			slog.Infof("节点[%v]的下节点不存在", node.NodeID)
		}
		nextNodes = append(nextNodes, next)
	}
	return &nextNodes
}
