// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameOrganizationInfo = "organization_info"

// OrganizationInfo mapped from table <organization_info>
type OrganizationInfo struct {
	ID         int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`                        // 唯一id
	OrgID      string    `gorm:"column:org_id;not null" json:"org_id"`                                     // 组id
	ParentID   string    `gorm:"column:parent_id;not null" json:"parent_id"`                               // 组父id
	OrgName    string    `gorm:"column:org_name;not null" json:"org_name"`                                 // 组名称
	Status     int32     `gorm:"column:status;not null;default:1" json:"status"`                           // 状态【1：未启用；2：已启用；3：锁定；】
	Remark     string    `gorm:"column:remark;not null" json:"remark"`                                     // 描述
	CreateUser string    `gorm:"column:create_user;not null" json:"create_user"`                           // 创建人
	UpdateUser string    `gorm:"column:update_user;not null" json:"update_user"`                           // 更新人
	CreateTime time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP" json:"update_time"`          // 更新时间
}

// TableName OrganizationInfo's table name
func (*OrganizationInfo) TableName() string {
	return TableNameOrganizationInfo
}
