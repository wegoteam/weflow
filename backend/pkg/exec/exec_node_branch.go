package exec

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/elliotchance/pie/v2"
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

// NewBranchNode
// @Description: 实例化执行节点对象
// @param: node
// @return *ExecBranchNode
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

// execCurrNodeModel
// @Description: 执行分支节点
//分支节点三个状态：1：分支节点未完成；2：分支节点完成且存在出口；3：分支节点完成无分支出口
//生成实例节点任务
//执行任务
//下节点
// @receiver execBranchNode
// @param: execution
// @return ExecResult
func (execBranchNode *ExecBranchNode) execCurrNodeModel(execution *Execution) ExecResult {
	hlog.Infof("实例任务[%s]的流程定义[%s]执行分支节点[%s]节点名称[%s]", execution.InstTaskID, execution.ProcessDefId, execBranchNode.NodeID, execBranchNode.NodeName)
	//判断当前的节点任务是否存在、或者是否在进行中
	//获取是否存在节点的执行任务
	_, ok := execution.ExecNodeTaskMap[execBranchNode.NodeID]
	if !ok {
		return bulidBrachNotStartResult(execution, execBranchNode)
	}
	//验证当前的分支节点是否完成
	finishFlag := getCurrBranchFinishFlag(execution, execBranchNode)
	hlog.Infof("实例任务[%s]的流程定义[%s]执行分支节点[%s]节点名称[%s]的节点完成标志[%v]；备注：1：分支节点未完成；2：分支节点完成且存在出口；3：分支节点完成无分支出口", execution.InstTaskID, execution.ProcessDefId, execBranchNode.NodeID, execBranchNode.NodeName, finishFlag)
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

// buildBranchFinishedHasOutResult
// @Description: 分支节点完成且存在出口
// @param: execution
// @param: execBranchNode
// @return ExecResult
func buildBranchFinishedHasOutResult(execution *Execution, execBranchNode *ExecBranchNode) ExecResult {
	processDefModel := execution.ProcessDefModel
	//修改当前执行节点任务状态为完成
	execNodeTask := execution.ExecNodeTaskMap[execBranchNode.NodeID]
	execNodeTask.Status = constant.InstanceNodeTaskStatusComplete
	execution.ExecNodeTaskMap[execBranchNode.NodeID] = execNodeTask
	//生成实例节点任务
	instNodeTasks := execution.InstNodeTasks
	var instNodeTask = entity.InstNodeTaskBO{
		ExecOpType: constant.OperationTypeUpdate,
		InstTaskID: execution.InstTaskID,
		NodeTaskID: execNodeTask.NodeTaskID,
		Status:     constant.InstanceNodeTaskStatusComplete,
		UpdateTime: execution.Now,
	}
	*instNodeTasks = append(*instNodeTasks, instNodeTask)
	//执行任务
	nextNodes := execBranchNode.execNextNodeModels(processDefModel.NodeModelMap)
	return ExecResult{
		NextNodes: nextNodes,
	}
}

// buildBranchFinishedNotOutResult
// @Description: 分支节点完成无分支出口
// @param: execution
// @param: execBranchNode
// @return ExecResult
func buildBranchFinishedNotOutResult(execution *Execution, execBranchNode *ExecBranchNode) ExecResult {
	processDefModel := execution.ProcessDefModel
	//修改当前执行节点任务状态为不通过
	execNodeTask := execution.ExecNodeTaskMap[execBranchNode.NodeID]
	execNodeTask.Status = constant.InstanceNodeTaskStatusNotPass
	execution.ExecNodeTaskMap[execBranchNode.NodeID] = execNodeTask
	//生成实例节点任务
	instNodeTasks := execution.InstNodeTasks
	var instNodeTask = entity.InstNodeTaskBO{
		ExecOpType: constant.OperationTypeUpdate,
		InstTaskID: execution.InstTaskID,
		NodeTaskID: execNodeTask.NodeTaskID,
		Status:     constant.InstanceNodeTaskStatusNotPass,
		UpdateTime: execution.Now,
	}
	*instNodeTasks = append(*instNodeTasks, instNodeTask)
	//添加分支节点不通过的操作日志
	var instTaskOpLog = entity.InstTaskOpLogBO{
		InstTaskID: execution.InstTaskID,
		NodeID:     execBranchNode.NodeID,
		NodeName:   execBranchNode.NodeName,
		CreateTime: execution.Now,
		UpdateTime: execution.Now,
		Type:       constant.InstTaskOpLogNode,
		Remark:     fmt.Sprintf("分支节点[%s]完成无分支出口，节点流转异常", execBranchNode.NodeName),
	}
	instTaskOpLogs := execution.InstTaskOpLogs
	*instTaskOpLogs = append(*instTaskOpLogs, instTaskOpLog)
	//判断下节点是否为父节点
	if isParent(execBranchNode.ParentID) {
		return ExecResult{}
	}
	nodeModelMap := processDefModel.NodeModelMap
	//判断节点的父节点是否是分支节点，节点是否在分支节点的最后节点上
	pNodeModel, ok := nodeModelMap[execBranchNode.ParentID]
	if !ok {
		hlog.Warnf("实例任务[%s]的流程定义[%s]执行分支节点[%s]节点名称[%s]的父节点不存在", execution.InstTaskID, execution.ProcessDefId, execBranchNode.NodeID, execBranchNode.NodeName)
		return ExecResult{}
	}
	if pNodeModel.NodeModel != constant.BranchNodeModel {
		hlog.Warnf("实例任务[%s]的流程定义[%s]执行分支节点[%s]节点名称[%s]的父节点错误，该节点的父节点不是分支节点", execution.InstTaskID, execution.ProcessDefId, execBranchNode.NodeID, execBranchNode.NodeName)
		return ExecResult{}
	}
	branchNodeModel := NewBranchNode(&pNodeModel)
	if branchNodeModel.LastNodes == nil {
		hlog.Infof("实例任务[%s]的流程定义[%s]执行分支节点[%s]节点名称[%s]的父节点错误，该分支节点的最后节点为空", execution.InstTaskID, execution.ProcessDefId, execBranchNode.NodeID, execBranchNode.NodeName)
		return ExecResult{}
	}
	var nextNodes = make([]entity.NodeModelBO, 0)
	if pie.Contains(branchNodeModel.LastNodes, execBranchNode.NodeID) {
		nextNodes = append(nextNodes, pNodeModel)
	}
	return ExecResult{
		NextNodes: &nextNodes,
	}
}

// getCurrBranchFinishFlag
// @Description: 验证当前的分支节点是否完成
//分支节点三个状态：1：分支节点未完成；2：分支节点完成且存在出口；3：分支节点完成无分支出口
// @param: execution
// @param: execBranchNode
// @return int8
func getCurrBranchFinishFlag(execution *Execution, execBranchNode *ExecBranchNode) int8 {
	execNodeTaskMap := execution.ExecNodeTaskMap
	processDefModel := execution.ProcessDefModel
	nodeModelMap := processDefModel.NodeModelMap
	if execNodeTaskMap == nil || len(execNodeTaskMap) == 0 {
		return constant.BranchNodeStatusNotComplete
	}
	//条件节点的分支
	if execBranchNode.ChildrenIds == nil || len(execBranchNode.ChildrenIds) == 0 {
		return constant.BranchNodeStatusNoBranch
	}
	//分支节点的出口
	out := len(execBranchNode.ChildrenIds)
	//嵌套节点完成条件：
	//全部出口：全部的条件、嵌套、正常节点验证通过
	//一个或多个出口：全部的条件、嵌套、正常节点验证通过
	//没有出口：全部的条件、嵌套、正常节点验证通过
	for _, childs := range execBranchNode.ChildrenIds {
		if childs == nil || len(childs) == 0 {
			out = out - 1
			continue
		}
		for _, childId := range childs {
			nodeModelBO, hasNode := nodeModelMap[childId]
			if !hasNode {
				out = out - 1
				hlog.Warnf("实例任务[%s]的流程定义[%s]的分支节点[%v]的子节点[%v]不存在", execution.InstTaskID, execution.ProcessDefId, execBranchNode.NodeID, childId)
				continue
			}
			execNodeTaskBO, existTask := execNodeTaskMap[childId]
			if !existTask {
				return constant.BranchNodeStatusNotComplete
			}
			//任务状态【1：未开始；2：处理中；3：完成；4：回退；5：终止；6：不通过】
			status := execNodeTaskBO.Status
			//节点任务未开始
			if status == constant.InstanceNodeTaskStatusNotStart {
				return constant.BranchNodeStatusNotComplete
			}
			switch nodeModelBO.NodeModel {
			case constant.ConditionNodeModel:
				//条件节点
				if status == constant.InstanceNodeTaskStatusNotPass {
					out = out - 1
					break
				}
			case constant.ApprovalNodeModel, constant.TransactNodeModel, constant.CustomNodeModel:
				//审批节点、办理节点、自定义节点
				if status == constant.InstanceNodeTaskStatusDoing {
					return constant.BranchNodeStatusNotComplete
				}
				if status == constant.InstanceNodeTaskStatusNotPass {
					out = out - 1
					break
				}
			case constant.NotifyNodeModel:
				//抄送节点
				continue
			case constant.BranchNodeModel:
				//分支节点
				if status == constant.InstanceNodeTaskStatusNotPass {
					out = out - 1
					break
				}
				if status == constant.InstanceNodeTaskStatusDoing {
					return constant.BranchNodeStatusNotComplete
				}
			case constant.ConvergenceNodeModel:
				continue
			default:
				hlog.Warnf("实例任务[%s]的流程定义[%s]的分支节点[%v]的子节点[%v]的节点类型[%v]不支持", execution.InstTaskID, execution.ProcessDefId, execBranchNode.NodeID, childId, nodeModelBO.NodeModel)
				return constant.BranchNodeStatusNoBranch
			}
		}
	}
	if out >= 1 {
		return constant.BranchNodeStatusComplete
	}
	return constant.BranchNodeStatusNoBranch
}

// bulidBrachNotStartResult
// @Description: 未开始，流转分支节点的子分支的头节点
// @param: execution
// @param: execBranchNode
// @return ExecResult
func bulidBrachNotStartResult(execution *Execution, execBranchNode *ExecBranchNode) ExecResult {
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
			hlog.Infof("节点[%v]的分支节点不存在", execBranchNode.NodeID)
		}
		branchNodes = append(branchNodes, bo)
	}
	return ExecResult{
		NextNodes: &branchNodes,
	}
}

