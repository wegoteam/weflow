package entity

// ProcessDefModel
// @Description: 流程定义模型
type ProcessDefModel struct {
	ProcessDefId string                 `json:"processDefId"` // 流程定义ID
	StartNodeId  string                 `json:"startNodeId"`  // 开始节点ID
	NodeModelMap map[string]NodeModelBO `json:"nodeModelMap"` // 节点map
	NodeModels   *[]NodeModelBO         `json:"nodeModels"`   // 节点列表
}

// ProcessConditions
// @Description: 流程条件
type ProcessConditions struct {
	Key           string //条件来源
	keyType       string //条件来源类型【字符串：string；整形数值：int；浮点型数值：float；数组：array；表格：table】
	Val           string //条件目标
	valType       string //条件目标类型【字符串：string；整形数值：int；浮点型数值：float；数组：array】
	operation     string //操作符 条件操作【条件语法：比较操作：等于，不等于，大于，大于等于，小于，小于等于；集合操作：包含，完全等于，包含任意；字符操作：包含，完全等于】
	OperationType string //操作类型【比较操作：comparison，逻辑操作：logical，集合操作：array，字符操作：string，其他：other】
	Mode          int    //条件分类【表单条件：0；人员条件：1】
}

type ConditionResult struct {
}
