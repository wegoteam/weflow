package exec

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/wegoteam/weflow/pkg/common/constant"
	"github.com/wegoteam/weflow/pkg/common/entity"
	"github.com/wegoteam/wepkg/id/snowflake"
	"time"
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

// NewStartNode
// @Description: 实例化执行节点对象
// @param: node
// @return *ExecStartNode
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

// execCurrNodeModel
// @Description: 执行开始节点
//生成实例节点任务
//执行任务
//下节点
// @receiver: execStartNode
// @param: execution
// @return ExecResult
func (execStartNode *ExecStartNode) execCurrNodeModel(execution *Execution) ExecResult {
	hlog.Infof("实例任务[%s]的流程定义[%s]执行开始节点[%s]节点名称[%s]生成节点任务", execution.InstTaskID, execution.ProcessDefId, execStartNode.NodeID, execStartNode.NodeName)
	nodeTaskId := snowflake.GetSnowflakeId()
	//生成执行节点任务
	var execNodeTask = &entity.ExecNodeTaskBO{
		NodeTaskID: nodeTaskId,
		NodeModel:  execStartNode.NodeModel,
		NodeID:     execStartNode.NodeID,
		Status:     constant.InstanceNodeTaskStatusComplete,
	}
	execution.ExecNodeTaskMap[execStartNode.NodeID] = *execNodeTask

	//生成实例节点任务
	instNodeTasks := execution.InstNodeTasks
	var instNodeTask = execStartNode.GetInstNodeTask(execution.InstTaskID, nodeTaskId, execution.Now)
	*instNodeTasks = append(*instNodeTasks, instNodeTask)

	processDefModel := execution.ProcessDefModel
	nextNodes := execStartNode.execNextNodeModels(processDefModel.NodeModelMap)
	return ExecResult{
		NextNodes: nextNodes,
	}
}

// GetInstNodeTask
// @Description: 获取实例节点任务
// @receiver execStartNode
// @param: instTaskID
// @param: nodeTaskID
// @param: now
// @return entity.InstNodeTaskBO
func (execStartNode *ExecStartNode) GetInstNodeTask(instTaskID, nodeTaskID string, now time.Time) entity.InstNodeTaskBO {
	//生成实例节点任务
	var instNodeTask = entity.InstNodeTaskBO{
		ExecOpType: constant.OperationTypeAdd,
		InstTaskID: instTaskID,
		NodeTaskID: nodeTaskID,
		NodeID:     execStartNode.NodeID,
		ParentID:   execStartNode.ParentID,
		NodeModel:  int32(execStartNode.NodeModel),
		NodeName:   execStartNode.NodeName,
		Status:     constant.InstanceNodeTaskStatusComplete,
		CreateTime: now,
		UpdateTime: now,
	}

	return instNodeTask
}

// execPreNodeModels
// @Description: 获取上节点
// @receiver: execStartNode
// @param: nodeModelMap
// @return *[]entity.NodeModelBO
func (execStartNode *ExecStartNode) execPreNodeModels(nodeModelMap map[string]entity.NodeModelBO) *[]entity.NodeModelBO {
	var preNodes = make([]entity.NodeModelBO, 0)
	if execStartNode.PreNodes == nil {
		return &preNodes
	}
	for _, val := range execStartNode.PreNodes {
		pre, ok := nodeModelMap[val]
		if !ok {
			hlog.Infof("节点[%v]的上节点不存在", execStartNode.NodeID)
			continue
		}
		preNodes = append(preNodes, pre)
	}
	return &preNodes
}

// execNextNodeModels
// @Description: 获取下节点
// @receiver: execStartNode
// @param: nodeModelMap
// @return *[]entity.NodeModelBO
func (execStartNode *ExecStartNode) execNextNodeModels(nodeModelMap map[string]entity.NodeModelBO) *[]entity.NodeModelBO {
	var nextNodes = make([]entity.NodeModelBO, 0)
	if execStartNode.NextNodes == nil {
		return &nextNodes
	}
	for _, val := range execStartNode.NextNodes {
		next, ok := nodeModelMap[val]
		if !ok {
			hlog.Infof("节点[%v]的下节点不存在", execStartNode.NodeID)
			continue
		}
		nextNodes = append(nextNodes, next)
	}
	return &nextNodes
}
