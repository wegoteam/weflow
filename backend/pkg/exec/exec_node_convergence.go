package exec

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/elliotchance/pie/v2"
	"github.com/wegoteam/weflow/pkg/common/constant"
	"github.com/wegoteam/weflow/pkg/common/entity"
	"github.com/wegoteam/wepkg/id/snowflake"
	"time"
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
func (execConvergenceNode *ExecConvergenceNode) execCurrNodeModel(execution *Execution) ExecResult {
	hlog.Infof("实例任务[%s]的流程定义[%s]执行汇聚节点[%s]节点名称[%s]生成节点任务", execution.InstTaskID, execution.ProcessDefId, execConvergenceNode.NodeID, execConvergenceNode.NodeName)
	processDefModel := execution.ProcessDefModel
	nodeTaskId := snowflake.GetSnowflakeId()

	//生成执行节点任务
	var execNodeTask = &entity.ExecNodeTaskBO{
		NodeTaskID: nodeTaskId,
		NodeModel:  execConvergenceNode.NodeModel,
		NodeID:     execConvergenceNode.NodeID,
		Status:     constant.InstanceNodeTaskStatusComplete,
	}
	execution.ExecNodeTaskMap[execConvergenceNode.NodeID] = *execNodeTask

	//生成实例节点任务
	instNodeTasks := execution.InstNodeTasks
	var instNodeTask = execConvergenceNode.GetInstNodeTask(execution.InstTaskID, nodeTaskId, execution.Now)
	*instNodeTasks = append(*instNodeTasks, instNodeTask)

	nextNodes := execConvergenceNode.execNextNodeModels(processDefModel.NodeModelMap)
	return ExecResult{
		NextNodes: nextNodes,
	}
}

/**
获取实例节点任务
*/
func (execConvergenceNode *ExecConvergenceNode) GetInstNodeTask(instTaskID, nodeTaskID string, now time.Time) entity.InstNodeTaskBO {
	//生成实例节点任务
	var instNodeTask = entity.InstNodeTaskBO{
		ExecOpType: constant.OperationTypeAdd,
		InstTaskID: instTaskID,
		NodeTaskID: nodeTaskID,
		NodeID:     execConvergenceNode.NodeID,
		ParentID:   execConvergenceNode.ParentID,
		NodeModel:  int32(execConvergenceNode.NodeModel),
		NodeName:   execConvergenceNode.NodeName,
		Status:     constant.InstanceNodeTaskStatusComplete,
		CreateTime: now,
		UpdateTime: now,
	}

	return instNodeTask
}

func (execConvergenceNode *ExecConvergenceNode) execPreNodeModels(nodeModelMap map[string]entity.NodeModelBO) *[]entity.NodeModelBO {
	var preNodes = make([]entity.NodeModelBO, 0)
	if execConvergenceNode.PreNodes == nil {
		return &preNodes
	}
	for _, val := range execConvergenceNode.PreNodes {
		pre, ok := nodeModelMap[val]
		if !ok {
			hlog.Infof("节点[%v]的上节点不存在", execConvergenceNode.NodeID)
			continue
		}
		preNodes = append(preNodes, pre)
	}
	return &preNodes
}

func (execConvergenceNode *ExecConvergenceNode) execNextNodeModels(nodeModelMap map[string]entity.NodeModelBO) *[]entity.NodeModelBO {
	var nextNodes = make([]entity.NodeModelBO, 0)

	//判断是否有下节点
	if execConvergenceNode.NextNodes != nil {
		for _, val := range execConvergenceNode.NextNodes {
			next, ok := nodeModelMap[val]
			if !ok {
				hlog.Infof("节点[%s]的下节点不存在", execConvergenceNode.NodeID)
				continue
			}
			nextNodes = append(nextNodes, next)
		}
	}

	//判断下节点是否为父节点
	if isParent(execConvergenceNode.ParentID) {
		return &nextNodes
	}
	//判断节点的父节点是否是分支节点，节点是否在分支节点的最后节点上
	pNodeModel, ok := nodeModelMap[execConvergenceNode.ParentID]
	if !ok {
		hlog.Warnf("节点[%s]的父节点不存在", execConvergenceNode.NodeID)
		return &nextNodes
	}
	if pNodeModel.NodeModel != constant.BranchNodeModel {
		hlog.Warnf("节点[%s]的父节点[%s]错误，该节点的父节点不是分支节点", execConvergenceNode.NodeID, execConvergenceNode.ParentID)
		return &nextNodes
	}
	branchNodeModel := NewBranchNode(&pNodeModel)
	if branchNodeModel.LastNodes == nil {
		hlog.Warnf("节点[%s]的父节点[%s]错误，该分支节点的最后节点为空", execConvergenceNode.NodeID, execConvergenceNode.ParentID)
		return &nextNodes
	}

	if pie.Contains(branchNodeModel.LastNodes, execConvergenceNode.NodeID) {
		nextNodes = append(nextNodes, pNodeModel)
	}
	return &nextNodes
}
