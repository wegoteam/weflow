// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameInstTaskOpLog = "inst_task_op_log"

// InstTaskOpLog mapped from table <inst_task_op_log>
type InstTaskOpLog struct {
	ID         int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`                        // 唯一id
	InstTaskID string    `gorm:"column:inst_task_id;not null" json:"inst_task_id"`                         // 实例任务id
	NodeID     string    `gorm:"column:node_id;not null" json:"node_id"`                                   // 节点任务id
	NodeName   string    `gorm:"column:node_name;not null" json:"node_name"`                               // 节点名称
	CreateTime time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP" json:"update_time"`          // 更新时间
	Type       int32     `gorm:"column:type;not null;default:1" json:"type"`                               // 类型【1：节点；2：任务；3：其他】
	Remark     string    `gorm:"column:remark;not null" json:"remark"`                                     // 描述
}

// TableName InstTaskOpLog's table name
func (*InstTaskOpLog) TableName() string {
	return TableNameInstTaskOpLog
}
