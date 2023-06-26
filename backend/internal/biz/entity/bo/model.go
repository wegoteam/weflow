package bo

import "time"

// ModelDetailResult
// @Description: 模型详情
type ModelDetailResult struct {
	ID           int64     `json:"id"`           // 唯一id
	ModelID      string    `json:"modelID"`      // 模板id
	ModelTitle   string    `json:"modelTitle"`   // 模板标题
	ProcessDefID string    `json:"processDefID"` // 流程定义id
	FormDefID    string    `json:"formDefID"`    // 表单定义id
	ModelGroupID string    `json:"modelGroupID"` // 模版组id
	IconURL      string    `json:"iconURL"`      // icon图标地址
	Status       int32     `json:"status"`       // 模板状态【1：草稿；2：发布；3：停用】默认草稿
	Remark       string    `json:"remark"`       // 描述
	CreateTime   time.Time `json:"createTime"`   // 创建时间 yyyy-MM-dd HH:mm:ss:SSS
	CreateUser   string    `json:"createUser"`   // 创建人
	UpdateTime   time.Time `json:"updateTime"`   // 更新时间 yyyy-MM-dd HH:mm:ss:SSS
	UpdateUser   string    `json:"updateUser"`   // 更新人
}

// ModelGroupResult
// @Description: 模型组
type ModelGroupResult struct {
	ID         int64     `json:"id"`         // 唯一id
	GroupID    string    `json:"groupID"`    // 组id
	GroupName  string    `json:"groupName"`  // 组名称
	Remark     string    `json:"remark"`     // 描述
	CreateUser string    `json:"createUser"` // 创建人
	UpdateUser string    `json:"updateUser"` // 更新人
	CreateTime time.Time `json:"createTime"` // 创建时间 yyyy-MM-dd HH:mm:ss:SSS
	UpdateTime time.Time `json:"updateTime"` // 更新时间 yyyy-MM-dd HH:mm:ss:SSS
}

// ModelGroupAddBO
// @Description: 模型组添加BO
type ModelGroupAddBO struct {
	GroupName string `json:"groupName"` // 组名称
	Remark    string `json:"remark"`    // 描述
	OpUser    string `json:"opUser"`    // 用户ID
}

// ModelGroupEditBO
// @Description: 模型组编辑BO
type ModelGroupEditBO struct {
	GroupID   string `json:"groupID"`   // 组id
	GroupName string `json:"groupName"` // 组名称
	Remark    string `json:"remark"`    // 描述
	OpUser    string `json:"opUser"`    // 用户ID
}

// ModelGroupDelBO
// @Description: 模型组删除BO
type ModelGroupDelBO struct {
	GroupID string `json:"groupID"` // 组id
}

// GroupModelQueryBO
// @Description: 模型组查询BO
type GroupModelQueryBO struct {
	ModelName string `json:"modelName"` // 组名称
}
