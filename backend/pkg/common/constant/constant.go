package constant

/*
	节点模型【1：开始节点；2：审批节点；3：办理节点；4：抄送节点；5：自定义节点；6：条件节点；7：分支节点；8：汇聚节点；9：结束节点】
*/
const (
	StartNodeModel       = 1
	ApprovalNodeModel    = 2
	TransactNodeModel    = 3
	NotifyNodeModel      = 4
	CustomNodeModel      = 5
	ConditionNodeModel   = 6
	BranchNodeModel      = 7
	ConvergenceNodeModel = 8
	EndNodeModel         = 9
)

/**
Redis key的前缀
*/
const (
	//流程定义Redis的前缀
	RedisProcessDefModel = "weflow:proess-def:"
)

const (
	HasRedisProcessDefModel = 1
)

/**
实例任务状态【1：创建中(草稿)；2：进行中； 3：终止； 4：完成； 5：挂起；6：回退】
*/
const (
	InstanceTaskStatusCreate   = 1
	InstanceTaskStatusDoing    = 2
	InstanceTaskStatusStop     = 3
	InstanceTaskStatusComplete = 4
	InstanceTaskStatusHangUp   = 5
	InstanceTaskStatusRollback = 6
)

/**
实例节点任务状态【1：未开始；2：处理中；3：完成；4：回退；5：终止；6：不通过】
*/
const (
	InstanceNodeTaskStatusNotStart = 1
	InstanceNodeTaskStatusDoing    = 2
	InstanceNodeTaskStatusComplete = 3
	InstanceNodeTaskStatusRollback = 4
	InstanceNodeTaskStatusStop     = 5
	InstanceNodeTaskStatusNotPass  = 6
)

/**
实例用户任务状态【1：处理中；2：完成；3：回退；4：终止；5：不通过】
*/
const (
	InstanceUserTaskStatusDoing    = 1
	InstanceUserTaskStatusComplete = 2
	InstanceUserTaskStatusRollback = 3
	InstanceUserTaskStatusStop     = 4
	InstanceUserTaskStatusNotPass  = 5
)

/**
实例用户任务处理意见【1：未发表；2：已阅；3：同意；4：不同意】
*/
const (
	InstanceUserTaskOpinionNotPublish = 1
	InstanceUserTaskOpinionRead       = 2
	InstanceUserTaskOpinionAgree      = 3
	InstanceUserTaskOpinionDisagree   = 4
)
