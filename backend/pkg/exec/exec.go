package exec

import (
	"github.com/wegoteam/weflow/pkg/common/config"
	"github.com/wegoteam/weflow/pkg/common/entity"
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
	start(modelID string, params map[string]any) string
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
	Execution      *Execution //执行对象
	ModelID        string     //模型ID
	VersionID      string     //版本ID
	CreateUserID   string     //创建人ID
	CreateUserName string     //创建人名称
}

// UserTaskExecution 执行用户任务
type UserTaskExecution struct {
	Execution      *Execution //执行对象
	ModelID        string     //模型ID
	VersionID      string     //版本ID
	CreateUserID   string     //创建人ID
	CreateUserName string     //创建人名称
	NodeTaskID     string     //节点任务ID
	UserTaskID     string     //用户任务ID
}
