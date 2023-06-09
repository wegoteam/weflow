// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameModelVersion = "model_version"

// ModelVersion mapped from table <model_version>
type ModelVersion struct {
	ID           int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`                        // 唯一id
	ModelID      string    `gorm:"column:model_id;not null" json:"model_id"`                                 // 模板id
	ModelTitle   string    `gorm:"column:model_title;not null" json:"model_title"`                           // 模板版本标题
	VersionID    string    `gorm:"column:version_id;not null" json:"version_id"`                             // 版本id
	ProcessDefID string    `gorm:"column:process_def_id;not null" json:"process_def_id"`                     // 流程定义id
	FormDefID    string    `gorm:"column:form_def_id;not null" json:"form_def_id"`                           // 表单定义id
	UseStatus    int32     `gorm:"column:use_status;not null;default:1" json:"use_status"`                   // 使用状态【1：使用；2：未使用】
	Remark       string    `gorm:"column:remark;not null" json:"remark"`                                     // 描述
	CreateTime   time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"` // 创建时间
	CreateUser   string    `gorm:"column:create_user;not null" json:"create_user"`                           // 创建人
	UpdateTime   time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP" json:"update_time"`          // 更新时间
	UpdateUser   string    `gorm:"column:update_user" json:"update_user"`                                    // 更新人
	NoticeURL    string    `gorm:"column:notice_url;not null" json:"notice_url"`                             // 回调通知推送url
	TitleProps   string    `gorm:"column:title_props;not null" json:"title_props"`                           // 标题配置
}

// TableName ModelVersion's table name
func (*ModelVersion) TableName() string {
	return TableNameModelVersion
}
