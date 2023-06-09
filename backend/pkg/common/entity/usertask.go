package entity

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
