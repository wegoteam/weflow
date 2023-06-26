package vo

// InstTaskQueryVO
// @Description: 实例任务查询vo
type InstTaskQueryVO struct {
	PageSize        int    `json:"pageSize"swaggertype:"integer" example:"30"`       // 每页条数
	PageNum         int    `json:"pageNum"  swaggertype:"integer" example:"1"`       // 页码
	TaskName        string `json:"taskName"  swaggertype:"string" example:""`        // 任务名称
	InstStatus      int    `json:"instStatus"  swaggertype:"integer" example:"0"`    // 任务状态，零值为全部
	ModelID         string `json:"modelID"  swaggertype:"string" example:""`         // 模板ID
	CreateTimeStart string `json:"createTimeStart"  swaggertype:"string" example:""` // 提交审批时间-开始 yyyy-MM-dd HH:mm:ss:SSS
	CreateTimeEnd   string `json:"createTimeEnd" swaggertype:"string" example:""`    // 提交审批时间-结束 yyyy-MM-dd HH:mm:ss:SSS
	FinishTimeStart string `json:"finishTimeStart"  swaggertype:"string" example:""` // 完成审批时间-开始 yyyy-MM-dd HH:mm:ss:SSS
	FinishTimeEnd   string `json:"finishTimeEnd"  swaggertype:"string" example:""`   // 完成审批时间-结束 yyyy-MM-dd HH:mm:ss:SSS
}

// InstTaskStartVO
// @Description: 实例任务发起参数
type InstTaskStartVO struct {
	ModelID string         // 模板ID
	Params  map[string]any // 参数
}

// InstTaskStopVO
// @Description: 实例任务停止参数
type InstTaskStopVO struct {
	InstTaskID  string // 实例任务ID
	OpinionDesc string // 操作意见
}

// InstTaskSuspendVO
// @Description: 实例任务挂起参数
type InstTaskSuspendVO struct {
	InstTaskID  string // 实例任务ID
	OpinionDesc string // 操作意见
}

// InstTaskSesumeVO
// @Description: 实例任务恢复参数
type InstTaskSesumeVO struct {
	InstTaskID  string // 实例任务ID
	OpinionDesc string // 操作意见
}

// InstTaskDeleteVO
// @Description: 实例任务删除参数
type InstTaskDeleteVO struct {
	InstTaskID  string // 实例任务ID
	OpinionDesc string // 操作意见
}
