package exec

import (
	"github.com/gookit/slog"
	"github.com/wegoteam/weflow/pkg/common/entity"
	"github.com/wegoteam/wepkg/snowflake"
)

// ExecTransactNode 办理节点
type ExecTransactNode struct {
	NodeModel      int8               `json:"nodeModel"`      // 节点类型
	NodeName       string             `json:"nodeName"`       // 节点名称
	NodeID         string             `json:"nodeId"`         // 节点ID
	ParentID       string             `json:"parentId"`       // 父节点ID
	FormPer        []entity.FormPer   `json:"formPer"`        // 表单权限
	NodeSetting    entity.NodeSetting `json:"nodeSetting"`    // 节点设置
	NodeHandler    entity.NodeHandler `json:"nodeHandler"`    // 节点处理人
	NoneHandler    int                `json:"noneHandler"`    //审批人为空时【自动通过：1；自动转交管理员：2；指定审批人：3】默认自动通过1--数字
	AppointHandler string             `json:"appointHandler"` //审批人为空时指定审批人ID
	HandleMode     int                `json:"handleMode"`     //审批方式【依次审批：1、会签（需要完成人数的审批人同意或拒绝才可完成节点）：2、或签（其中一名审批人同意或拒绝即可）：3】默认会签2
	FinishMode     int                `json:"finishMode"`     //完成人数：依次审批默认0所有人不可选人，会签默认0所有人（可选人大于0），或签默认1一个人（可选人大于0）
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
func NewTransactNode(node *entity.NodeModelBO) *ExecTransactNode {

	return &ExecTransactNode{
		NodeModel:   node.NodeModel,
		NodeName:    node.NodeName,
		NodeID:      node.NodeID,
		ParentID:    node.ParentID,
		FormPer:     node.FormPer,
		NodeSetting: node.NodeSetting,
		NodeHandler: node.NodeHandler,
		NoneHandler: node.NoneHandler,
		HandleMode:  node.HandleMode,
		FinishMode:  node.FinishMode,
		PreNodes:    node.PreNodes,
		NextNodes:   node.NextNodes,
		LastNodes:   node.LastNodes,
		Index:       node.Index,
		BranchIndex: node.BranchIndex,
	}
}

/**
执行节点
生成实例节点任务
执行任务
下节点
*/
func (execTransactNode *ExecTransactNode) ExecCurrNodeModel(execution *entity.Execution) ExecResult {

	slog.Infof("ExecTransactNode 执行办理节点")
	processDefModel := execution.ProcessDefModel
	nodeTaskId := snowflake.GetSnowflakeId()

	//生成执行节点任务
	var execNodeTask = &entity.ExecNodeTaskBO{
		NodeTaskID: nodeTaskId,
		NodeModel:  execTransactNode.NodeModel,
		NodeID:     execTransactNode.NodeID,
		Status:     2,
	}
	execution.ExecNodeTaskMap[execTransactNode.NodeID] = *execNodeTask

	//生成实例节点任务
	var instNodeTask = entity.InstNodeTaskBO{}
	instNodeTasks := *execution.InstNodeTasks
	instNodeTasks = append(instNodeTasks, instNodeTask)

	//生成用户任务
	var userTask = entity.UserTaskBO{}
	userTasks := *execution.UserTasks
	userTasks = append(userTasks, userTask)

	//执行任务

	nextNodes := execTransactNode.ExecNextNodeModels(processDefModel.NodeModelMap)
	return ExecResult{
		NextNodes: nextNodes,
	}
}

func (execTransactNode *ExecTransactNode) ExecPreNodeModels(nodeModelMap map[string]entity.NodeModelBO) *[]entity.NodeModelBO {
	var preNodes = make([]entity.NodeModelBO, 0)
	if execTransactNode.PreNodes == nil {
		return &preNodes
	}
	for _, val := range execTransactNode.PreNodes {
		pre, ok := nodeModelMap[val]
		if !ok {
			slog.Infof("节点[%v]的上节点不存在", execTransactNode.NodeID)
		}
		preNodes = append(preNodes, pre)
	}
	return &preNodes
}

func (execTransactNode *ExecTransactNode) ExecNextNodeModels(nodeModelMap map[string]entity.NodeModelBO) *[]entity.NodeModelBO {
	var nextNodes = make([]entity.NodeModelBO, 0)
	if execTransactNode.NextNodes == nil {
		return &nextNodes
	}
	for _, val := range execTransactNode.NextNodes {
		next, ok := nodeModelMap[val]
		if !ok {
			slog.Infof("节点[%v]的下节点不存在", execTransactNode.NodeID)
		}
		nextNodes = append(nextNodes, next)
	}
	return &nextNodes
}
