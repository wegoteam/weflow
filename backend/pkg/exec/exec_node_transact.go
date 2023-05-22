package exec

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/elliotchance/pie/v2"
	"github.com/wegoteam/weflow/pkg/common/constant"
	"github.com/wegoteam/weflow/pkg/common/entity"
	"github.com/wegoteam/wepkg/snowflake"
	"time"
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
	_, ok := execution.ExecNodeTaskMap[execTransactNode.NodeID]
	if ok {
		hlog.Warnf("实例任务[%s]的流程定义[%s]执行办理节点[%s]节点名称[%s]已经生成节点任务，该节点重复执行", execution.InstTaskID, execution.ProcessDefId, execTransactNode.NodeID, execTransactNode.NodeName)
		return ExecResult{}
	}
	hlog.Infof("实例任务[%s]的流程定义[%s]执行办理节点[%s]节点名称[%s]生成节点任务", execution.InstTaskID, execution.ProcessDefId, execTransactNode.NodeID, execTransactNode.NodeName)
	processDefModel := execution.ProcessDefModel
	nodeTaskId := snowflake.GetSnowflakeId()
	//生成执行节点任务
	var execNodeTask = &entity.ExecNodeTaskBO{
		NodeTaskID: nodeTaskId,
		NodeModel:  execTransactNode.NodeModel,
		NodeID:     execTransactNode.NodeID,
		Status:     constant.InstanceNodeTaskStatusDoing,
	}
	execution.ExecNodeTaskMap[execTransactNode.NodeID] = *execNodeTask
	//生成实例节点任务
	instNodeTasks := execution.InstNodeTasks
	var instNodeTask = execTransactNode.GetInstNodeTask(execution.InstTaskID, nodeTaskId, execution.Now)
	*instNodeTasks = append(*instNodeTasks, instNodeTask)
	//生成实例节点任务表单权限
	instNodeTaskForms := execution.TaskFormPers
	addInstNodeTaskForms := execTransactNode.GetTaskFormPers(execTransactNode.FormPer, instNodeTask)
	*instNodeTaskForms = append(*instNodeTaskForms, addInstNodeTaskForms...)
	//生成用户任务
	userTasks := execution.UserTasks
	//addUserTasks := GetUserTask(instNodeTask, execTransactNode.NodeHandler)
	addUserTasks := ExecUserTask(*execution, instNodeTask, execTransactNode.NodeHandler)
	*userTasks = append(*userTasks, addUserTasks...)
	//执行任务
	nextNodes := execTransactNode.ExecNextNodeModels(processDefModel.NodeModelMap)
	return ExecResult{
		NextNodes: nextNodes,
	}
}

/**
获取实例节点任务
*/
func (execTransactNode *ExecTransactNode) GetInstNodeTask(instTaskID, nodeTaskID string, now time.Time) entity.InstNodeTaskBO {
	//生成实例节点任务
	var instNodeTask = entity.InstNodeTaskBO{
		ExecOpType:     constant.OperationTypeAdd,
		InstTaskID:     instTaskID,
		NodeTaskID:     nodeTaskID,
		NodeID:         execTransactNode.NodeID,
		ParentID:       execTransactNode.ParentID,
		NodeModel:      int32(execTransactNode.NodeModel),
		NodeName:       execTransactNode.NodeName,
		NoneHandler:    int32(execTransactNode.NoneHandler),
		AppointHandler: execTransactNode.AppointHandler,
		HandleMode:     int32(execTransactNode.HandleMode),
		FinishMode:     int32(execTransactNode.FinishMode),
		Status:         constant.InstanceNodeTaskStatusDoing,
		CreateTime:     now,
		UpdateTime:     now,
	}
	return instNodeTask
}

/**
获取实例节点任务表单权限
*/
func (execTransactNode *ExecTransactNode) GetTaskFormPers(formPers []entity.FormPer, instNodeTask entity.InstNodeTaskBO) []entity.TaskFormPerBO {
	var taskFormPers = make([]entity.TaskFormPerBO, len(formPers))
	for ind, formPer := range formPers {
		var taskFormPerBO = entity.TaskFormPerBO{
			ExecOpType: constant.OperationTypeAdd,
			InstTaskID: instNodeTask.InstTaskID,
			NodeTaskID: instNodeTask.NodeTaskID,
			NodeID:     instNodeTask.NodeID,
			ElemID:     formPer.ElemID,
			ElemPID:    formPer.ElemPID,
			Per:        int32(formPer.Per),
		}
		taskFormPers[ind] = taskFormPerBO
	}
	return taskFormPers
}

func (execTransactNode *ExecTransactNode) ExecPreNodeModels(nodeModelMap map[string]entity.NodeModelBO) *[]entity.NodeModelBO {
	var preNodes = make([]entity.NodeModelBO, 0)
	if execTransactNode.PreNodes == nil {
		return &preNodes
	}
	for _, val := range execTransactNode.PreNodes {
		pre, ok := nodeModelMap[val]
		if !ok {
			hlog.Infof("节点[%v]的上节点不存在", execTransactNode.NodeID)
		}
		preNodes = append(preNodes, pre)
	}
	return &preNodes
}

func (execTransactNode *ExecTransactNode) ExecNextNodeModels(nodeModelMap map[string]entity.NodeModelBO) *[]entity.NodeModelBO {
	var nextNodes = make([]entity.NodeModelBO, 0)
	//判断是否有下节点
	if execTransactNode.NextNodes != nil {
		for _, val := range execTransactNode.NextNodes {
			next, ok := nodeModelMap[val]
			if !ok {
				hlog.Infof("节点[%s]的下节点不存在", execTransactNode.NodeID)
			}
			nextNodes = append(nextNodes, next)
		}
	}
	//判断下节点是否为父节点
	if isParent(execTransactNode.ParentID) {
		return &nextNodes
	}
	//判断节点的父节点是否是分支节点，节点是否在分支节点的最后节点上
	pNodeModel, ok := nodeModelMap[execTransactNode.ParentID]
	if !ok {
		hlog.Warnf("节点[%s]的父节点不存在", execTransactNode.NodeID)
		return &nextNodes
	}
	if pNodeModel.NodeModel != constant.BranchNodeModel {
		hlog.Warnf("节点[%s]的父节点[%s]错误，该节点的父节点不是分支节点", execTransactNode.NodeID, execTransactNode.ParentID)
		return &nextNodes
	}
	branchNodeModel := NewBranchNode(&pNodeModel)
	if branchNodeModel.LastNodes == nil {
		hlog.Warnf("节点[%s]的父节点[%s]错误，该分支节点的最后节点为空", execTransactNode.NodeID, execTransactNode.ParentID)
		return &nextNodes
	}
	if pie.Contains(branchNodeModel.LastNodes, execTransactNode.NodeID) {
		nextNodes = append(nextNodes, pNodeModel)
	}
	return &nextNodes
}
