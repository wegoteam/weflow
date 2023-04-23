// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameProcessDefNode = "process_def_node"

// ProcessDefNode mapped from table <process_def_node>
type ProcessDefNode struct {
	ID             int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`                        // 唯一id
	ProcessDefID   string    `gorm:"column:process_def_id;not null" json:"process_def_id"`                     // 流程定义id
	NodeID         string    `gorm:"column:node_id;not null" json:"node_id"`                                   // 节点id
	NodeType       int32     `gorm:"column:node_type;not null;default:1" json:"node_type"`                     // 节点类型;1：正常节点；2：开始节点；3：结束节点；4：汇聚节点；5：条件节点
	NodeName       string    `gorm:"column:node_name;not null" json:"node_name"`                               // 节点名称
	ForwardMode    int32     `gorm:"column:forward_mode;not null;default:1" json:"forward_mode"`               // 进行模式【1：并行 2：串行】
	CompleteConn   int32     `gorm:"column:complete_conn;not null" json:"complete_conn"`                       // 节点完成条件;通过的人数，0表示所有人通过，节点才算完成
	PermissionMode int32     `gorm:"column:permission_mode;not null;default:1" json:"permission_mode"`         // 权限模式【1：协同 2：知会 3：审批】
	AllowAdd       int32     `gorm:"column:allow_add;not null;default:1" json:"allow_add"`                     // 允许加签【1：不能加签；2：允许加签】
	ProcessMode    int32     `gorm:"column:process_mode;not null;default:1" json:"process_mode"`               // 处理模式【1：人工； 2：自动】
	BusID          string    `gorm:"column:bus_id;not null" json:"bus_id"`                                     // 业务id
	BusType        string    `gorm:"column:bus_type;not null" json:"bus_type"`                                 // 业务类型
	TimeLimit      int32     `gorm:"column:time_limit;not null" json:"time_limit"`                             // 处理期限时长;单位秒，0表示无期限；
	ConnData       string    `gorm:"column:conn_data;not null" json:"conn_data"`                               // 条件表达式;条件节点才有条件表达式
	FormPerData    string    `gorm:"column:form_per_data;not null" json:"form_per_data"`                       // 表单权限数据;节点表单权限配置，json格式
	Remark         string    `gorm:"column:remark;not null" json:"remark"`                                     // 节点描述
	CreateTime     time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"` // 创建时间
	CreateUser     string    `gorm:"column:create_user;not null" json:"create_user"`                           // 创建人
	UpdateTime     time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP" json:"update_time"`          // 更新时间
	UpdateUser     string    `gorm:"column:update_user" json:"update_user"`                                    // 更新人
}

// TableName ProcessDefNode's table name
func (*ProcessDefNode) TableName() string {
	return TableNameProcessDefNode
}