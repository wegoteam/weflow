package exec

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/elliotchance/pie/v2"
	"github.com/wegoteam/weflow/pkg/common/constant"
	"github.com/wegoteam/weflow/pkg/common/entity"
	"github.com/wegoteam/weflow/pkg/expr"
	"github.com/wegoteam/wepkg/snowflake"
	"time"
)

// ExecConditionNode 条件节点
type ExecConditionNode struct {
	NodeModel      int8   `json:"nodeModel"`      // 节点类型
	NodeName       string `json:"nodeName"`       // 节点名称
	NodeID         string `json:"nodeId"`         // 节点ID
	ParentID       string `json:"parentId"`       // 父节点ID
	Level          int    `json:"level"`          //优先级，分支执行方式为多分支处理方式无优先级应为0
	ConditionGroup string `json:"conditionGroup"` //条件组前端描述展示条件组
	ConditionExpr  string `json:"conditionExpr"`  //条件组解析后的表达式
	//实例节点位置信息
	PreNodes    []string `json:"preNodes,omitempty"`    //上节点ID
	NextNodes   []string `json:"nextNodes,omitempty"`   //下节点ID
	LastNodes   []string `json:"lastNodes,omitempty"`   //分支节点尾节点ID
	Index       int      `json:"index,omitempty"`       // 下标
	BranchIndex int      `json:"branchIndex,omitempty"` // 分支下标
}

/**
实例化执行条件节点对象
*/
func NewConditionNode(node *entity.NodeModelBO) *ExecConditionNode {

	return &ExecConditionNode{
		NodeModel:      node.NodeModel,
		NodeName:       node.NodeName,
		NodeID:         node.NodeID,
		ParentID:       node.ParentID,
		Level:          node.Level,
		ConditionGroup: node.ConditionGroup,
		ConditionExpr:  node.ConditionExpr,
		PreNodes:       node.PreNodes,
		NextNodes:      node.NextNodes,
		LastNodes:      node.LastNodes,
		Index:          node.Index,
		BranchIndex:    node.BranchIndex,
	}
}

func (execConditionNode *ExecConditionNode) ExecCurrNodeModel(execution *entity.Execution) ExecResult {
	_, ok := execution.ExecNodeTaskMap[execConditionNode.NodeID]
	if ok {
		hlog.Warnf("实例任务[%s]的流程定义[%s]执行条件节点[%s]节点名称[%s]已经生成节点任务，该节点重复执行", execution.InstTaskID, execution.ProcessDefId, execConditionNode.NodeID, execConditionNode.NodeName)
		return ExecResult{}
	}

	hlog.Infof("实例任务[%s]的流程定义[%s]执行条件节点[%s]节点名称[%s]生成节点任务", execution.InstTaskID, execution.ProcessDefId, execConditionNode.NodeID, execConditionNode.NodeName)
	//条件
	conditions := execConditionNode.ConditionExpr
	//参数
	paramMap := execution.InstTaskParamMap
	//执行条件
	flag := expr.ExecExpr(conditions, paramMap)
	hlog.Infof("实例任务[%v]的流程定义[%v]执行条件节点[%v]节点名称[%v]的表达式：%v", execution.InstTaskID, execution.ProcessDefId, execConditionNode.NodeID, execConditionNode.NodeName, conditions)
	hlog.Infof("实例任务[%v]的流程定义[%v]执行条件节点[%v]节点名称[%v]的条件参数：%v", execution.InstTaskID, execution.ProcessDefId, execConditionNode.NodeID, execConditionNode.NodeName, paramMap)
	if !flag {
		return buildNoPass(execution, execConditionNode)
	}
	return buildPass(execution, execConditionNode)
}

/**
获取通过的返回结果
*/
func buildPass(execution *entity.Execution, execConditionNode *ExecConditionNode) ExecResult {
	hlog.Infof("实例任务[%v]的流程定义[%v]执行条件节点[%v]节点名称[%v]的条件成立", execution.InstTaskID, execution.ProcessDefId, execConditionNode.NodeID, execConditionNode.NodeName)

	nodeTaskId := snowflake.GetSnowflakeId()
	//流程定义
	processDefModel := execution.ProcessDefModel
	//生成执行节点任务
	var execNodeTask = &entity.ExecNodeTaskBO{
		NodeTaskID: nodeTaskId,
		NodeModel:  execConditionNode.NodeModel,
		NodeID:     execConditionNode.NodeID,
		Status:     constant.InstanceNodeTaskStatusComplete,
	}
	execution.ExecNodeTaskMap[execConditionNode.NodeID] = *execNodeTask

	//生成实例节点任务
	instNodeTasks := execution.InstNodeTasks
	var instNodeTask = execConditionNode.GetInstNodeTask(execution.InstTaskID, nodeTaskId, execution.Now)
	instNodeTask.Status = constant.InstanceNodeTaskStatusComplete
	*instNodeTasks = append(*instNodeTasks, instNodeTask)

	nextNodes := execConditionNode.ExecNextNodeModels(processDefModel.NodeModelMap)
	return ExecResult{
		NextNodes: nextNodes,
	}
}

