package exec

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/wegoteam/weflow/pkg/common/entity"
	"github.com/wegoteam/weflow/pkg/parser"
	"github.com/wegoteam/wepkg/snowflake"
	"testing"
	"time"
)

func TestStartInstTask(t *testing.T) {
	processDefModel := parser.GetProcessDefModel("1640993392605401001")

	execution := &entity.Execution{}
	execution.ProcessDefModel = processDefModel
	execution.InstTaskID = snowflake.GetSnowflakeId()
	execution.InstTaskName = "测试流程"
	execution.InstTaskStatus = 1
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

	execution.ProcessDefId = "1640993392605401001"
	execution.FormDefId = "1640993392605401001"
	Exec(&startNode, execution)

	hlog.Infof("执行结果%+v", execution)
}
