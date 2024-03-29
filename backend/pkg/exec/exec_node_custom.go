package exec

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/elliotchance/pie/v2"
	"github.com/wegoteam/weflow/pkg/common/constant"
	"github.com/wegoteam/weflow/pkg/common/entity"
	"github.com/wegoteam/wepkg/id/snowflake"
	"time"
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

// NewCustomNode
// @Description: 实例化执行节点对象
// @param: node
// @return *ExecCustomNode
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

// execCurrNodeModel
// @Description: 执行自定义节点
//生成实例节点任务
//执行任务
//下节点
// @receiver: execCustomNode
// @param: execution
// @return ExecResult
func (execCustomNode *ExecCustomNode) execCurrNodeModel(execution *Execution) ExecResult {
	_, ok := execution.ExecNodeTaskMap[execCustomNode.NodeID]
	if ok {
		hlog.Warnf("实例任务[%s]的流程定义[%s]执行自定义节点[%s]节点名称[%s]已经生成节点任务，该节点重复执行", execution.InstTaskID, execution.ProcessDefId, execCustomNode.NodeID, execCustomNode.NodeName)
		return ExecResult{}
	}
	hlog.Infof("实例任务[%s]的流程定义[%s]执行自定义节点[%s]节点名称[%s]生成节点任务", execution.InstTaskID, execution.ProcessDefId, execCustomNode.NodeID, execCustomNode.NodeName)
	nodeTaskId := snowflake.GetSnowflakeId()
	//生成执行节点任务
	var execNodeTask = &entity.ExecNodeTaskBO{
		NodeTaskID: nodeTaskId,
		NodeModel:  execCustomNode.NodeModel,
		NodeID:     execCustomNode.NodeID,
		Status:     constant.InstanceNodeTaskStatusDoing,
	}
	execution.ExecNodeTaskMap[execCustomNode.NodeID] = *execNodeTask

	//生成实例节点任务
	instNodeTasks := execution.InstNodeTasks
	var instNodeTask = execCustomNode.GetInstNodeTask(execution.InstTaskID, nodeTaskId, execution.Now)
	*instNodeTasks = append(*instNodeTasks, instNodeTask)

	processDefModel := execution.ProcessDefModel
	nextNodes := execCustomNode.execNextNodeModels(processDefModel.NodeModelMap)
	return ExecResult{
		NextNodes: nextNodes,
	}
}

// GetInstNodeTask
// @Description: 获取实例节点任务
// @receiver: execCustomNode
// @param: instTaskID
// @param: nodeTaskID
// @param: now
// @return entity.InstNodeTaskBO
func (execCustomNode *ExecCustomNode) GetInstNodeTask(instTaskID, nodeTaskID string, now time.Time) entity.InstNodeTaskBO {
	//生成实例节点任务
	var instNodeTask = entity.InstNodeTaskBO{
		ExecOpType: constant.OperationTypeAdd,
		InstTaskID: instTaskID,
		NodeTaskID: nodeTaskID,
		ParentID:   execCustomNode.ParentID,
		NodeModel:  int32(execCustomNode.NodeModel),
		NodeName:   execCustomNode.NodeName,
		Status:     constant.InstanceNodeTaskStatusDoing,
		CreateTime: now,
		UpdateTime: now,
	}

	return instNodeTask
}

// execPreNodeModels
// @Description: 执行上节点
// @receiver: execCustomNode
// @param: nodeModelMap
// @return *[]entity.NodeModelBO
func (execCustomNode *ExecCustomNode) execPreNodeModels(nodeModelMap map[string]entity.NodeModelBO) *[]entity.NodeModelBO {
	var preNodes = make([]entity.NodeModelBO, 0)
	if execCustomNode.PreNodes == nil {
		return &preNodes
	}
	for _, val := range execCustomNode.PreNodes {
		pre, ok := nodeModelMap[val]
		if !ok {
			hlog.Infof("节点[%v]的上节点不存在", execCustomNode.NodeID)
			continue
		}
		preNodes = append(preNodes, pre)
	}
	return &preNodes
}

// execNextNodeModels
// @Description: 执行下节点
// @receiver: execCustomNode
// @param: nodeModelMap
// @return *[]entity.NodeModelBO
func (execCustomNode *ExecCustomNode) execNextNodeModels(nodeModelMap map[string]entity.NodeModelBO) *[]entity.NodeModelBO {
	var nextNodes = make([]entity.NodeModelBO, 0)

	//判断是否有下节点
	if execCustomNode.NextNodes != nil {
		for _, val := range execCustomNode.NextNodes {
			next, ok := nodeModelMap[val]
			if !ok {
				hlog.Infof("节点[%s]的下节点不存在", execCustomNode.NodeID)
				continue
			}
			nextNodes = append(nextNodes, next)
		}
	}
	//判断下节点是否为父节点
	if isParent(execCustomNode.ParentID) {
		return &nextNodes
	}
	//判断节点的父节点是否是分支节点，节点是否在分支节点的最后节点上
	pNodeModel, ok := nodeModelMap[execCustomNode.ParentID]
	if !ok {
		hlog.Warnf("节点[%s]的父节点不存在", execCustomNode.NodeID)
		return &nextNodes
	}
	if pNodeModel.NodeModel != constant.BranchNodeModel {
		hlog.Warnf("节点[%s]的父节点[%s]错误，该节点的父节点不是分支节点", execCustomNode.NodeID, execCustomNode.ParentID)
		return &nextNodes
	}
	branchNodeModel := NewBranchNode(&pNodeModel)
	if branchNodeModel.LastNodes == nil {
		hlog.Warnf("节点[%s]的父节点[%s]错误，该分支节点的最后节点为空", execCustomNode.NodeID, execCustomNode.ParentID)
		return &nextNodes
	}
	if pie.Contains(branchNodeModel.LastNodes, execCustomNode.NodeID) {
		nextNodes = append(nextNodes, pNodeModel)
	}
	return &nextNodes
}
