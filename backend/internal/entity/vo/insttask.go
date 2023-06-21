package vo

// InstTaskQueryVO
// @Description: 实例任务查询vo
type InstTaskQueryVO struct {
	PageSize        int    `json:"pageSize"query:"pageSize" swaggertype:"integer" example:"30"`                    // 每页条数
	PageNum         int    `json:"pageNum"  query:"pageNum" swaggertype:"integer" example:"1"`                     // 页码
	TaskName        string `json:"taskName"  query:"taskName" swaggertype:"string" example:"任务名称"`                 // 任务名称
	InstStatus      int    `json:"instStatus"  query:"instStatus" swaggertype:"integer" example:"1"`               // 任务状态
	ModelId         string `json:"modelId"  query:"modelId" swaggertype:"string" example:"模板ID"`                   // 模板ID
	CreateTimeStart string `json:"createTimeStart"  query:"createTimeStart" swaggertype:"string" example:"提交审批时间"` // 提交审批时间-开始 格式2021-01-28 13:14:15
	CreateTimeEnd   string `json:"createTimeEnd"  query:"createTimeEnd" swaggertype:"string" example:"提交审批时间"`     // 提交审批时间-结束
	FinishTimeStart string `json:"finishTimeStart"  query:"finishTimeStart" swaggertype:"string" example:"完成审批时间"` // 完成审批时间-开始
	FinishTimeEnd   string `json:"finishTimeEnd"  query:"finishTimeEnd" swaggertype:"string" example:"完成审批时间"`     // 完成审批时间-结束
}
