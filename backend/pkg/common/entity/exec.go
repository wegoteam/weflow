package entity

import "time"

// Execution 执行对象
type Execution struct {
	InstTaskId       string                     `json:"instTaskId"`       //实例任务ID
	InstStatus       int8                       `json:"instStatus"`       //实例状态
	Now              time.Time                  `json:"now"`              //当前时间
	InstTaskName     string                     `json:"instTaskName"`     //实例任务名称
	ProcessDefModel  ProcessDefModel            `json:"processDefModel"`  //流程定义
	InstTaskParamMap *map[string]interface{}    `json:"instTaskParamMap"` //实例任务参数
	ExecNodeTaskMap  *map[string]ExecNodeTaskBO `json:"execNodeTaskMap"`  //实例节点任务执行缓存数据
	UserTasks        *[]UserTaskBO              `json:"userTasks"`        //用户任务
	InstNodeTasks    *[]InstNodeTask            `json:"instNodeTasks"`    //实例节点任务
}

// ExecNodeTaskBO 执行的节点任务，执行流转任务
type ExecNodeTaskBO struct {
}

// UserTaskBO 用户任务
type UserTaskBO struct {
}

// InstNodeTask 实例节点任务
type InstNodeTask struct {
}

// InstTaskExecution 执行实例
type InstTaskExecution struct {
	Execution Execution `json:"execution"` //执行对象
}

// UserTaskExecution 执行用户任务
type UserTaskExecution struct {
	Execution Execution `json:"execution"` //执行对象
}
