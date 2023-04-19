package exec

import "github.com/wegoteam/weflow/pkg/common/entity"

// Execution 执行
type Execution struct {
	InstTaskId      string                 //实例任务ID
	ProcessDefModel entity.ProcessDefModel //流程定义
	InstTaskParam   map[string]interface{} //实例任务参数

}

// InstTaskExecution 执行实例
type InstTaskExecution struct {
}

// TaskExecution 执行任务
type TaskExecution struct {
}

type NodeExecTaskBO struct {
}
