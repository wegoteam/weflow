package exec

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/gookit/slog"
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
	hlog.Infof("实例任务[%s]的流程定义[%s]执行条件节点[%s]生成节点任务", execution.InstTaskID, execution.ProcessDefId, execConditionNode.NodeID)

	nodeTaskId := snowflake.GetSnowflakeId()

	//生成执行节点任务
	var execNodeTask = &entity.ExecNodeTaskBO{
		NodeTaskID: nodeTaskId,
		NodeModel:  execConditionNode.NodeModel,
		NodeID:     execConditionNode.NodeID,
		Status:     1,
	}
	execution.ExecNodeTaskMap[execConditionNode.NodeID] = *execNodeTask

	//生成实例节点任务
	instNodeTasks := execution.InstNodeTasks
	var instNodeTask = execConditionNode.GetInstNodeTask(execution.InstTaskID, nodeTaskId, execution.Now)
	*instNodeTasks = append(*instNodeTasks, instNodeTask)

	//条件
	conditions := execConditionNode.ConditionExpr
	//参数
	paramMap := execution.InstTaskParamMap

	//执行条件
	flag := expr.ExecExpr(conditions, paramMap)
	if !flag {
		slog.Infof("节点[%v]的条件不成立", execConditionNode.NodeID)
		return ExecResult{}
	}
	processDefModel := execution.ProcessDefModel
	nextNodes := execConditionNode.ExecNextNodeModels(processDefModel.NodeModelMap)
	return ExecResult{
		NextNodes: nextNodes,
	}
}

/**
获取实例节点任务
*/
func (execConditionNode *ExecConditionNode) GetInstNodeTask(instTaskID, nodeTaskID string, now time.Time) entity.InstNodeTaskBO {
	//生成实例节点任务
	var instNodeTask = entity.InstNodeTaskBO{
		InstTaskID:     instTaskID,
		NodeTaskID:     nodeTaskID,
		ParentID:       execConditionNode.ParentID,
		NodeModel:      int32(execConditionNode.NodeModel),
		NodeName:       execConditionNode.NodeName,
		BranchLevel:    int32(execConditionNode.Level),
		ConditionExpr:  execConditionNode.ConditionExpr,
		ConditionGroup: execConditionNode.ConditionGroup,
		Status:         1,
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
			slog.Infof("节点[%v]的上节点不存在", execConditionNode.NodeID)
		}
		preNodes = append(preNodes, pre)
	}
	return &preNodes
}

func (execConditionNode *ExecConditionNode) ExecNextNodeModels(nodeModelMap map[string]entity.NodeModelBO) *[]entity.NodeModelBO {
	var nextNodes = make([]entity.NodeModelBO, 0)
	if execConditionNode.NextNodes == nil {
		return &nextNodes
	}
	for _, val := range execConditionNode.NextNodes {
		next, ok := nodeModelMap[val]
		if !ok {
			slog.Infof("节点[%v]的下节点不存在", execConditionNode.NodeID)
		}
		nextNodes = append(nextNodes, next)
	}
	return &nextNodes
}
