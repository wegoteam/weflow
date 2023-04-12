package constant

/*
	节点模型【1：开始节点；2：审批节点；3：知会节点；4：自定义节点；5：条件节点；6：分支节点；7：汇聚节点；8：结束节点】
*/
const (
	START_NODE_MODEL       = 1
	APPROVAL_NODE_MODEL    = 2
	NOTIFY_NODE_MODEL      = 3
	CUSTOM_NODE_MODEL      = 4
	CONDITION_NODE_MODEL   = 5
	BRANCH_NODE_MODEL      = 6
	CONVERGENCE_NODE_MODEL = 7
	END_NODE_MODEL         = 8
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
