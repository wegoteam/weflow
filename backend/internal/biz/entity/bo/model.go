package bo

import (
	"time"
)

// ModelDetailResult
// @Description: 模型详情
type ModelDetailResult struct {
	ID           int64  `json:"id"`           // 唯一id
	ModelID      string `json:"modelID"`      // 模板id
	ModelTitle   string `json:"modelTitle"`   // 模板标题
	ProcessDefID string `json:"processDefID"` // 流程定义id
	FormDefID    string `json:"formDefID"`    // 表单定义id
	ModelGroupID string `json:"modelGroupID"` // 模版组id
	IconURL      string `json:"iconURL"`      // icon图标地址
	Status       int32  `json:"status"`       // 模板状态【1：草稿；2：发布；3：停用】默认草稿
	Remark       string `json:"remark"`       // 描述
	CreateTime   string `json:"createTime"`   // 创建时间 yyyy-MM-dd HH:mm:ss:SSS
	CreateUser   string `json:"createUser"`   // 创建人
	UpdateTime   string `json:"updateTime"`   // 更新时间 yyyy-MM-dd HH:mm:ss:SSS
	UpdateUser   string `json:"updateUser"`   // 更新人
}

// ModelGroupResult
// @Description: 模型组
type ModelGroupResult struct {
	ID         int64  `json:"id"`         // 唯一id
	GroupID    string `json:"groupID"`    // 组id
	GroupName  string `json:"groupName"`  // 组名称
	Remark     string `json:"remark"`     // 描述
	CreateUser string `json:"createUser"` // 创建人
	UpdateUser string `json:"updateUser"` // 更新人
	CreateTime string `json:"createTime"` // 创建时间 yyyy-MM-dd HH:mm:ss:SSS
	UpdateTime string `json:"updateTime"` // 更新时间 yyyy-MM-dd HH:mm:ss:SSS
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

// ModelSaveResult
// @Description: 模型保存结果
type ModelSaveResult struct {
	ModelID string `json:"modelID"` // 模板id
}

// ModelVersionResult
// @Description: 模型版本
type ModelVersionResult struct {
	ID           int64  `json:"id"`           // 唯一id
	ModelID      string `json:"modelID"`      // 模板id
	ModelTitle   string `json:"modelTitle"`   // 模板版本标题
	VersionID    string `json:"versionID"`    // 版本id
	ProcessDefID string `json:"processDefID"` // 流程定义id
	FormDefID    string `json:"formDefID"`    // 表单定义id
	UseStatus    int32  `json:"useStatus"`    // 使用状态【1：使用；2：未使用】
	Remark       string `json:"remark"`       // 描述
	CreateTime   string `json:"createTime"`   // 创建时间
	CreateUser   string `json:"createUser"`   // 创建人
	UpdateTime   string `json:"updateTime"`   // 更新时间
	UpdateUser   string `json:"updateUser"`   // 更新人
	NoticeURL    string `json:"noticeURL"`    // 回调通知推送url
	TitleProps   string `json:"titleProps"`   // 标题配置
}

// ModelAndVersionInfoResult
// @Description: 模型及版本信息
type ModelAndVersionInfoResult struct {
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
	FlowContent  string    `json:"flowContent" ` // 流程内容
	FormContent  string    `json:"formContent" ` // 表单内容
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
	CreateTime string              // 创建时间
	UpdateTime string              // 更新时间
	Models     []ModelDetailResult // 模型列表
}
