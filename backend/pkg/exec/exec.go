package exec

import (
	"github.com/wegoteam/weflow/pkg/common/config"
	"github.com/wegoteam/weflow/pkg/common/entity"
	"github.com/wegoteam/weflow/pkg/parser"
	"github.com/wegoteam/weflow/pkg/service"
	"time"
)

var (
	MysqlDB     = config.MysqlDB
	RedisCliet  = config.RedisCliet
	MongoClient = config.MongoClient
)

// IExecution
// @Description: 执行接口
type IExecution interface {
	//开始
	start(modelID string, params map[string]any) string
	//停止
	stop(instTaskID string) bool
	//暂停
	suspend(instTaskID string) bool
	//恢复
	resume(instTaskID string) bool
}

// ITaskExecution
// @Description: 执行接口
type ITaskExecution interface {
	//同意处理审批、办理、抄送、自定义节点的任务
	agree(userTaskID string, params map[string]any) bool
	//不同意处理审批、办理、抄送、自定义节点的任务
	disagree(userTaskID string) bool
	//转办任务，将任务交接给他人办理，办理完成后继续下步骤
	turn() bool
	//委托任务，将任务委托给他人，他人办理完成后再回到委托人
	delegate() bool
	//退回：退回上节点
	rollback() bool
	//退回：退回开始节点
	rollbackStartNode() bool
	//退回：退回指定节点
	rollbackAnyNode() bool
	//撤回，处理人撤回
	revoke() bool
	//撤销：发起人撤销
	cancel() bool
	//催办
	urge() bool
	//保存
	save() bool
	//加签
	addSign() bool
	//减签
	reduceSign() bool
	//抄送
	cc() bool
	//抄送回复
	ccReply() bool
	//抄送撤回
	ccRevoke() bool
}

// Execution 执行对象
type Execution struct {
	InstTaskID       string                           //实例任务ID
	ProcessDefId     string                           //流程定义ID
	FormDefId        string                           //表单定义ID
	InstTaskStatus   int8                             //实例任务状态
	InstTaskName     string                           //实例任务名称
	CreateUserID     string                           //创建人ID
	CreateUserName   string                           //创建人名称
	Now              time.Time                        //当前时间
	ProcessDefModel  *entity.ProcessDefModel          //流程定义
	InstTaskParamMap map[string]interface{}           //实例任务参数
	ExecNodeTaskMap  map[string]entity.ExecNodeTaskBO //实例节点任务执行缓存数据
	UserTasks        *[]entity.UserTaskBO             //用户任务
	InstNodeTasks    *[]entity.InstNodeTaskBO         //实例节点任务
	TaskFormPers     *[]entity.TaskFormPerBO          //实例节点任务表单权限
	InstTaskOpLogs   *[]entity.InstTaskOpLogBO        //实例任务操作日志
}

// InstTaskExecution 执行实例
type InstTaskExecution struct {
	Execution  *Execution //执行对象
	ModelID    string     //模型ID
	VersionID  string     //版本ID
	OpUserID   string     //操作用户ID
	OpUserName string     //操作用户名称
}

// UserTaskExecution 执行用户任务
type UserTaskExecution struct {
	Execution      *Execution //执行对象
	ModelID        string     //模型ID
	VersionID      string     //版本ID
	NodeID         string     //节点ID
	NodeTaskID     string     //节点任务ID
	ApproveType    int        // 审批类型【人工审批：1；自动通过：2；自动拒绝】默认人工审批1
	NoneHandler    int        // 审批人为空时【自动通过：1；自动转交管理员：2；指定审批人：3】默认自动通过1
	AppointHandler string     // 审批人为空时指定审批人ID
	HandleMode     int        // 审批方式【依次审批：1；会签（需要完成人数的审批人同意或拒绝才可完成节点）：2；或签（其中一名审批人同意或拒绝即可）：3】默认会签2
	FinishMode     int        // 完成人数：依次审批默认0所有人不可选人，会签默认0所有人（可选人大于0），或签默认1一个人（可选人大于0）
	NodeTaskStatus int        //节点任务状态
	UserTaskID     string     //用户任务ID
	UserTaskStatus int        //用户任务状态
	OpUserID       string     //操作用户ID
	OpUserName     string     //操作用户名称
	OpinionDesc    string     //意见描述
	Opinion        string     //意见
	OpSort         int        //操作排序
}

