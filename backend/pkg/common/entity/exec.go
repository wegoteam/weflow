package entity

import "time"

// Execution 执行对象
type Execution struct {
	InstTaskId       string                    `json:"instTaskId"`       //实例任务ID
	InstTaskStatus   int8                      `json:"instStatus"`       //实例任务状态
	Now              time.Time                 `json:"now"`              //当前时间
	InstTaskName     string                    `json:"instTaskName"`     //实例任务名称
	ProcessDefModel  *ProcessDefModel          `json:"processDefModel"`  //流程定义
	InstTaskParamMap map[string]interface{}    `json:"instTaskParamMap"` //实例任务参数
	ExecNodeTaskMap  map[string]ExecNodeTaskBO `json:"execNodeTaskMap"`  //实例节点任务执行缓存数据
	UserTasks        *[]UserTaskBO             `json:"userTasks"`        //用户任务
	InstNodeTasks    *[]InstNodeTask           `json:"instNodeTasks"`    //实例节点任务
}

// ExecNodeTaskBO 执行的节点任务，执行流转任务
type ExecNodeTaskBO struct {
	NodeTaskID string // 节点任务id
	NodeID     string // 节点id
	Status     int8   // 任务状态【0：未开始；1：处理中；2：完成；3：回退；4：终止；5：条件验证通过；6：条件验证不通过】
	NodeModel  int8   // 节点模型【1：开始节点；2：审批节点；3：知会节点；4：自定义节点；5：条件节点；6：分支节点；7：汇聚节点；8：结束节点】
}

// UserTaskBO 用户任务
type UserTaskBO struct {
	InstTaskID   string    // 实例任务id
	NodeTaskID   string    // 节点任务id
	NodeID       string    // 节点任务id
	UserTaskID   string    // 处理人任务id
	NodeUserID   string    // 节点处理人id
	NodeUserName string    // 处理人名称
	NodeUserType int32     // 处理人类型【1：操作员；2：部门；3：相对岗位；4：表单控件；5：角色；6：岗位；7：组织；8：自定义】
	OpOrigin     int32     // 操作来源【1：正常；2：加签】
	TimeLimit    int64     // 处理期限;格式：yyyymmddhhmm 可直接指定到期限的具体时间，期限支持到分钟； 0表示无期限
	Status       int32     // 任务状态【1：处理中；2：完成；3：回退；4：终止】
	CreateTime   time.Time // 创建时间
	UpdateTime   time.Time // 更新时间
	HandleTime   time.Time // 处理时间
	OpUserID     string    // 操作用户id
	OpUserName   string    // 操作用户名称
	HandlerSort  int32     // 处理人排序;处理人当前的处理排序
	Opinion      int32     // 处理意见【1：未发表；2：已阅；3：同意；4：不同意】
	OpinionDesc  string    // 处理意见描述

}

// InstNodeTask 实例节点任务
type InstNodeTask struct {
	InstTaskID     string    // 实例任务id
	NodeTaskID     string    // 节点任务id
	NodeID         string    // 节点id
	ParentID       string    // 父节点id
	NodeModel      int8      //节点模型【1：开始节点；2：审批节点；3：知会节点；4：自定义节点；5：条件节点；6：分支节点；7：汇聚节点；8：结束节点】
	NodeName       string    // 节点名称
	ForwardMode    int8      // 进行模式【1：并行 2：串行】
	CompleteConn   int32     // 节点完成条件;通过的人数，0表示所有人通过，节点才算完成
	PermissionMode int8      // 权限模式【1：协同 2：知会 3：审批；4：业务】
	AllowAdd       int8      // 允许加签【1：不能加签；2：允许加签】
	ProcessMode    int8      // 处理模式【1：人工； 2：自动；3：自动转人工】
	TimeLimit      int64     // 处理期限
	ConnData       string    // 条件数据
	FormPerData    string    // 表单权限数据
	Status         int8      // 任务状态【0：未开始；1：处理中；2：完成；3：回退；4：终止；5：条件验证通过；6：条件验证不通过】
	CreateTime     time.Time // 创建时间
	UpdateTime     time.Time // 更新时间
}

// InstTaskExecution 执行实例
type InstTaskExecution struct {
	Execution Execution `json:"execution"` //执行对象
}

// UserTaskExecution 执行用户任务
type UserTaskExecution struct {
	Execution Execution `json:"execution"` //执行对象
}
