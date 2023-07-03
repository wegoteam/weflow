package bo

// UserTaskResult @description And so forth.
type UserTaskResult struct {
	InstTaskID     string `json:"instTaskID"`     // 实例任务id
	TaskName       string `json:"taskName"`       // 实例任务名称
	InstStatus     int32  `json:"instStatus"`     // 任务状态【1：创建中(草稿)；2：进行中； 3：终止； 4：完成； 5：挂起；6：回退】
	StartTime      string `json:"startTime"`      // 发起时间
	EndTime        string `json:"endTime"`        // 结束时间
	CreateUserID   string `json:"createUserID"`   // 创建人id
	CreateUserName string `json:"createUserName"` // 创建人名称
	NodeTaskID     string `json:"nodeTaskID"`     // 节点任务id
	NodeID         string `json:"nodeID"`         // 节点id
	ParentID       string `json:"parentID"`       // 父节点id
	NodeModel      int32  `json:"nodeModel"`      // 节点模型【1：开始节点；2：审批节点；3：办理节点；4：抄送节点；5：自定义节点；6：条件节点；7：分支节点；8：汇聚节点；9：结束节点】
	NodeName       string `json:"nodeName"`       // 节点名称
	UserTaskID     string `json:"userTaskID"`     // 处理人任务id
}

// UserTaskTodoResult
// @Description: 用户待办任务查询结果
type UserTaskTodoResult struct {
	InstTaskID     string `json:"instTaskID"`     // 实例任务id
	TaskName       string `json:"taskName"`       // 实例任务名称
	InstStatus     int32  `json:"instStatus"`     // 任务状态【1：创建中(草稿)；2：进行中； 3：终止； 4：完成； 5：挂起；6：回退】
	StartTime      string `json:"startTime"`      // 发起时间 yyyy-MM-dd HH:mm:ss:SSS
	EndTime        string `json:"endTime"`        // 结束时间 yyyy-MM-dd HH:mm:ss:SSS
	CreateUserID   string `json:"createUserID"`   // 创建人id
	CreateUserName string `json:"createUserName"` // 创建人名称
	NodeTaskID     string `json:"nodeTaskID"`     // 节点任务id
	NodeID         string `json:"nodeID"`         // 节点id
	ParentID       string `json:"parentID"`       // 父节点id
	NodeModel      int32  `json:"nodeModel"`      // 节点模型【1：开始节点；2：审批节点；3：办理节点；4：抄送节点；5：自定义节点；6：条件节点；7：分支节点；8：汇聚节点；9：结束节点】
	NodeName       string `json:"nodeName"`       // 节点名称
	UserTaskID     string `json:"userTaskID"`     // 处理人任务id
}

// UserTaskAgreeBO
// @Description: 同意用户任务
type UserTaskAgreeBO struct {
	UserTaskID  string         // 用户任务id
	OpUserID    string         // 操作人id
	OpUserName  string         // 操作人名称
	OpinionDesc string         // 操作意见
	Params      map[string]any // 参数
}

// UserTaskDisagreeBO
// @Description: 不同意用户任务
type UserTaskDisagreeBO struct {
	UserTaskID  string // 用户任务id
	OpUserID    string // 操作人id
	OpUserName  string // 操作人名称
	OpinionDesc string // 操作意见
}

// UserTaskSaveBO
// @Description: 保存用户任务
type UserTaskSaveBO struct {
	UserTaskID  string         // 用户任务id
	OpUserID    string         // 操作人id
	OpUserName  string         // 操作人名称
	OpinionDesc string         // 操作意见
	Params      map[string]any // 参数
}
