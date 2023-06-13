package entity

import "time"

// InstTaskResult
// @Description: 实例任务返回结果
type InstTaskResult struct {
	ID             int64     // 唯一id
	InstTaskID     string    // 实例任务id
	ModelID        string    // 模板id
	ProcessDefID   string    // 流程定义id
	FormDefID      string    // 表单定义id
	VersionID      string    // 版本id
	TaskName       string    // 实例任务名称
	Status         int32     // 任务状态【1：创建中(草稿)；2：进行中； 3：终止； 4：完成； 5：挂起；6：回退】
	Remark         string    // 描述
	CreateTime     time.Time // 创建时间
	CreateUserID   string    // 创建人id
	CreateUserName string    // 创建人名称
	UpdateTime     time.Time // 更新时间
	UpdateUserID   string    // 更新人id
	UpdateUserName string    // 更新人名称
	StartTime      time.Time // 发起时间
	EndTime        time.Time // 结束时间
}

// InstNodeAndUserTaskResult
// @Description: 实例任务、节点任务、用户任务返回结果
type InstNodeAndUserTaskResult struct {
	//实例任务
	TID            int64     // 唯一id
	InstTaskID     string    // 实例任务id
	ModelID        string    // 模板id
	ProcessDefID   string    // 流程定义id
	FormDefID      string    // 表单定义id
	VersionID      string    // 版本id
	TaskName       string    // 实例任务名称
	TStatus        int32     // 任务状态【1：创建中(草稿)；2：进行中； 3：终止； 4：完成； 5：挂起；6：回退】
	TRemark        string    // 描述
	TCreateTime    time.Time // 创建时间
	CreateUserID   string    // 创建人id
	CreateUserName string    // 创建人名称
	TUpdateTime    time.Time // 更新时间
	UpdateUserID   string    // 更新人id
	UpdateUserName string    // 更新人名称
	StartTime      time.Time // 发起时间
	EndTime        time.Time // 结束时间
	//实例节点任务
	NID            int64     // 唯一id
	NodeTaskID     string    // 节点任务id
	NodeID         string    // 节点id
	ParentID       string    // 父节点id
	NodeModel      int32     // 节点模型【1：开始节点；2：审批节点；3：办理节点；4：抄送节点；5：自定义节点；6：条件节点；7：分支节点；8：汇聚节点；9：结束节点】
	NodeName       string    // 节点名称
	ApproveType    int32     // 审批类型【人工审批：1；自动通过：2；自动拒绝】默认人工审批1
	NoneHandler    int32     // 审批人为空时【自动通过：1；自动转交管理员：2；指定审批人：3】默认自动通过1
	AppointHandler string    // 审批人为空时指定审批人ID
	HandleMode     int32     // 审批方式【依次审批：1；会签（需要完成人数的审批人同意或拒绝才可完成节点）：2；或签（其中一名审批人同意或拒绝即可）：3】默认会签2
	FinishMode     int32     // 完成人数：依次审批默认0所有人不可选人，会签默认0所有人（可选人大于0），或签默认1一个人（可选人大于0）
	BranchMode     int32     // 分支执行方式【单分支：1；多分支：2】默认多分支2
	DefaultBranch  int32     // 单分支处理需要默认分支，在条件优先级无法处理时候执行默认分支，取值分支下标
	BranchLevel    int32     // 优先级，分支执行方式为多分支处理方式无优先级应为0
	ConditionGroup string    // 条件组前端描述展示条件组
	ConditionExpr  string    // 条件组解析后的表达式
	NRemark        string    // 节点描述
	NStatus        int32     // 任务状态【1：未开始；2：处理中；3：完成；4：回退；5：终止；6：不通过】
	NCreateTime    time.Time // 创建时间
	NUpdateTime    time.Time // 更新时间
	//实例用户任务
	UID          int64     // 唯一id
	UserTaskID   string    // 处理人任务id
	Type         int32     // 常用审批人【指定成员：1；发起人自己：2；发起人自选：3：角色：4；部门：5】主管（相对岗位）【直属主管：1；部门主管：2；连续多级主管：3；部门控件对应主管：4】其他【表单人员控件：1；部门控件：2；角色控件：3】
	Strategy     int32     // 处理人策略【常用审批人：1；主管（相对岗位）：2；其他：3】
	NodeUserName string    // 处理人名称
	NodeUserID   string    // 处理人id
	Sort         int32     // 处理人顺序;正序排序
	Obj          string    // 扩展字段，设计中可忽略
	Relative     string    // 相对发起人的直属主管，设计中可忽略
	UStatus      int32     // 实例用户任务状态【1：处理中；2：完成（同意）；3：不通过（不同意）；4：回退；5：终止】
	UCreateTime  time.Time // 创建时间
	UUpdateTime  time.Time // 更新时间
	HandleTime   time.Time // 处理时间
	OpUserID     string    // 操作用户id
	OpUserName   string    // 操作用户名称
	Opinion      int32     // 任务处理意见【1：未发表；2：已阅；3：同意；4：不同意】
	OpinionDesc  string    // 处理意见描述

}

// InstUserTaskResult
// @Description: 实例用户任务
type InstUserTaskResult struct {
	ID           int64     // 唯一id
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
	Status       int32     // 实例用户任务状态【1：处理中；2：完成（同意）；3：不通过（不同意）；4：回退；5：终止】
	CreateTime   time.Time // 创建时间
	UpdateTime   time.Time // 更新时间
	HandleTime   time.Time // 处理时间
	OpUserID     string    // 操作用户id
	OpUserName   string    // 操作用户名称
	Opinion      int32     // 任务处理意见【1：未发表；2：同意；3：不同意；4：xxx】
	OpinionDesc  string    // 处理意见描述
}
