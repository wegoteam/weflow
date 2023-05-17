package exec

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/wegoteam/weflow/pkg/common/constant"
	"github.com/wegoteam/weflow/pkg/common/entity"
	"github.com/wegoteam/wepkg/snowflake"
	"time"
)

// ExecEndNode 结束节点
type ExecEndNode struct {
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
func NewEndNode(node *entity.NodeModelBO) *ExecEndNode {

	return &ExecEndNode{
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
执行结束节点
生成实例节点任务
执行任务
下节点
*/
func (execEndNode *ExecEndNode) ExecCurrNodeModel(execution *entity.Execution) ExecResult {
	hlog.Infof("实例任务[%s]的流程定义[%s]执行结束节点[%s]节点名称[%s]生成节点任务", execution.InstTaskID, execution.ProcessDefId, execEndNode.NodeID, execEndNode.NodeName)
	nodeTaskId := snowflake.GetSnowflakeId()

	//生成执行节点任务
	var execNodeTask = &entity.ExecNodeTaskBO{
		NodeTaskID: nodeTaskId,
		NodeModel:  execEndNode.NodeModel,
		NodeID:     execEndNode.NodeID,
		Status:     constant.InstanceNodeTaskStatusComplete,
	}
	execution.ExecNodeTaskMap[execEndNode.NodeID] = *execNodeTask

	//生成实例节点任务
	instNodeTasks := execution.InstNodeTasks
	var instNodeTask = execEndNode.GetInstNodeTask(execution.InstTaskID, nodeTaskId, execution.Now)
	*instNodeTasks = append(*instNodeTasks, instNodeTask)

	//实例任务状态完成
	execution.InstTaskStatus = constant.InstanceTaskStatusComplete
	return ExecResult{
		NextNodes: &[]entity.NodeModelBO{},
	}
}

/**
获取实例节点任务
*/
func (execEndNode *ExecEndNode) GetInstNodeTask(instTaskID, nodeTaskID string, now time.Time) entity.InstNodeTaskBO {
	//生成实例节点任务
	var instNodeTask = entity.InstNodeTaskBO{
		InstTaskID: instTaskID,
		NodeTaskID: nodeTaskID,
		ParentID:   execEndNode.ParentID,
		NodeModel:  int32(execEndNode.NodeModel),
		NodeName:   execEndNode.NodeName,
		Status:     constant.InstanceNodeTaskStatusComplete,
		CreateTime: now,
		UpdateTime: now,
	}

	return instNodeTask
}

func (execEndNode *ExecEndNode) ExecPreNodeModels(nodeModelMap map[string]entity.NodeModelBO) *[]entity.NodeModelBO {
	var preNodes = make([]entity.NodeModelBO, 0)
	if execEndNode.PreNodes == nil {
		return &preNodes
	}
	for _, val := range execEndNode.PreNodes {
		pre, ok := nodeModelMap[val]
		if !ok {
			hlog.Infof("节点[%v]的上节点不存在", execEndNode.NodeID)
		}
		preNodes = append(preNodes, pre)
	}
	return &preNodes
}

func (execEndNode *ExecEndNode) ExecNextNodeModels(nodeModelMap map[string]entity.NodeModelBO) *[]entity.NodeModelBO {
	var nextNodes = make([]entity.NodeModelBO, 0)
	if execEndNode.NextNodes == nil {
		return &nextNodes
	}
	for _, val := range execEndNode.NextNodes {
		next, ok := nodeModelMap[val]
		if !ok {
			hlog.Infof("节点[%v]的下节点不存在", execEndNode.NodeID)
		}
		nextNodes = append(nextNodes, next)
	}
	return &nextNodes
}
