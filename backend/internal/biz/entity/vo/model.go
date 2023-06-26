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
	ModelName string `json:"modelName"` // 组名称
}
