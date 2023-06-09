// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameProcessDefNodeUser = "process_def_node_user"

// ProcessDefNodeUser mapped from table <process_def_node_user>
type ProcessDefNodeUser struct {
	ID           int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`    // 唯一id
	ProcessDefID string `gorm:"column:process_def_id;not null" json:"process_def_id"` // 流程定义id
	NodeID       string `gorm:"column:node_id;not null" json:"node_id"`               // 节点id
	Type         int32  `gorm:"column:type;not null;default:1" json:"type"`           // 常用审批人【指定成员：1；发起人自己：2；发起人自选：3：角色：4；部门：5】主管（相对岗位）【直属主管：1；部门主管：2；连续多级主管：3；部门控件对应主管：4】其他【表单人员控件：1；部门控件：2；角色控件：3】
	Strategy     int32  `gorm:"column:strategy;not null;default:1" json:"strategy"`   // 处理人策略【常用审批人：1；主管（相对岗位）：2；其他：3】
	NodeUserName string `gorm:"column:node_user_name;not null" json:"node_user_name"` // 处理人名称
	NodeUserID   string `gorm:"column:node_user_id;not null" json:"node_user_id"`     // 处理人id
	Sort         int32  `gorm:"column:sort;not null;default:1" json:"sort"`           // 处理人顺序;正序排序
	Obj          string `gorm:"column:obj;not null" json:"obj"`                       // 扩展字段，设计中可忽略
	Relative     string `gorm:"column:relative;not null" json:"relative"`             // 相对发起人的直属主管，设计中可忽略
}

// TableName ProcessDefNodeUser's table name
func (*ProcessDefNodeUser) TableName() string {
	return TableNameProcessDefNodeUser
}
