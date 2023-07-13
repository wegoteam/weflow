package entity

import "time"

// InstNodeTaskResult
// @Description: 节点任务返回结果
type InstNodeTaskResult struct {
	ID             int64     // 唯一id
	InstTaskID     string    // 实例任务id
	NodeTaskID     string    // 节点任务id
	NodeID         string    // 节点id
	ParentID       string    // 父节点id
	NodeModel      int32     // 节点模型【1：开始节点；2：审批节点；3：办理节点；4：抄送节点；5：自定义节点；6：条件节点；7：分支节点；8：汇聚节点；9：结束节点】
	NodeName       string    // 节点名称
	ApproveType    int32     // 审批类型【人工审批：1；自动通过：2；自动拒绝】默认人工审批1
	NoneHandler    int32     // 审批人为空时【自动通过：1；自动转交管理员：2；指定审批人：3】默认自动通过1
	AppointHandler string    // 审批人为空时指定审批人ID
	HandleMode     int32     // 审批方式【依次审批：1；会签（需要完成人数的审批人同意或拒绝才可完成节点）：2；或签（其中一名审批人同意或拒绝即可）：3】默认会签2
	FinishMode     int32     // 完成人数：依次审批默认0所有人不可选人，会签默认0所有人（可选人大于0），或签默认1一个人（可选人大于0）
	BranchMode     int32     // 分支执行方式【单分支：1；多分支：2】默认多分支2
	DefaultBranch  int32     // 单分支处理需要默认分支，在条件优先级无法处理时候执行默认分支，取值分支下标
	BranchLevel    int32     // 优先级，分支执行方式为多分支处理方式无优先级应为0
	ConditionGroup string    // 条件组前端描述展示条件组
	ConditionExpr  string    // 条件组解析后的表达式
	Remark         string    // 节点描述
	Status         int32     // 任务状态【1：未开始；2：处理中；3：完成；4：回退；5：终止；6：不通过】
	CreateTime     time.Time // 创建时间
	UpdateTime     time.Time // 更新时间
}

// InstNodeTaskFormperResult
// @Description: 节点任务表单权限返回结果
type InstNodeTaskFormperResult struct {
	ID         int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"` // 唯一id
	InstTaskID string `gorm:"column:inst_task_id;not null" json:"inst_task_id"`  // 实例任务id
	NodeTaskID string `gorm:"column:node_task_id;not null" json:"node_task_id"`  // 节点任务id
	NodeID     string `gorm:"column:node_id;not null" json:"node_id"`            // 节点id
	ElemID     string `gorm:"column:elemId;not null" json:"elemId"`              // 表单元素ID
	ElemPID    string `gorm:"column:elemPId;not null" json:"elemPId"`            // 表单元素父ID
	Per        int32  `gorm:"column:per;not null;default:2" json:"per"`          // 表单权限【可编辑：1；只读：2；隐藏：3;必填：4】默认只读2
}