/**
获取不通过的返回结果
*/
func buildNoPass(execution *entity.Execution, execConditionNode *ExecConditionNode) ExecResult {
	hlog.Warnf("实例任务[%v]的流程定义[%v]执行条件节点[%v]节点名称[%v]的条件不成立", execution.InstTaskID, execution.ProcessDefId, execConditionNode.NodeID, execConditionNode.NodeName)

	nodeTaskId := snowflake.GetSnowflakeId()
	//流程定义
	processDefModel := execution.ProcessDefModel
	//生成执行节点任务
	var execNodeTask = &entity.ExecNodeTaskBO{
		NodeTaskID: nodeTaskId,
		NodeModel:  execConditionNode.NodeModel,
		NodeID:     execConditionNode.NodeID,
		Status:     constant.InstanceNodeTaskStatusNotPass,
	}
	execution.ExecNodeTaskMap[execConditionNode.NodeID] = *execNodeTask

	//生成实例节点任务
	instNodeTasks := execution.InstNodeTasks
	var instNodeTask = execConditionNode.GetInstNodeTask(execution.InstTaskID, nodeTaskId, execution.Now)
	instNodeTask.Status = constant.InstanceNodeTaskStatusNotPass
	*instNodeTasks = append(*instNodeTasks, instNodeTask)

	//条件不成立，验证是父节点的分支是否有出口
	if isParent(execConditionNode.ParentID) {
		return ExecResult{}
	}
	var nextNodes = make([]entity.NodeModelBO, 0)
	//返回父的分支节点，验证分支节点是否有出口
	//判断下节点是否为父节点
	//判断节点的父节点是否是分支节点，节点是否在分支节点的最后节点上
	pNodeModel, ok := processDefModel.NodeModelMap[execConditionNode.ParentID]
	if !ok {
		hlog.Warnf("节点[%s]的父节点不存在", execConditionNode.NodeID)
		return ExecResult{}
	}
	if pNodeModel.NodeModel != constant.BranchNodeModel {
		hlog.Warnf("节点[%s]的父节点[%s]错误，该节点的父节点不是分支节点", execConditionNode.NodeID, execConditionNode.ParentID)
		return ExecResult{}
	}
	branchNodeModel := NewBranchNode(&pNodeModel)
	if branchNodeModel.LastNodes == nil {
		hlog.Warnf("节点[%s]的父节点[%s]错误，该分支节点的最后节点为空", execConditionNode.NodeID, execConditionNode.ParentID)
		return ExecResult{}
	}
	if pie.Contains(branchNodeModel.LastNodes, execConditionNode.NodeID) {
		nextNodes = append(nextNodes, pNodeModel)
	}
	return ExecResult{
		NextNodes: &nextNodes,
	}
}

/**
获取实例节点任务
*/
func (execConditionNode *ExecConditionNode) GetInstNodeTask(instTaskID, nodeTaskID string, now time.Time) entity.InstNodeTaskBO {
	//生成实例节点任务
	var instNodeTask = entity.InstNodeTaskBO{
		ExecOpType:     constant.OperationTypeAdd,
		InstTaskID:     instTaskID,
		NodeTaskID:     nodeTaskID,
		ParentID:       execConditionNode.ParentID,
		NodeModel:      int32(execConditionNode.NodeModel),
		NodeName:       execConditionNode.NodeName,
		BranchLevel:    int32(execConditionNode.Level),
		ConditionExpr:  execConditionNode.ConditionExpr,
		ConditionGroup: execConditionNode.ConditionGroup,
		Status:         constant.InstanceNodeTaskStatusDoing,
		CreateTime:     now,
		UpdateTime:     now,
	}
	return instNodeTask
}

func (execConditionNode *ExecConditionNode) ExecPreNodeModels(nodeModelMap map[string]entity.NodeModelBO) *[]entity.NodeModelBO {
	var preNodes = make([]entity.NodeModelBO, 0)
	if execConditionNode.PreNodes == nil {
		return &preNodes
	}
	for _, val := range execConditionNode.PreNodes {
		pre, ok := nodeModelMap[val]
		if !ok {
			hlog.Infof("节点[%v]的上节点不存在", execConditionNode.NodeID)
		}
		preNodes = append(preNodes, pre)
	}
	return &preNodes
}

func (execConditionNode *ExecConditionNode) ExecNextNodeModels(nodeModelMap map[string]entity.NodeModelBO) *[]entity.NodeModelBO {
	var nextNodes = make([]entity.NodeModelBO, 0)
	//判断是否有下节点
	if execConditionNode.NextNodes != nil {
		for _, val := range execConditionNode.NextNodes {
			next, ok := nodeModelMap[val]
			if !ok {
				hlog.Infof("节点[%s]的下节点不存在", execConditionNode.NodeID)
			}
			nextNodes = append(nextNodes, next)
		}
	}
	//判断下节点是否为父节点
	if isParent(execConditionNode.ParentID) {
		return &nextNodes
	}
	//判断节点的父节点是否是分支节点，节点是否在分支节点的最后节点上
	pNodeModel, ok := nodeModelMap[execConditionNode.ParentID]
	if !ok {
		hlog.Warnf("节点[%s]的父节点不存在", execConditionNode.NodeID)
		return &nextNodes
	}
	if pNodeModel.NodeModel != constant.BranchNodeModel {
		hlog.Warnf("节点[%s]的父节点[%s]错误，该节点的父节点不是分支节点", execConditionNode.NodeID, execConditionNode.ParentID)
		return &nextNodes
	}
	branchNodeModel := NewBranchNode(&pNodeModel)
	if branchNodeModel.LastNodes == nil {
		hlog.Warnf("节点[%s]的父节点[%s]错误，该分支节点的最后节点为空", execConditionNode.NodeID, execConditionNode.ParentID)
		return &nextNodes
	}
	if pie.Contains(branchNodeModel.LastNodes, execConditionNode.NodeID) {
		nextNodes = append(nextNodes, pNodeModel)
	}
	return &nextNodes
}
