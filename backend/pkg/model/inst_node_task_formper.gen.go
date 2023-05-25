// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameInstNodeTaskFormper = "inst_node_task_formper"

// InstNodeTaskFormper mapped from table <inst_node_task_formper>
type InstNodeTaskFormper struct {
	ID         int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"` // 唯一id
	InstTaskID string `gorm:"column:inst_task_id;not null" json:"inst_task_id"`  // 实例任务id
	NodeTaskID string `gorm:"column:node_task_id;not null" json:"node_task_id"`  // 节点任务id
	NodeID     string `gorm:"column:node_id;not null" json:"node_id"`            // 节点id
	ElemID     string `gorm:"column:elemId;not null" json:"elemId"`              // 表单元素ID
	ElemPID    string `gorm:"column:elemPId;not null" json:"elemPId"`            // 表单元素父ID
	Per        int32  `gorm:"column:per;not null;default:2" json:"per"`          // 表单权限【可编辑：1；只读：2；隐藏：3;必填：4】默认只读2
}

// TableName InstNodeTaskFormper's table name
func (*InstNodeTaskFormper) TableName() string {
	return TableNameInstNodeTaskFormper
}