// GetInstNodeTask
// @Description: 获取实例节点任务
// @receiver execBranchNode
// @param: instTaskID
// @param: nodeTaskID
// @param: now
// @return entity.InstNodeTaskBO
func (execBranchNode *ExecBranchNode) GetInstNodeTask(instTaskID, nodeTaskID string, now time.Time) entity.InstNodeTaskBO {
	//生成实例节点任务
	var instNodeTask = entity.InstNodeTaskBO{
		ExecOpType:    constant.OperationTypeAdd,
		InstTaskID:    instTaskID,
		NodeTaskID:    nodeTaskID,
		NodeID:        execBranchNode.NodeID,
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

// execPreNodeModels
// @Description: 获取上一节点
// @receiver execBranchNode
// @param: nodeModelMap
// @return *[]entity.NodeModelBO
func (execBranchNode *ExecBranchNode) execPreNodeModels(nodeModelMap map[string]entity.NodeModelBO) *[]entity.NodeModelBO {
	var preNodes = make([]entity.NodeModelBO, 0)
	if execBranchNode.PreNodes == nil {
		return &preNodes
	}
	for _, val := range execBranchNode.PreNodes {
		pre, ok := nodeModelMap[val]
		if !ok {
			hlog.Warnf("节点[%v]的上节点不存在", execBranchNode.NodeID)
			continue
		}
		preNodes = append(preNodes, pre)
	}
	return &preNodes
}

// execNextNodeModels
// @Description: 获取下一节点
// @receiver execBranchNode
// @param: nodeModelMap
// @return *[]entity.NodeModelBO
func (execBranchNode *ExecBranchNode) execNextNodeModels(nodeModelMap map[string]entity.NodeModelBO) *[]entity.NodeModelBO {
	var nextNodes = make([]entity.NodeModelBO, 0)
	//判断是否有下节点
	if execBranchNode.NextNodes != nil {
		for _, val := range execBranchNode.NextNodes {
			next, ok := nodeModelMap[val]
			if !ok {
				hlog.Warnf("节点[%s]的下节点不存在", execBranchNode.NodeID)
				continue
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
