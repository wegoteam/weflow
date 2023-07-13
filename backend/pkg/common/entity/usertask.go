package entity

import "time"

// UserTaskQueryBO
// @Description: 用户任务查询BO
type UserTaskQueryBO struct {
	UserID          string `json:"userID"`          // 用户id
	PageSize        int    `json:"pageSize"`        // 每页条数
	PageNum         int    `json:"pageNum"`         // 页码
	TaskName        string `json:"taskName"`        // 任务名称
	InstStatus      int    `json:"instStatus"`      // 任务状态
	ModelID         string `json:"modelID"`         // 模板名称
	CreateUserID    string `json:"createUserID"`    // 创建人id
	CreateTimeStart string `json:"createTimeStart"` // 提交审批时间-开始
	CreateTimeEnd   string `json:"createTimeEnd"`   // 提交审批时间-结束
	FinishTimeStart string `json:"finishTimeStart"` // 完成审批时间-开始
	FinishTimeEnd   string `json:"finishTimeEnd"`   // 完成审批时间-结束
}

// InstUserTaskResult
// @Description: 实例用户任务
type InstUserTaskResult struct {
	ID           int64     `json:"id"`           // 唯一id
	InstTaskID   string    `json:"instTaskID"`   // 实例任务id
	NodeTaskID   string    `json:"nodeTaskID"`   // 节点任务id
	NodeID       string    `json:"nodeID"`       // 节点任务id
	UserTaskID   string    `json:"userTaskID"`   // 处理人任务id
	Type         int32     `json:"type"`         // 常用审批人【指定成员：1；发起人自己：2；发起人自选：3：角色：4；部门：5】主管（相对岗位）【直属主管：1；部门主管：2；连续多级主管：3；部门控件对应主管：4】其他【表单人员控件：1；部门控件：2；角色控件：3】
	Strategy     int32     `json:"strategy"`     // 处理人策略【常用审批人：1；主管（相对岗位）：2；其他：3】
	NodeUserName string    `json:"nodeUserName"` // 处理人名称
	NodeUserID   string    `json:"nodeUserID"`   // 处理人id
	Sort         int32     `json:"sort"`         // 处理人顺序;正序排序
	Obj          string    `json:"obj"`          // 扩展字段，设计中可忽略
	Relative     string    `json:"relative"`     // 相对发起人的直属主管，设计中可忽略
	Status       int32     `json:"status"`       // 实例用户任务状态【1：处理中；2：完成（同意）；3：不通过（不同意）；4：回退；5：终止】
	CreateTime   time.Time `json:"createTime"`   // 创建时间
	UpdateTime   time.Time `json:"updateTime"`   // 更新时间
	HandleTime   time.Time `json:"handleTime"`   // 处理时间
	OpUserID     string    `json:"opUserID"`     // 操作用户id
	OpUserName   string    `json:"opUserName"`   // 操作用户名称
	Opinion      int32     `json:"opinion"`      // 任务处理意见【1：未发表；2：同意；3：不同意；4：xxx】
	OpinionDesc  string    `json:"opinionDesc"`  // 处理意见描述
}

// InstUserTaskOpinionResult
// @Description: 实例用户任务意见
type InstUserTaskOpinionResult struct {
	ID          int64     // 唯一id
	InstTaskID  string    // 实例任务id
	NodeTaskID  string    // 节点任务id
	UserTaskID  string    // 用户任务id
	NodeID      string    // 节点id
	OpinionID   string    // 意见id
	Opinion     int32     // 处理意见【1：未处理；2：同意；3：不同意；4：回退；5：终止；6：撤回】
	OpinionDesc string    // 处理意见描述
	OpUserID    string    // 操作用户id
	OpUserName  string    // 操作用户名称
	CreateTime  time.Time // 创建时间
	UpdateTime  time.Time // 更新时间
	OpinionTime time.Time // 发表意见时间
}
