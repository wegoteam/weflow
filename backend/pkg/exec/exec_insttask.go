package exec

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/wegoteam/weflow/pkg/common/constant"
	"github.com/wegoteam/weflow/pkg/common/entity"
	"github.com/wegoteam/weflow/pkg/common/utils"
	"github.com/wegoteam/weflow/pkg/parser"
	"github.com/wegoteam/weflow/pkg/service"
	"github.com/wegoteam/wepkg/snowflake"
	"time"
)

// Start
// @Description: 发起实例任务
// @param modelID 模板ID
// @param userID 发起人ID
// @param userName 发起人名称
// @param params 参数
func Start(modelID, userID, userName string, params map[string]any) string {
	instTaskExecution := &InstTaskExecution{
		Execution:  &Execution{},
		ModelID:    modelID,
		OpUserName: userName,
		OpUserID:   userID,
	}
	return instTaskExecution.start(modelID, params)
}

// Stop
// @Description: 停止实例任务
// @param instTaskID 实例任务ID
// @param opUserID 操作人ID
// @param opUserName 操作人名称
// @param opinionDesc 意见描述
// @return bool
func Stop(instTaskID, opUserID, opUserName, opinionDesc string) bool {
	instTaskExecution := NewInstTaskExecution(instTaskID)
	instTaskExecution.OpUserID = opUserID
	instTaskExecution.OpUserName = opUserName
	instTaskExecution.OpinionDesc = opinionDesc
	instTaskExecution.stop(instTaskID)
	return true
}

// Suspend
// @Description: 暂停、挂起实例任务
// @param instTaskID 实例任务ID
// @param opUserID 操作人ID
// @param opUserName 操作人名称
// @param opinionDesc 意见描述
// @return bool
func Suspend(instTaskID, opUserID, opUserName, opinionDesc string) bool {
	instTaskExecution := NewInstTaskExecution(instTaskID)
	instTaskExecution.OpUserID = opUserID
	instTaskExecution.OpUserName = opUserName
	instTaskExecution.OpinionDesc = opinionDesc
	instTaskExecution.suspend(instTaskID)
	return true
}

// Sesume
// @Description: 恢复实例任务
// @param instTaskID 实例任务ID
// @param opUserID 操作人ID
// @param opUserName 操作人名称
// @param opinionDesc 意见描述
// @return bool
func Sesume(instTaskID, opUserID, opUserName, opinionDesc string) bool {
	instTaskExecution := NewInstTaskExecution(instTaskID)
	instTaskExecution.OpUserID = opUserID
	instTaskExecution.OpUserName = opUserName
	instTaskExecution.OpinionDesc = opinionDesc
	instTaskExecution.resume(instTaskID)
	return true
}

// start
// @Description: 发起实例任务
// @receiver execution
// @param modelId 模板ID
// @param userName 发起人名称
// @param userID 发起人ID
// @param params 参数
// @return string
func (instTaskExecution *InstTaskExecution) start(modelID string, params map[string]any) string {
	execution := instTaskExecution.Execution
	modelVersion := service.GetEnableModelVersion(modelID)
	if modelVersion == nil {
		hlog.Errorf("模板ID[{}]不存在或者模板未发布可用版本", modelID)
		panic("模板不存在或者模板未发布可用版本")
	}
	instTaskExecution.VersionID = modelVersion.VersionID
	execution.CreateUserName = instTaskExecution.OpUserName
	execution.CreateUserID = instTaskExecution.OpUserID
	execution.ProcessDefId = modelVersion.ProcessDefID
	execution.FormDefId = modelVersion.FormDefID
	//获取流程定义模型
	processDefModel := parser.GetProcessDefModel(modelVersion.ProcessDefID)
	execution.ProcessDefModel = processDefModel
	execution.InstTaskID = snowflake.GetSnowflakeId()
	execution.InstTaskName = modelVersion.ModelTitle
	execution.InstTaskStatus = constant.InstanceTaskStatusDoing
	execution.Now = time.Now()
	startNodeId := processDefModel.StartNodeId
	startNode := processDefModel.NodeModelMap[startNodeId]
	//实例任务参数
	if params == nil {
		var instTaskParamMap = make(map[string]interface{})
		execution.InstTaskParamMap = instTaskParamMap
	} else {
		execution.InstTaskParamMap = params
	}
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
	//执行节点
	execNode(&startNode, execution)
	//实例任务数据
	instTaskExecution.execStartInstData()
	hlog.Infof("实例任务[%v]的发起人[%v]发起版本[%v]的实例任务执行成功，发起参数为%v", execution.InstTaskID, instTaskExecution.OpUserID, instTaskExecution.VersionID, params)
	return execution.InstTaskID
}

// stop
// @Description: 停止实例任务
// @receiver instTaskExecution
// @param instTaskID
// @return bool
func (instTaskExecution *InstTaskExecution) stop(instTaskID string) bool {
	execution := instTaskExecution.Execution

	if utils.IsNotContainsSlice(instCanStopList, int(execution.InstTaskStatus)) {
		panic("该实例任务状态不允许终止，请检查实例任务状态")
	}
	//终止操作执行的实例数据，进行数据处理
	instTaskExecution.execStopInstData()
	hlog.Infof("实例任务[%v]的版本[%v]的实例任务终止执行成功", instTaskID, instTaskExecution.VersionID)
	return true
}

// suspend
// @Description: 挂起实例任务
// @receiver instTaskExecution
// @param instTaskID
// @return bool
func (instTaskExecution *InstTaskExecution) suspend(instTaskID string) bool {

	//挂起实例任务操作执行的实例数据，进行数据处理
	instTaskExecution.execSuspendInstData()
	hlog.Infof("实例任务[%v]的版本[%v]的实例任务挂起执行成功", instTaskID, instTaskExecution.VersionID)
	return true
}

// resume
// @Description: 恢复实例任务
// @receiver instTaskExecution
// @param instTaskID
// @return bool
func (instTaskExecution *InstTaskExecution) resume(instTaskID string) bool {

	//终止操作执行的实例数据，进行数据处理
	instTaskExecution.execResumeInstData()
	hlog.Infof("实例任务[%v]的版本[%v]的实例任务恢复执行成功", instTaskID, instTaskExecution.VersionID)
	return true
}
