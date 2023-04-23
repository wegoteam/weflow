// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameInstUserTask = "inst_user_task"

// InstUserTask mapped from table <inst_user_task>
type InstUserTask struct {
	ID           int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`                        // 唯一id
	InstTaskID   string    `gorm:"column:inst_task_id;not null" json:"inst_task_id"`                         // 实例任务id
	NodeTaskID   string    `gorm:"column:node_task_id;not null" json:"node_task_id"`                         // 节点任务id
	NodeID       string    `gorm:"column:node_id;not null" json:"node_id"`                                   // 节点任务id
	UserTaskID   string    `gorm:"column:user_task_id;not null" json:"user_task_id"`                         // 处理人任务id
	NodeUserID   string    `gorm:"column:node_user_id;not null" json:"node_user_id"`                         // 节点处理人id
	NodeUserName string    `gorm:"column:node_user_name;not null" json:"node_user_name"`                     // 处理人名称
	NodeUserType int32     `gorm:"column:node_user_type;not null;default:1" json:"node_user_type"`           // 处理人类型【1：操作员；2：部门；3：相对岗位；4：表单控件；5：角色；6：岗位；7：组织；8：自定义】
	OpOrigin     int32     `gorm:"column:op_origin;not null;default:1" json:"op_origin"`                     // 操作来源【1：正常；2：加签】
	TimeLimit    int64     `gorm:"column:time_limit;not null" json:"time_limit"`                             // 处理期限;格式：yyyymmddhhmm 可直接指定到期限的具体时间，期限支持到分钟； 0表示无期限
	Status       int32     `gorm:"column:status;not null;default:1" json:"status"`                           // 任务状态【1：处理中；2：完成；3：回退；4：终止】
	CreateTime   time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"` // 创建时间
	UpdateTime   time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP" json:"update_time"`          // 更新时间
	HandleTime   time.Time `gorm:"column:handle_time;default:CURRENT_TIMESTAMP" json:"handle_time"`          // 处理时间
	OpUserID     string    `gorm:"column:op_user_id;not null" json:"op_user_id"`                             // 操作用户id
	OpUserName   string    `gorm:"column:op_user_name;not null" json:"op_user_name"`                         // 操作用户名称
	HandlerSort  int32     `gorm:"column:handler_sort;not null;default:1" json:"handler_sort"`               // 处理人排序;处理人当前的处理排序
	Opinion      int32     `gorm:"column:opinion;not null;default:1" json:"opinion"`                         // 处理意见【1：未发表；2：已阅；3：同意；4：不同意】
	OpinionDesc  string    `gorm:"column:opinion_desc;not null" json:"opinion_desc"`                         // 处理意见描述
}

// TableName InstUserTask's table name
func (*InstUserTask) TableName() string {
	return TableNameInstUserTask
}