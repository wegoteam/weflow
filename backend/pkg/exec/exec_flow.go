package exec

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/wegoteam/weflow/pkg/common/constant"
	"github.com/wegoteam/weflow/pkg/common/entity"
	"github.com/wegoteam/weflow/pkg/parser"
	"github.com/wegoteam/weflow/pkg/service"
	"github.com/wegoteam/wepkg/snowflake"
	"time"
)

// StartProcessInstTask
// @Description: 发起实例任务
// @param modelID
// @param userID
// @param userName
// @param params
func StartProcessInstTask(modelID, userID, userName string, params map[string]any) {

	instTaskExecution := &InstTaskExecution{
		CreateUserName: userName,
		CreateUserID:   userID,
	}
	instTaskExecution.start(modelID, params)
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
	execution.CreateUserName = instTaskExecution.CreateUserName
	execution.CreateUserID = instTaskExecution.CreateUserID
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
	hlog.Infof("实例任务[%v]的发起人[%v]发起版本[%v]的实例任务执行成功，发起参数为%v", execution.InstTaskID, instTaskExecution.CreateUserID, instTaskExecution.VersionID, params)
	return execution.InstTaskID
}

func (execution *Execution) stop(instTaskID string) {

}

func (execution *Execution) hangup(instTaskID string) {

}
