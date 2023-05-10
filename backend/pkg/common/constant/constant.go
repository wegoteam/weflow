package constant

/*
	节点模型【1：开始节点；2：审批节点；3：办理节点；4：抄送节点；5：自定义节点；6：条件节点；7：分支节点；8：汇聚节点；9：结束节点】
*/
const (
	START_NODE_MODEL       = 1
	APPROVAL_NODE_MODEL    = 2
	TRANSACT_NODE_MODEL    = 3
	NOTIFY_NODE_MODEL      = 4
	CUSTOM_NODE_MODEL      = 5
	CONDITION_NODE_MODEL   = 6
	BRANCH_NODE_MODEL      = 7
	CONVERGENCE_NODE_MODEL = 8
	END_NODE_MODEL         = 9
)

/**
Redis key的前缀
*/
const (
	//流程定义Redis的前缀
	REDIS_PROCESS_DEF_MODEL = "weflow:proess-def:"
)

const (
	HAS_REDIS_PROCESS_DEF_MODEL = 1
)
