package entity

import "time"

// ModelVersionResult
// @Description: 模型版本
type ModelVersionResult struct {
	ID           int64     `json:"id"`           // 唯一id
	ModelID      string    `json:"modelID"`      // 模板id
	ModelTitle   string    `json:"modelTitle"`   // 模板版本标题
	VersionID    string    `json:"versionID"`    // 版本id
	ProcessDefID string    `json:"processDefID"` // 流程定义id
	FormDefID    string    `json:"formDefID"`    // 表单定义id
	UseStatus    int32     `json:"useStatus"`    // 使用状态【1：使用；2：未使用】
	Remark       string    `json:"remark"`       // 描述
	CreateTime   time.Time `json:"createTime"`   // 创建时间
	CreateUser   string    `json:"createUser"`   // 创建人
	UpdateTime   time.Time `json:"updateTime"`   // 更新时间
	UpdateUser   string    `json:"updateUser"`   // 更新人
	NoticeURL    string    `json:"noticeURL"`    // 回调通知推送url
	TitleProps   string    `json:"titleProps"`   // 标题配置
}

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
	CreateTime   time.Time `json:"createTime"`   // 创建时间
	CreateUser   string    `json:"createUser"`   // 创建人
	UpdateTime   time.Time `json:"updateTime"`   // 更新时间
	UpdateUser   string    `json:"updateUser"`   // 更新人
}

// ModelGroupResult
// @Description: 模型组
type ModelGroupResult struct {
	ID         int64     // 唯一id
	GroupID    string    // 组id
	GroupName  string    // 组名称
	Remark     string    // 描述
	CreateUser string    // 创建人
	UpdateUser string    // 更新人
	CreateTime time.Time // 创建时间
	UpdateTime time.Time // 更新时间
}

// ModelGroupAddBO
// @Description: 模型组添加BO
type ModelGroupAddBO struct {
	GroupName  string    `json:"groupName"`  // 组名称
	Remark     string    `json:"remark"`     // 描述
	CreateUser string    `json:"createUser"` // 创建人
	UpdateUser string    `json:"updateUser"` // 更新人
	CreateTime time.Time `json:"createTime"` // 创建时间
	UpdateTime time.Time `json:"updateTime"` // 更新时间
}

// ModelGroupEditBO
// @Description: 模型组编辑BO
type ModelGroupEditBO struct {
	GroupID    string    `json:"groupID"`    // 组id
	GroupName  string    `json:"groupName"`  // 组名称
	Remark     string    `json:"remark"`     // 描述
	UpdateUser string    `json:"updateUser"` // 更新人
	UpdateTime time.Time `json:"updateTime"` // 更新时间
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

// GroupModelDetailsResult
// @Description: 模型组详情
type GroupModelDetailsResult struct {
	ID         int64               // 唯一id
	GroupID    string              // 组id
	GroupName  string              // 组名称
	Remark     string              // 描述
	CreateUser string              // 创建人
	UpdateUser string              // 更新人
	CreateTime time.Time           // 创建时间
	UpdateTime time.Time           // 更新时间
	Models     []ModelDetailResult // 模型列表
}

// ModelQueryBO
// @Description: 模型查询BO
type ModelQueryBO struct {
	ModelName string `json:"modelName"` // 模型名称
	Status    int    `json:"status"`    // 模板状态【1：草稿；2：发布；3：停用】默认草稿
}

// ModelPageBO
// @Description: 模型分页BO
type ModelPageBO struct {
	ModelName string `json:"modelName" ` // 模型名称
	Status    int    `json:"status"`     // 模板状态【1：草稿；2：发布；3：停用】默认草稿
	PageSize  int    `json:"pageSize"`   // 每页条数
	PageNum   int    `json:"pageNum"`    // 页码
}

// ModelSaveBO
// @Description: 模型保存BO
type ModelSaveBO struct {
	ModelID     string             `json:"modelID" `     // 模型id
	Base        ModelBaseSetup     `json:"base" `        // 基础设置
	FlowContent string             `json:"flowContent" ` // 流程内容
	FormContent string             `json:"formContent" ` // 表单内容
	Advanced    ModelAdvancedSetup `json:"advanced" `    // 高级设置
	UserID      string             `json:"userID" `      // 用户id
	UserName    string             `json:"userName" `    // 用户名称
}

// ModelBaseSetup
// @Description: 模型基础设置
type ModelBaseSetup struct {
	ModelName string `json:"modelName" ` // 模型名称
	IconURL   string `json:"iconURL" `   // 图标地址
	GroupID   string `json:"groupID" `   // 组id
	Remark    string `json:"remark" `    // 描述
}

// ModelAdvancedSetup
// @Description: 模型高级设置
type ModelAdvancedSetup struct {
	TitleType    int    `json:"titleType" `    // 标题类型【1：默认；2：自定义】默认为1
	TitleContent string `json:"titleContent" ` // 标题内容
}
