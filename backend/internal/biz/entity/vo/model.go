package vo

// ModelGroupAddVO
// @Description: 模型组添加VO
type ModelGroupAddVO struct {
	GroupName string `json:"groupName" swaggertype:"string" example:"" vd:"len($)>1; msg:'组名称不能为空'"` // 组名称
	Remark    string `json:"remark" swaggertype:"string" example:""`                                 // 描述
}

// ModelGroupEditVO
// @Description: 模型组编辑VO
type ModelGroupEditVO struct {
	GroupID   string `json:"groupID" swaggertype:"string" example:"" vd:"len($)>1; msg:'组id不能为空'"`   // 组id
	GroupName string `json:"groupName" swaggertype:"string" example:"" vd:"len($)>1; msg:'组名称不能为空'"` // 组名称
	Remark    string `json:"remark" swaggertype:"string" example:""`                                 // 描述
}

// ModelGroupDelVO
// @Description: 模型组删除VO
type ModelGroupDelVO struct {
	GroupID string `json:"groupID" swaggertype:"string" example:"" example:"" vd:"len($)>1; msg:'组id不能为空'"` // 组id
}

// GroupModelQueryVO
// @Description: 模型组查询VO
type GroupModelQueryVO struct {
	ModelName string `json:"modelName" swaggertype:"string" example:""` // 模型名称
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

// ModelSaveVO
// @Description: 模型添加VO
type ModelSaveVO struct {
	ModelID     string             `json:"modelID" swaggertype:"string" example:""`     // 模型id
	Base        ModelBaseSetup     `json:"base"`                                        // 基础设置
	FlowContent string             `json:"flowContent" swaggertype:"string" example:""` // 流程内容
	FormContent string             `json:"formContent" swaggertype:"string" example:""` // 表单内容
	Advanced    ModelAdvancedSetup `json:"advanced"`                                    // 高级设置
}

// ModelBaseSetup
// @Description: 模型基础设置
type ModelBaseSetup struct {
	ModelName string `json:"modelName" swaggertype:"string" example:"" vd:"len($)>1; msg:'模板名称不能为空'"` // 模型名称
	IconURL   string `json:"iconURL" swaggertype:"string" example:""`                                 // 图标地址
	GroupID   string `json:"groupID" swaggertype:"string" example:"" vd:"len($)>1; msg:'组id不能为空'"`    // 组id
	Remark    string `json:"remark" swaggertype:"string" example:""`                                  // 描述
}

// ModelAdvancedSetup
// @Description: 模型高级设置
type ModelAdvancedSetup struct {
	TitleType    int    `json:"titleType" swaggertype:"integer" example:"0"`  // 标题类型【1：默认；2：自定义】默认为1
	TitleContent string `json:"titleContent" swaggertype:"string" example:""` // 标题内容
}

// ModelInvalidVO
// @Description: 模型失效VO
type ModelInvalidVO struct {
	ModelID string `json:"modelID" swaggertype:"string" example:"" vd:"len($)>1; msg:'模板id不能为空'"` // 模型id
}

// ModelVersionQueryVO
// @Description: 模型版本查询VO
type ModelVersionQueryVO struct {
	ModelID string `json:"modelID" query:"modelID" swaggertype:"string" example:""` // 模型id
}

// ReleaseModelVersionVO
// @Description: 发布模型版本VO
type ReleaseModelVersionVO struct {
	VersionID string `json:"versionID" swaggertype:"string" example:"" vd:"len($)>1; msg:'版本id不能为空'"` // 版本id
}

// ModelAndVersionQueryVO
// @Description: 模型详情查询VO
type ModelAndVersionQueryVO struct {
	ModelID   string `json:"modelID" query:"modelID" swaggertype:"string" example:""`     // 模型id
	VersionID string `json:"versionID" query:"versionID" swaggertype:"string" example:""` // 版本id
}
