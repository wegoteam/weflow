package exec

import (
	"github.com/gookit/slog"
	"github.com/wegoteam/weflow/pkg/common/entity"
	"github.com/wegoteam/weflow/pkg/expr"
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

func (execConditionNode *ExecConditionNode) ExecCurrNodeModel(exec *entity.Execution) ExecResult {
	slog.Infof("ExecConditionNode 执行条件节点")

	//条件
	conditions := execConditionNode.ConditionExpr
	//参数
	paramMap := exec.InstTaskParamMap

	//执行条件
	flag := expr.ExecExpr(conditions, paramMap)
	if !flag {
		slog.Infof("节点[%v]的条件不成立", execConditionNode.NodeID)
		return ExecResult{}
	}
	processDefModel := exec.ProcessDefModel
	nextNodes := execConditionNode.ExecNextNodeModels(processDefModel.NodeModelMap)
	return ExecResult{
		NextNodes: nextNodes,
	}
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
