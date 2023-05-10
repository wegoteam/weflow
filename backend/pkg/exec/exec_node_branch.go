package exec

import (
	"github.com/gookit/slog"
	"github.com/wegoteam/weflow/pkg/common/entity"
)

// ExecBranchNode 分支节点
type ExecBranchNode struct {
	NodeModel     int8       `json:"nodeModel"`             // 节点类型
	NodeName      string     `json:"nodeName"`              // 节点名称
	NodeID        string     `json:"nodeId"`                // 节点ID
	ParentID      string     `json:"parentId"`              // 父节点ID
	BranchMode    int        `json:"branchMode"`            // 分支执行方式【单分支：1；多分支：2】默认多分支2
	DefaultBranch int        `json:"defaultBranch"`         // 单分支处理需要默认分支，在条件优先级无法处理时候执行默认分支，取值分支下标
	ChildrenIds   [][]string `json:"childrenIds,omitempty"` // 子节点ID
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
func NewBranchNode(node *entity.NodeModelBO) *ExecBranchNode {

	return &ExecBranchNode{
		NodeModel:     node.NodeModel,
		NodeName:      node.NodeName,
		NodeID:        node.NodeID,
		ParentID:      node.ParentID,
		BranchMode:    node.BranchMode,
		DefaultBranch: node.DefaultBranch,
		ChildrenIds:   node.ChildrenIDs,
		PreNodes:      node.PreNodes,
		NextNodes:     node.NextNodes,
		LastNodes:     node.LastNodes,
		Index:         node.Index,
		BranchIndex:   node.BranchIndex,
	}
}

func (execBranchNode *ExecBranchNode) ExecCurrNodeModel(exec *entity.Execution) ExecResult {
	slog.Infof("ExecBranchNode 执行分支节点")
	var branchNodes = make([]entity.NodeModelBO, 0)

	execNodeTaskMap := exec.ExecNodeTaskMap
	processDefModel := exec.ProcessDefModel
	nodeModelMap := processDefModel.NodeModelMap
	_, ok := execNodeTaskMap[execBranchNode.NodeID]
	if !ok {
		slog.Infof("节点[%s]的分支节点未执行", execBranchNode.NodeID)
		for _, childBranchs := range execBranchNode.ChildrenIds {
			if childBranchs == nil {
				continue
			}
			bo, hasNode := nodeModelMap[childBranchs[0]]
			if !hasNode {
				slog.Infof("节点[%v]的分支节点不存在", execBranchNode.NodeID)
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

/**
执行分支节点
分支节点三个状态：1：分支节点未完成；2：分支节点完成且存在出口；3：分支节点完成无分支出口
生成实例节点任务
执行任务
下节点
*/
func (execBranchNode *ExecBranchNode) ExecCurrNode(node *entity.NodeModelBO, exec *entity.Execution) ExecResult {
	slog.Infof("ExecBranchNode 执行分支节点")
	var branchNodes = make([]entity.NodeModelBO, 0)

	execNodeTaskMap := exec.ExecNodeTaskMap
	processDefModel := exec.ProcessDefModel
	nodeModelMap := processDefModel.NodeModelMap
	_, ok := execNodeTaskMap[node.NodeID]
	if !ok {
		slog.Infof("节点[%s]的分支节点未执行", node.NodeID)
		for _, childBranchs := range node.ChildrenIDs {
			if childBranchs == nil {
				continue
			}
			bo, hasNode := nodeModelMap[childBranchs[0]]
			if !hasNode {
				slog.Infof("节点[%v]的分支节点不存在", node.NodeID)
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
func (execBranchNode *ExecBranchNode) ExecPreNodeModels(nodeModelMap map[string]entity.NodeModelBO) *[]entity.NodeModelBO {
	var preNodes = make([]entity.NodeModelBO, 0)
	if execBranchNode.PreNodes == nil {
		return &preNodes
	}
	for _, val := range execBranchNode.PreNodes {
		pre, ok := nodeModelMap[val]
		if !ok {
			slog.Infof("节点[%v]的上节点不存在", execBranchNode.NodeID)
		}
		preNodes = append(preNodes, pre)
	}
	return &preNodes
}
func (execBranchNode *ExecBranchNode) ExecPreNodes(node *entity.NodeModelBO, nodeModelMap map[string]entity.NodeModelBO) *[]entity.NodeModelBO {
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

func (execBranchNode *ExecBranchNode) ExecNextNodeModels(nodeModelMap map[string]entity.NodeModelBO) *[]entity.NodeModelBO {
	var nextNodes = make([]entity.NodeModelBO, 0)
	if execBranchNode.NextNodes == nil {
		return &nextNodes
	}
	for _, val := range execBranchNode.NextNodes {
		next, ok := nodeModelMap[val]
		if !ok {
			slog.Infof("节点[%v]的下节点不存在", execBranchNode.NodeID)
		}
		nextNodes = append(nextNodes, next)
	}
	return &nextNodes
}

func (execBranchNode *ExecBranchNode) ExecNextNodes(node *entity.NodeModelBO, nodeModelMap map[string]entity.NodeModelBO) *[]entity.NodeModelBO {
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
