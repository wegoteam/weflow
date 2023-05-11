package exec

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/wegoteam/weflow/pkg/common/entity"
	"github.com/wegoteam/wepkg/snowflake"
	"time"
)

// ExecApprovalNode 审批节点
type ExecApprovalNode struct {
	NodeModel      int8               `json:"nodeModel"`      // 节点类型
	NodeName       string             `json:"nodeName"`       // 节点名称
	NodeID         string             `json:"nodeId"`         // 节点ID
	ParentID       string             `json:"parentId"`       // 父节点ID
	ApproveType    int                `json:"approveType"`    //审批类型【人工审批：1；自动通过：2；自动拒绝】默认人工审批1
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
实例化执行审批节点对象
*/
func NewApprovalNode(node *entity.NodeModelBO) *ExecApprovalNode {

	return &ExecApprovalNode{
		NodeModel:   node.NodeModel,
		NodeName:    node.NodeName,
		NodeID:      node.NodeID,
		ParentID:    node.ParentID,
		ApproveType: node.ApproveType,
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

// ExecCurrNodeModel 执行当前节点
func (execApprovalNode *ExecApprovalNode) ExecCurrNodeModel(execution *entity.Execution) ExecResult {
	processDefModel := execution.ProcessDefModel
	nodeTaskId := snowflake.GetSnowflakeId()

	//生成执行节点任务
	var execNodeTask = &entity.ExecNodeTaskBO{
		NodeTaskID: nodeTaskId,
		NodeModel:  execApprovalNode.NodeModel,
		NodeID:     execApprovalNode.NodeID,
		Status:     1,
	}
	execution.ExecNodeTaskMap[execApprovalNode.NodeID] = *execNodeTask

	//生成实例节点任务
	instNodeTasks := execution.InstNodeTasks
	var instNodeTask = execApprovalNode.GetInstNodeTask(execution.InstTaskID, nodeTaskId, execution.Now)
	*instNodeTasks = append(*instNodeTasks, instNodeTask)

	//生成用户任务
	userTasks := execution.UserTasks
	addUserTasks := GetUserTask(instNodeTask, execApprovalNode.NodeHandler)
	*userTasks = append(*userTasks, addUserTasks...)

	//执行任务
	nextNodes := execApprovalNode.ExecNextNodeModels(processDefModel.NodeModelMap)
	return ExecResult{
		NextNodes: nextNodes,
	}
}

/**
获取实例节点任务
*/
func (execApprovalNode *ExecApprovalNode) GetInstNodeTask(instTaskID, nodeTaskID string, now time.Time) entity.InstNodeTaskBO {
	//生成实例节点任务
	var instNodeTask = entity.InstNodeTaskBO{
		InstTaskID:     instTaskID,
		NodeTaskID:     nodeTaskID,
		ParentID:       execApprovalNode.ParentID,
		NodeModel:      int32(execApprovalNode.NodeModel),
		NodeName:       execApprovalNode.NodeName,
		ApproveType:    int32(execApprovalNode.ApproveType),
		NoneHandler:    int32(execApprovalNode.NoneHandler),
		AppointHandler: execApprovalNode.AppointHandler,
		HandleMode:     int32(execApprovalNode.HandleMode),
		FinishMode:     int32(execApprovalNode.FinishMode),
		Status:         1,
		CreateTime:     now,
		UpdateTime:     now,
	}

	return instNodeTask
}

// ExecPreNodeModels 获取上一节点
func (execApprovalNode *ExecApprovalNode) ExecPreNodeModels(nodeModelMap map[string]entity.NodeModelBO) *[]entity.NodeModelBO {
	var preNodes = make([]entity.NodeModelBO, 0)
	if execApprovalNode.PreNodes == nil {
		return &preNodes
	}
	for _, val := range execApprovalNode.PreNodes {
		pre, ok := nodeModelMap[val]
		if !ok {
			hlog.Infof("节点[%v]的上节点不存在", execApprovalNode.NodeID)
		}
		preNodes = append(preNodes, pre)
	}
	return &preNodes
}

// ExecNextNodeModels 获取下一节点
func (execApprovalNode *ExecApprovalNode) ExecNextNodeModels(nodeModelMap map[string]entity.NodeModelBO) *[]entity.NodeModelBO {
	var nextNodes = make([]entity.NodeModelBO, 0)
	if execApprovalNode.NextNodes == nil {
		return &nextNodes
	}
	for _, val := range execApprovalNode.NextNodes {
		next, ok := nodeModelMap[val]
		if !ok {
			hlog.Infof("节点[%v]的下节点不存在", execApprovalNode.NodeID)
		}
		nextNodes = append(nextNodes, next)
	}
	return &nextNodes
}
