// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameFlowDefInfo = "flow_def_info"

// FlowDefInfo mapped from table <flow_def_info>
type FlowDefInfo struct {
	ID          int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`                        // 唯一id
	FlowDefID   string    `gorm:"column:flow_def_id;not null" json:"flow_def_id"`                           // 流程定义id
	FlowDefName string    `gorm:"column:flow_def_name;not null" json:"flow_def_name"`                       // 流程定义名称
	Status      int32     `gorm:"column:status;not null;default:1" json:"status"`                           // 状态【1：草稿；2：发布可用；3：停用】
	Remark      string    `gorm:"column:remark;not null" json:"remark"`                                     // 描述
	StructData  string    `gorm:"column:struct_data" json:"struct_data"`                                    // 流程结构化数据
	CreateTime  time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"` // 创建时间
	CreateUser  string    `gorm:"column:create_user;not null" json:"create_user"`                           // 创建人
	UpdateTime  time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP" json:"update_time"`          // 更新时间
	UpdateUser  string    `gorm:"column:update_user" json:"update_user"`                                    // 更新人
}

// TableName FlowDefInfo's table name
func (*FlowDefInfo) TableName() string {
	return TableNameFlowDefInfo
}
