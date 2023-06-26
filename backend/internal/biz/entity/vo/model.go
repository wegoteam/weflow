package vo

// ModelGroupAddVO
// @Description: 模型组添加VO
type ModelGroupAddVO struct {
	GroupName string `json:"groupName"` // 组名称
	Remark    string `json:"remark"`    // 描述
}

// ModelGroupEditVO
// @Description: 模型组编辑VO
type ModelGroupEditVO struct {
	GroupID   string `json:"groupID"`   // 组id
	GroupName string `json:"groupName"` // 组名称
	Remark    string `json:"remark"`    // 描述
}

// ModelGroupDelVO
// @Description: 模型组删除VO
type ModelGroupDelVO struct {
	GroupID string `json:"groupID"` // 组id
}

// GroupModelQueryVO
// @Description: 模型组查询VO
type GroupModelQueryVO struct {
	ModelName string `json:"modelName"` // 模型名称
}

// ModelQueryVO
// @Description: 模型查询VO
type ModelQueryVO struct {
	ModelName string `json:"modelName" swaggertype:"string" example:""` // 模型名称
	Status    int    `json:"status" swaggertype:"integer" example:"0"`  // 模板状态【1：草稿；2：发布；3：停用】默认草稿
}

// ModelPageVO
// @Description: 模型分页VO
type ModelPageVO struct {
	ModelName string `json:"modelName" swaggertype:"string" example:""`  // 模型名称
	Status    int    `json:"status" swaggertype:"integer" example:"0"`   // 模板状态【1：草稿；2：发布；3：停用】默认草稿
	PageSize  int    `json:"pageSize"swaggertype:"integer" example:"30"` // 每页条数
	PageNum   int    `json:"pageNum"  swaggertype:"integer" example:"1"` // 页码
}