// NewInstTaskExecution
// @Description: 创建实例任务执行对象
// @param instTaskID
// @return *InstTaskExecution
func NewInstTaskExecution(instTaskID string) *InstTaskExecution {
	//实例任务
	instTask := service.GetInstTask(instTaskID)
	if instTask == nil {
		panic("实例任务不存在")
	}
	var execution = &Execution{}
	//获取流程定义模型
	processDefModel := parser.GetProcessDefModel(instTask.ProcessDefID)
	execution.ProcessDefModel = processDefModel
	execution.InstTaskID = instTask.InstTaskID
	execution.ProcessDefId = instTask.ProcessDefID
	execution.FormDefId = instTask.FormDefID
	execution.CreateUserName = instTask.CreateUserName
	execution.CreateUserID = instTask.CreateUserID
	execution.InstTaskName = instTask.TaskName
	execution.InstTaskStatus = int8(instTask.Status)
	execution.Now = time.Now()
	//实例任务参数
	var instTaskParamMap = make(map[string]interface{})
	execution.InstTaskParamMap = instTaskParamMap
	//实例节点任务执行缓存数据
	var execNodeTaskMap = make(map[string]entity.ExecNodeTaskBO)
	execution.ExecNodeTaskMap = execNodeTaskMap
	//用户任务
	var userTasks = make([]entity.UserTaskBO, 0)
	execution.UserTasks = &userTasks
	//实例节点任务
	var instNodeTasks = make([]entity.InstNodeTaskBO, 0)
	execution.InstNodeTasks = &instNodeTasks
	//实例节点任务表单权限
	var taskFormPers = make([]entity.TaskFormPerBO, 0)
	execution.TaskFormPers = &taskFormPers
	//实例任务操作日志
	var instTaskOpLogs = make([]entity.InstTaskOpLogBO, 0)
	execution.InstTaskOpLogs = &instTaskOpLogs
	return &InstTaskExecution{
		Execution:  execution,
		ModelID:    instTask.ModelID,
		VersionID:  instTask.VersionID,
		OpUserID:   instTask.CreateUserID,
		OpUserName: instTask.CreateUserName,
	}
}

// NewUserTaskExecution
// @Description: 创建用户任务执行对象
// @param userTaskID
// @return *UserTaskExecution
func NewUserTaskExecution(userTaskID string) *UserTaskExecution {
	//实例任务
	instNodeUserTask := service.GetInstNodeUserTask(userTaskID)
	if instNodeUserTask == nil {
		panic("实例用户任务不存在")
	}
	var execution = &Execution{}
	//获取流程定义模型
	processDefModel := parser.GetProcessDefModel(instNodeUserTask.ProcessDefID)
	execution.ProcessDefModel = processDefModel
	execution.InstTaskID = instNodeUserTask.InstTaskID
	execution.ProcessDefId = instNodeUserTask.ProcessDefID
	execution.FormDefId = instNodeUserTask.FormDefID
	execution.CreateUserName = instNodeUserTask.CreateUserName
	execution.CreateUserID = instNodeUserTask.CreateUserID
	execution.InstTaskName = instNodeUserTask.TaskName
	execution.InstTaskStatus = int8(instNodeUserTask.TStatus)
	execution.Now = time.Now()
	//实例任务参数
	var instTaskParamMap = make(map[string]interface{})
	execution.InstTaskParamMap = instTaskParamMap
	//实例节点任务执行缓存数据
	execNodeTaskMap := service.GetExecNodeTaskMap(instNodeUserTask.InstTaskID)
	execution.ExecNodeTaskMap = execNodeTaskMap
	//用户任务
	var userTasks = make([]entity.UserTaskBO, 0)
	execution.UserTasks = &userTasks
	//实例节点任务
	var instNodeTasks = make([]entity.InstNodeTaskBO, 0)
	execution.InstNodeTasks = &instNodeTasks
	//实例节点任务表单权限
	var taskFormPers = make([]entity.TaskFormPerBO, 0)
	execution.TaskFormPers = &taskFormPers
	//实例任务操作日志
	var instTaskOpLogs = make([]entity.InstTaskOpLogBO, 0)
	execution.InstTaskOpLogs = &instTaskOpLogs
	return &UserTaskExecution{
		Execution:      execution,
		ModelID:        instNodeUserTask.ModelID,
		VersionID:      instNodeUserTask.VersionID,
		NodeID:         instNodeUserTask.NodeID,
		NodeTaskID:     instNodeUserTask.NodeTaskID,
		ApproveType:    int(instNodeUserTask.ApproveType),
		NoneHandler:    int(instNodeUserTask.NoneHandler),
		AppointHandler: instNodeUserTask.AppointHandler,
		HandleMode:     int(instNodeUserTask.HandleMode),
		FinishMode:     int(instNodeUserTask.FinishMode),
		NodeTaskStatus: int(instNodeUserTask.NStatus),
		OpUserID:       instNodeUserTask.OpUserID,
		OpUserName:     instNodeUserTask.OpUserName,
		UserTaskID:     instNodeUserTask.UserTaskID,
		OpSort:         int(instNodeUserTask.Sort),
	}
}
