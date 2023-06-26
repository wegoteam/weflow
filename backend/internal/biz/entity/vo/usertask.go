package vo

// UserTaskQueryVO
// @Description: 用户任务查询VO
type UserTaskQueryVO struct {
	PageSize        int    `json:"pageSize"swaggertype:"integer" example:"30"`       // 每页条数
	PageNum         int    `json:"pageNum"  swaggertype:"integer" example:"1"`       // 页码
	TaskName        string `json:"taskName"  swaggertype:"string" example:""`        // 任务名称
	InstStatus      int    `json:"instStatus"  swaggertype:"integer" example:"0"`    // 任务状态，零值为全部
	ModelID         string `json:"modelID"  swaggertype:"string" example:""`         // 模板ID
	CreateUserID    string `json:"createUserID"  swaggertype:"string" example:""`    // 创建人id
	CreateTimeStart string `json:"createTimeStart"  swaggertype:"string" example:""` // 提交审批时间-开始 yyyy-MM-dd HH:mm:ss:SSS
	CreateTimeEnd   string `json:"createTimeEnd" swaggertype:"string" example:""`    // 提交审批时间-结束 yyyy-MM-dd HH:mm:ss:SSS
	FinishTimeStart string `json:"finishTimeStart"  swaggertype:"string" example:""` // 完成审批时间-开始 yyyy-MM-dd HH:mm:ss:SSS
	FinishTimeEnd   string `json:"finishTimeEnd"  swaggertype:"string" example:""`   // 完成审批时间-结束 yyyy-MM-dd HH:mm:ss:SSS
}

// UserTaskAgreeVO
// @Description: 用户任务同意VO
type UserTaskAgreeVO struct {
	UserTaskID  string         // 用户任务id
	OpinionDesc string         // 操作意见
	Params      map[string]any // 参数
}

// UserTaskDisagreeVO
// @Description: 用户任务不同意VO
type UserTaskDisagreeVO struct {
	UserTaskID  string         // 用户任务id
	OpinionDesc string         // 操作意见
	Params      map[string]any // 参数
}

// UserTaskSaveVO
// @Description: 用户任务保存VO
type UserTaskSaveVO struct {
	UserTaskID  string         // 用户任务id
	OpinionDesc string         // 操作意见
	Params      map[string]any // 参数
}
