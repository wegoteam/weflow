package entity

import "time"

// Execution 执行对象
type Execution struct {
	InstTaskID       string                    `json:"instTaskId"`       //实例任务ID
	ProcessDefId     string                    `json:"processDefId"`     //流程定义ID
	FormDefId        string                    `json:"formDefId"`        //表单定义ID
	InstTaskStatus   int8                      `json:"instTaskStatus"`   //实例任务状态
	InstTaskName     string                    `json:"instTaskName"`     //实例任务名称
	CreateUserID     string                    `json:"createUserId"`     //创建人ID
	CreateUserName   string                    `json:"createUserName"`   //创建人名称
	Now              time.Time                 `json:"now"`              //当前时间
	ProcessDefModel  *ProcessDefModel          `json:"processDefModel"`  //流程定义
	InstTaskParamMap map[string]interface{}    `json:"instTaskParamMap"` //实例任务参数
	ExecNodeTaskMap  map[string]ExecNodeTaskBO `json:"execNodeTaskMap"`  //实例节点任务执行缓存数据
	UserTasks        *[]UserTaskBO             `json:"userTasks"`        //用户任务
	InstNodeTasks    *[]InstNodeTaskBO         `json:"instNodeTasks"`    //实例节点任务
	TaskFormPers     *[]TaskFormPerBO          `json:"taskFormPers"`     //实例节点任务表单权限
	InstTaskOpLogs   *[]InstTaskOpLogBO        `json:"instTaskOpLogs"`   //实例任务操作日志
}

// ExecNodeTaskBO 执行的节点任务，执行流转任务
type ExecNodeTaskBO struct {
	NodeTaskID string // 节点任务id
	NodeID     string // 节点id
	Status     int8   // 任务状态【1：未开始；2：处理中；3：完成；4：回退；5：终止；6：不通过】
	NodeModel  int8   // 节点模型【1：开始节点；2：审批节点；3：办理节点；4：抄送节点；5：自定义节点；6：条件节点；7：分支节点；8：汇聚节点；9：结束节点】
}

// UserTaskBO 用户任务
type UserTaskBO struct {
	ExecOpType   int8      // 执行操作类型【添加：1；修改：2；删除：3】
	InstTaskID   string    // 实例任务id
	NodeTaskID   string    // 节点任务id
	NodeID       string    // 节点任务id
	UserTaskID   string    // 处理人任务id
	Type         int32     // 常用审批人【指定成员：1；发起人自己：2；发起人自选：3：角色：4；部门：5】主管（相对岗位）【直属主管：1；部门主管：2；连续多级主管：3；部门控件对应主管：4】其他【表单人员控件：1；部门控件：2；角色控件：3】
	Strategy     int32     // 处理人策略【常用审批人：1；主管（相对岗位）：2；其他：3】
	NodeUserName string    // 处理人名称
	NodeUserID   string    // 处理人id
	Sort         int32     // 处理人顺序;正序排序
	Obj          string    // 扩展字段，设计中可忽略
	Relative     string    // 相对发起人的直属主管，设计中可忽略
	Status       int32     // 任务状态【1：处理中；2：完成；3：回退；4：终止；5：不通过】
	CreateTime   time.Time // 创建时间
	UpdateTime   time.Time // 更新时间
	HandleTime   time.Time // 处理时间
	OpUserID     string    // 操作用户id
	OpUserName   string    // 操作用户名称
	Opinion      int32     // 任务处理意见【1：未发表；2：已阅；3：同意；4：不同意】
	OpinionDesc  string    // 处理意见描述
}

// InstNodeTaskBO 实例节点任务
type InstNodeTaskBO struct {
	ExecOpType     int8      // 执行操作类型【添加：1；修改：2；删除：3】
	InstTaskID     string    // 实例任务id
	NodeTaskID     string    // 节点任务id
	NodeID         string    // 节点id
	ParentID       string    // 父节点id
	NodeModel      int32     // 节点模型【1：开始节点；2：审批节点；3：办理节点；4：抄送节点；5：自定义节点；6：条件节点；7：分支节点；8：汇聚节点；9：结束节点】
	NodeName       string    // 节点名称
	ApproveType    int32     // 审批类型【人工审批：1；自动通过：2；自动拒绝】默认人工审批1
	NoneHandler    int32     // 审批人为空时【自动通过：1；自动转交管理员：2；指定审批人：3】默认自动通过1
	AppointHandler string    // 审批人为空时指定审批人ID
	HandleMode     int32     // 审批方式【依次审批：1、会签（需要完成人数的审批人同意或拒绝才可完成节点）：2、或签（其中一名审批人同意或拒绝即可）：3】默认会签2
	FinishMode     int32     // 完成人数：依次审批默认0所有人不可选人，会签默认0所有人（可选人大于0），或签默认1一个人（可选人大于0）
	BranchMode     int32     // 分支执行方式【单分支：1；多分支：2】默认多分支2
	DefaultBranch  int32     // 单分支处理需要默认分支，在条件优先级无法处理时候执行默认分支，取值分支下标
	BranchLevel    int32     // 优先级，分支执行方式为多分支处理方式无优先级应为0
	ConditionGroup string    // 条件组前端描述展示条件组
	ConditionExpr  string    // 条件组解析后的表达式
	Remark         string    // 节点描述
	Status         int32     // 任务状态【1：未开始；2：处理中；3：完成；4：回退；5：终止；6：不通过】
	CreateTime     time.Time // 创建时间
	UpdateTime     time.Time // 更新时间
}

// TaskFormPerBO 任务表单权限
type TaskFormPerBO struct {
	ExecOpType int8   // 执行操作类型【添加：1；修改：2；删除：3】
	InstTaskID string // 实例任务id
	NodeTaskID string // 节点任务id
	NodeID     string // 节点ID
	ElemID     string //表单元素ID
	ElemPID    string //表单元素父ID
	Per        int32  // 表单权限【可编辑：1；只读：2；隐藏：3】默认只读2
}

// InstTaskOpLogBO 实例任务操作日志
type InstTaskOpLogBO struct {
	InstTaskID string    // 实例任务id
	NodeID     string    // 节点任务id
	NodeName   string    // 节点名称
	CreateTime time.Time // 创建时间
	UpdateTime time.Time // 更新时间
	Type       int32     // 类型【1：节点；2：任务；3：其他】
	Remark     string    // 描述
}

// InstTaskExecution 执行实例
type InstTaskExecution struct {
	Execution Execution `json:"execution"` //执行对象
}

// UserTaskExecution 执行用户任务
type UserTaskExecution struct {
	Execution Execution `json:"execution"` //执行对象
}
