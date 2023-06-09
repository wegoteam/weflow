// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameInstTaskParamAttr = "inst_task_param_attr"

// InstTaskParamAttr mapped from table <inst_task_param_attr>
type InstTaskParamAttr struct {
	ID         int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`                        // 唯一id
	InstTaskID string    `gorm:"column:inst_task_id;not null" json:"inst_task_id"`                         // 实例任务id
	ParamID    string    `gorm:"column:param_id;not null" json:"param_id"`                                 // 参数id
	ParamName  string    `gorm:"column:param_name;not null" json:"param_name"`                             // 参数名称
	ParamAttr  string    `gorm:"column:param_attr;not null" json:"param_attr"`                             // 参数值
	CreateTime time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP" json:"update_time"`          // 更新时间
}

// TableName InstTaskParamAttr's table name
func (*InstTaskParamAttr) TableName() string {
	return TableNameInstTaskParamAttr
}
