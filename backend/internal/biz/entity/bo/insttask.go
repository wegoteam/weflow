package bo

import "time"

// InstTaskResult
// @Description: 实例任务返回结果
type InstTaskResult struct {
	ID             int64     `json:"id"`             // 唯一id
	InstTaskID     string    `json:"instTaskID"`     // 实例任务id
	ModelID        string    `json:"modelID"`        // 模板id
	ProcessDefID   string    `json:"processDefID"`   // 流程定义id
	FormDefID      string    `json:"formDefID"`      // 表单定义id
	VersionID      string    `json:"versionID"`      // 版本id
	TaskName       string    `json:"taskName"`       // 实例任务名称
	Status         int32     `json:"status"`         // 任务状态【1：创建中(草稿)；2：进行中； 3：终止； 4：完成； 5：挂起；6：回退】
	Remark         string    `json:"remark"`         // 描述
	CreateTime     time.Time `json:"createTime"`     // 创建时间 yyyy-MM-dd HH:mm:ss:SSS
	CreateUserID   string    `json:"createUserID"`   // 创建人id
	CreateUserName string    `json:"createUserName"` // 创建人名称
	UpdateTime     time.Time `json:"updateTime"`     // 更新时间 yyyy-MM-dd HH:mm:ss:SSS
	UpdateUserID   string    `json:"updateUserID"`   // 更新人id
	UpdateUserName string    `json:"updateUserName"` // 更新人名称
	StartTime      time.Time `json:"startTime"`      // 发起时间 yyyy-MM-dd HH:mm:ss:SSS
	EndTime        time.Time `json:"endTime"`        // 结束时间 yyyy-MM-dd HH:mm:ss:SSS
}
