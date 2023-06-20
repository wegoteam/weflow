package entity

// UserTaskQueryBO
// @Description: 用户任务查询BO
type UserTaskQueryBO struct {
	PageSize        int    `json:"pageSize"`        // 每页条数
	PageNum         int    `json:"pageNum"`         // 页码
	TaskName        string `json:"taskName"`        // 任务名称
	InstStatus      int    `json:"instStatus"`      // 任务状态
	ModelId         string `json:"groupId"`         // 组名称
	CreateUserId    string `json:"createUserId"`    // 创建人id
	CreateTimeStart string `json:"createTimeStart"` // 提交审批时间-开始
	CreateTimeEnd   string `json:"createTimeEnd"`   // 提交审批时间-结束
	FinishTimeStart string `json:"finishTimeStart"` // 完成审批时间-开始
	FinishTimeEnd   string `json:"finishTimeEnd"`   // 完成审批时间-结束
}
