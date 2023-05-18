package exec

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/elliotchance/pie/v2"
	"github.com/gookit/slog"
	"github.com/wegoteam/weflow/pkg/common/constant"
	"github.com/wegoteam/weflow/pkg/common/entity"
	"github.com/wegoteam/wepkg/snowflake"
	"time"
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

/**
执行分支节点
分支节点三个状态：1：分支节点未完成；2：分支节点完成且存在出口；3：分支节点完成无分支出口
生成实例节点任务
执行任务
下节点
*/
func (execBranchNode *ExecBranchNode) ExecCurrNodeModel(execution *entity.Execution) ExecResult {
	hlog.Infof("实例任务[%s]的流程定义[%s]执行分支节点[%s]节点名称[%s]", execution.InstTaskID, execution.ProcessDefId, execBranchNode.NodeID, execBranchNode.NodeName)

	//判断当前的节点任务是否存在、或者是否在进行中
	//获取是否存在节点的执行任务
	_, ok := execution.ExecNodeTaskMap[execBranchNode.NodeID]
	if !ok {
		return bulidBrachNotStartResult(execution, execBranchNode)
	}
	//验证当前的分支节点是否完成
	finishFlag := getCurrBranchFinishFlag(execution, execBranchNode)
	//分支节点未完成
	//分支节点完成无分支出口，分支节点完成状态为不通过
	switch finishFlag {
	case constant.BranchNodeStatusNotComplete:
		//分支节点未完成
		return ExecResult{}
	case constant.BranchNodeStatusComplete:
		//分支节点完成且存在出口
		return buildBranchFinishedHasOutResult(execution, execBranchNode)
	case constant.BranchNodeStatusNoBranch:
		//分支节点完成无分支出口
		return buildBranchFinishedNotOutResult(execution, execBranchNode)
	default:
		return ExecResult{}
	}
}

/**
分支节点完成且存在出口
*/
func buildBranchFinishedHasOutResult(execution *entity.Execution, execBranchNode *ExecBranchNode) ExecResult {
	processDefModel := execution.ProcessDefModel
	//执行任务
	nextNodes := execBranchNode.ExecNextNodeModels(processDefModel.NodeModelMap)
	return ExecResult{
		NextNodes: nextNodes,
	}
}

/**
分支节点完成无分支出口
*/
func buildBranchFinishedNotOutResult(execution *entity.Execution, execBranchNode *ExecBranchNode) ExecResult {

	return ExecResult{}
}

/**
验证当前的分支节点是否完成
分支节点三个状态：1：分支节点未完成；2：分支节点完成且存在出口；3：分支节点完成无分支出口
*/
func getCurrBranchFinishFlag(execution *entity.Execution, execBranchNode *ExecBranchNode) int8 {

	return constant.BranchNodeStatusComplete
}

/**
未开始，流转分支节点的子分支的头节点
*/
func bulidBrachNotStartResult(execution *entity.Execution, execBranchNode *ExecBranchNode) ExecResult {
	hlog.Infof("实例任务[%s]的流程定义[%s]的分支节点[%s]节点名称[%s]未执行，生成新的实例节点任务", execution.InstTaskID, execution.ProcessDefId, execBranchNode.NodeID, execBranchNode.NodeName)
	var branchNodes = make([]entity.NodeModelBO, 0)
	nodeTaskId := snowflake.GetSnowflakeId()
	processDefModel := execution.ProcessDefModel
	nodeModelMap := processDefModel.NodeModelMap

	//生成执行节点任务
	var execNodeTask = &entity.ExecNodeTaskBO{
		NodeTaskID: nodeTaskId,
		NodeModel:  execBranchNode.NodeModel,
		NodeID:     execBranchNode.NodeID,
		Status:     constant.InstanceNodeTaskStatusDoing,
	}
	execution.ExecNodeTaskMap[execBranchNode.NodeID] = *execNodeTask

	//生成实例节点任务
	instNodeTasks := execution.InstNodeTasks
	var instNodeTask = execBranchNode.GetInstNodeTask(execution.InstTaskID, nodeTaskId, execution.Now)
	*instNodeTasks = append(*instNodeTasks, instNodeTask)

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

/**
获取实例节点任务
*/
func (execBranchNode *ExecBranchNode) GetInstNodeTask(instTaskID, nodeTaskID string, now time.Time) entity.InstNodeTaskBO {
	//生成实例节点任务
	var instNodeTask = entity.InstNodeTaskBO{
		ExecOpType:    constant.OperationTypeAdd,
		InstTaskID:    instTaskID,
		NodeTaskID:    nodeTaskID,
		ParentID:      execBranchNode.ParentID,
		NodeModel:     int32(execBranchNode.NodeModel),
		NodeName:      execBranchNode.NodeName,
		BranchMode:    int32(execBranchNode.BranchMode),
		DefaultBranch: int32(execBranchNode.DefaultBranch),
		Status:        constant.InstanceNodeTaskStatusDoing,
		CreateTime:    now,
		UpdateTime:    now,
	}

	return instNodeTask
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

func (execBranchNode *ExecBranchNode) ExecNextNodeModels(nodeModelMap map[string]entity.NodeModelBO) *[]entity.NodeModelBO {
	var nextNodes = make([]entity.NodeModelBO, 0)

	//判断是否有下节点
	if execBranchNode.NextNodes != nil {
		for _, val := range execBranchNode.NextNodes {
			next, ok := nodeModelMap[val]
			if !ok {
				hlog.Infof("节点[%s]的下节点不存在", execBranchNode.NodeID)
			}
			nextNodes = append(nextNodes, next)
		}
	}

	//判断下节点是否为父节点
	if isParent(execBranchNode.ParentID) {
		return &nextNodes
	}
	//判断节点的父节点是否是分支节点，节点是否在分支节点的最后节点上
	pNodeModel, ok := nodeModelMap[execBranchNode.ParentID]
	if !ok {
		hlog.Warnf("节点[%s]的父节点不存在", execBranchNode.NodeID)
		return &nextNodes
	}
	if pNodeModel.NodeModel != constant.BranchNodeModel {
		hlog.Warnf("节点[%s]的父节点[%s]错误，该节点的父节点不是分支节点", execBranchNode.NodeID, execBranchNode.ParentID)
		return &nextNodes
	}
	branchNodeModel := NewBranchNode(&pNodeModel)
	if branchNodeModel.LastNodes == nil {
		hlog.Warnf("节点[%s]的父节点[%s]错误，该分支节点的最后节点为空", execBranchNode.NodeID, execBranchNode.ParentID)
		return &nextNodes
	}

	if pie.Contains(branchNodeModel.LastNodes, execBranchNode.NodeID) {
		nextNodes = append(nextNodes, pNodeModel)
	}
	return &nextNodes
}
