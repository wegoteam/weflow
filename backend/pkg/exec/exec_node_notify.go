package exec

import (
	"github.com/gookit/slog"
	"github.com/wegoteam/weflow/pkg/common/entity"
)

// ExecNotifyNode 知会节点
type ExecNotifyNode struct {
	NodeModel   int8               `json:"nodeModel"`   // 节点类型
	NodeName    string             `json:"nodeName"`    // 节点名称
	NodeID      string             `json:"nodeId"`      // 节点ID
	ParentID    string             `json:"parentId"`    // 父节点ID
	FormPer     []entity.FormPer   `json:"formPer"`     // 表单权限
	NodeSetting entity.NodeSetting `json:"nodeSetting"` // 节点设置
	NodeHandler entity.NodeHandler `json:"nodeHandler"` // 节点处理人
	//实例节点位置信息
	PreNodes    []string `json:"preNodes,omitempty"`    //上节点ID
	NextNodes   []string `json:"nextNodes,omitempty"`   //下节点ID
	LastNodes   []string `json:"lastNodes,omitempty"`   //分支节点尾节点ID
	Index       int      `json:"index,omitempty"`       // 下标
	BranchIndex int      `json:"branchIndex,omitempty"` // 分支下标
}

/**
实例化执行审批节点对象
*/
func NewNotifyNode(node *entity.NodeModelBO) *ExecNotifyNode {

	return &ExecNotifyNode{
		NodeModel:   node.NodeModel,
		NodeName:    node.NodeName,
		NodeID:      node.NodeID,
		ParentID:    node.ParentID,
		FormPer:     node.FormPer,
		NodeSetting: node.NodeSetting,
		NodeHandler: node.NodeHandler,
		PreNodes:    node.PreNodes,
		NextNodes:   node.NextNodes,
		LastNodes:   node.LastNodes,
		Index:       node.Index,
		BranchIndex: node.BranchIndex,
	}
}

/**
执行知会节点
生成实例节点任务
执行任务
生成知会用户任务
下节点
*/
func (execNotifyNode *ExecNotifyNode) ExecCurrNodeModel(exec *entity.Execution) ExecResult {
	slog.Infof("ExecNotifyNode 执行知会节点")
	processDefModel := exec.ProcessDefModel
	nextNodes := execNotifyNode.ExecNextNodeModels(processDefModel.NodeModelMap)
	return ExecResult{
		NextNodes: nextNodes,
	}
}

func (execNotifyNode *ExecNotifyNode) ExecPreNodeModels(nodeModelMap map[string]entity.NodeModelBO) *[]entity.NodeModelBO {
	var preNodes = make([]entity.NodeModelBO, 0)
	if execNotifyNode.PreNodes == nil {
		return &preNodes
	}
	for _, val := range execNotifyNode.PreNodes {
		pre, ok := nodeModelMap[val]
		if !ok {
			slog.Infof("节点[%v]的上节点不存在", execNotifyNode.NodeID)
		}
		preNodes = append(preNodes, pre)
	}
	return &preNodes
}

func (execNotifyNode *ExecNotifyNode) ExecNextNodeModels(nodeModelMap map[string]entity.NodeModelBO) *[]entity.NodeModelBO {
	var nextNodes = make([]entity.NodeModelBO, 0)
	if execNotifyNode.NextNodes == nil {
		return &nextNodes
	}
	for _, val := range execNotifyNode.NextNodes {
		next, ok := nodeModelMap[val]
		if !ok {
			slog.Infof("节点[%v]的下节点不存在", execNotifyNode.NodeID)
		}
		nextNodes = append(nextNodes, next)
	}
	return &nextNodes
}
