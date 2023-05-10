package exec

import (
	"github.com/gookit/slog"
	"github.com/wegoteam/weflow/pkg/common/entity"
	"github.com/wegoteam/wepkg/snowflake"
)

// ExecStartNode 开始节点
type ExecStartNode struct {
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
func NewStartNode(node *entity.NodeModelBO) *ExecStartNode {

	return &ExecStartNode{
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

func (execStartNode *ExecStartNode) ExecCurrNodeModel(exec *entity.Execution) ExecResult {
	slog.Infof("ExecStartNode 执行开始节点")
	nodeTaskId := snowflake.GetSnowflakeId()
	//生成执行节点任务
	var execNodeTask = &entity.ExecNodeTaskBO{
		NodeTaskID: nodeTaskId,
		NodeModel:  execStartNode.NodeModel,
		NodeID:     execStartNode.NodeID,
		Status:     2,
	}
	exec.ExecNodeTaskMap[execStartNode.NodeID] = *execNodeTask

	//生成实例节点任务
	var instNodeTask = entity.InstNodeTaskBO{}
	instNodeTasks := *exec.InstNodeTasks
	instNodeTasks = append(instNodeTasks, instNodeTask)

	processDefModel := exec.ProcessDefModel
	nextNodes := execStartNode.ExecNextNodeModels(processDefModel.NodeModelMap)
	return ExecResult{
		NextNodes: nextNodes,
	}
}

/**
执行开始节点
生成实例节点任务
执行任务
下节点
*/
func (execStartNode *ExecStartNode) ExecCurrNode(node *entity.NodeModelBO, exec *entity.Execution) ExecResult {
	slog.Infof("ExecStartNode 执行开始节点")
	nodeTaskId := snowflake.GetSnowflakeId()
	//生成执行节点任务
	var execNodeTask = &entity.ExecNodeTaskBO{
		NodeTaskID: nodeTaskId,
		NodeModel:  node.NodeModel,
		NodeID:     node.NodeID,
		Status:     2,
	}
	exec.ExecNodeTaskMap[node.NodeID] = *execNodeTask

	//生成实例节点任务
	var instNodeTask = entity.InstNodeTaskBO{}
	instNodeTasks := *exec.InstNodeTasks
	instNodeTasks = append(instNodeTasks, instNodeTask)

	processDefModel := exec.ProcessDefModel
	nextNodes := execStartNode.ExecNextNodes(node, processDefModel.NodeModelMap)
	return ExecResult{
		NextNodes: nextNodes,
	}
}
func (execStartNode *ExecStartNode) ExecPreNodeModels(nodeModelMap map[string]entity.NodeModelBO) *[]entity.NodeModelBO {
	var preNodes = make([]entity.NodeModelBO, 0)
	if execStartNode.PreNodes == nil {
		return &preNodes
	}
	for _, val := range execStartNode.PreNodes {
		pre, ok := nodeModelMap[val]
		if !ok {
			slog.Infof("节点[%v]的上节点不存在", execStartNode.NodeID)
		}
		preNodes = append(preNodes, pre)
	}
	return &preNodes
}

func (execStartNode *ExecStartNode) ExecPreNodes(node *entity.NodeModelBO, nodeModelMap map[string]entity.NodeModelBO) *[]entity.NodeModelBO {
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

func (execStartNode *ExecStartNode) ExecNextNodeModels(nodeModelMap map[string]entity.NodeModelBO) *[]entity.NodeModelBO {
	var nextNodes = make([]entity.NodeModelBO, 0)
	if execStartNode.NextNodes == nil {
		return &nextNodes
	}
	for _, val := range execStartNode.NextNodes {
		next, ok := nodeModelMap[val]
		if !ok {
			slog.Infof("节点[%v]的下节点不存在", execStartNode.NodeID)
		}
		nextNodes = append(nextNodes, next)
	}
	return &nextNodes
}

func (execStartNode *ExecStartNode) ExecNextNodes(node *entity.NodeModelBO, nodeModelMap map[string]entity.NodeModelBO) *[]entity.NodeModelBO {
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
