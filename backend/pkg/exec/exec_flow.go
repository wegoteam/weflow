package exec

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/wegoteam/weflow/pkg/common/constant"
	"github.com/wegoteam/weflow/pkg/common/entity"
	"github.com/wegoteam/weflow/pkg/parser"
	"github.com/wegoteam/wepkg/snowflake"
	"time"
)

func StartProcessInstTask(modelID string, params map[string]any) {
	processDefModel := parser.GetProcessDefModel("1640993392605401001")

	execution := &Execution{}
	execution.ProcessDefModel = processDefModel
	execution.InstTaskID = snowflake.GetSnowflakeId()
	execution.InstTaskName = "测试流程"
	execution.InstTaskStatus = constant.InstanceTaskStatusDoing
	execution.Now = time.Now()
	startNodeId := processDefModel.StartNodeId
	startNode := processDefModel.NodeModelMap[startNodeId]

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

	execution.CreateUserName = "xuch01"
	execution.CreateUserID = "547"
	execution.ProcessDefId = "1640993392605401001"
	execution.FormDefId = "1640993392605401001"
	execNode(&startNode, execution)

	hlog.Infof("执行结果%+v", execution)
}

// start
// @Description: 发起实例任务
// @receiver execution
// @param modelId 模板ID
// @param params 参数
// @return string
func (instTaskExecution *InstTaskExecution) start(modelID string, params map[string]any) string {
	execution := instTaskExecution.Execution
	processDefModel := parser.GetProcessDefModel("1640993392605401001")

	execution.ProcessDefModel = processDefModel
	execution.InstTaskID = snowflake.GetSnowflakeId()
	execution.InstTaskName = "测试流程"
	execution.InstTaskStatus = constant.InstanceTaskStatusDoing
	execution.Now = time.Now()
	startNodeId := processDefModel.StartNodeId
	startNode := processDefModel.NodeModelMap[startNodeId]

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

	execution.CreateUserName = "xuch01"
	execution.CreateUserID = "547"
	execution.ProcessDefId = "1640993392605401001"
	execution.FormDefId = "1640993392605401001"
	execNode(&startNode, execution)

	hlog.Infof("执行结果%+v", execution)
	return execution.InstTaskID
}

func (execution *Execution) stop(instTaskID string) {

}

func (execution *Execution) hangup(instTaskID string) {

}
