// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameInstTaskDetail = "inst_task_detail"

// InstTaskDetail mapped from table <inst_task_detail>
type InstTaskDetail struct {
	ID             int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`                        // 唯一id
	InstTaskID     string    `gorm:"column:inst_task_id;not null" json:"inst_task_id"`                         // 实例任务id
	ModelID        string    `gorm:"column:model_id;not null" json:"model_id"`                                 // 模板id
	ProcessDefID   string    `gorm:"column:process_def_id;not null" json:"process_def_id"`                     // 流程定义id
	FormDefID      string    `gorm:"column:form_def_id;not null" json:"form_def_id"`                           // 表单定义id
	VersionID      string    `gorm:"column:version_id;not null" json:"version_id"`                             // 版本id
	TaskName       string    `gorm:"column:task_name;not null" json:"task_name"`                               // 实例任务名称
	Status         int32     `gorm:"column:status;not null;default:1" json:"status"`                           // 任务状态【1：创建中(草稿)；2：进行中； 3：终止； 4：完成； 5：挂起；6：回退】
	Remark         string    `gorm:"column:remark;not null" json:"remark"`                                     // 描述
	CreateTime     time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"` // 创建时间
	CreateUserID   string    `gorm:"column:create_user_id;not null" json:"create_user_id"`                     // 创建人id
	CreateUserName string    `gorm:"column:create_user_name;not null" json:"create_user_name"`                 // 创建人名称
	UpdateTime     time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP" json:"update_time"`          // 更新时间
	UpdateUserID   string    `gorm:"column:update_user_id" json:"update_user_id"`                              // 更新人id
	UpdateUserName string    `gorm:"column:update_user_name" json:"update_user_name"`                          // 更新人名称
	StartTime      time.Time `gorm:"column:start_time;not null;default:CURRENT_TIMESTAMP" json:"start_time"`   // 发起时间
	EndTime        time.Time `gorm:"column:end_time;not null;default:CURRENT_TIMESTAMP" json:"end_time"`       // 结束时间
}

// TableName InstTaskDetail's table name
func (*InstTaskDetail) TableName() string {
	return TableNameInstTaskDetail
}
