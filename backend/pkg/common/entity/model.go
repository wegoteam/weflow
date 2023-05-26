package entity

import "time"

type ModelVersionResult struct {
	ID           int64     // 唯一id
	ModelID      string    // 模板id
	ModelTitle   string    // 模板版本标题
	VersionID    string    // 版本id
	ProcessDefID string    // 流程定义id
	FormDefID    string    // 表单定义id
	UseStatus    int32     // 使用状态【1：使用；2：未使用】
	Remark       string    // 描述
	CreateTime   time.Time // 创建时间
	CreateUser   string    // 创建人
	UpdateTime   time.Time // 更新时间
	UpdateUser   string    // 更新人
	NoticeURL    string    // 回调通知推送url
	TitleProps   string    // 标题配置
}
