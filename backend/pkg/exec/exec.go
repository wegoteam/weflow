package exec

// Execution 执行
type Execution struct {
}

// InstTaskExecution 执行实例
type InstTaskExecution struct {
}

// TaskExecution 执行任务
type TaskExecution struct {
}

type INodeExec interface {
	NextNode()
	PreNode()
	IsParent()
	IsBranchFristNode()
}
