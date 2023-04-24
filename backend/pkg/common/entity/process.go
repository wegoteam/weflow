package entity

type ProcessDefModel struct {
	ProcessDefId string                 `json:"processDefId"` // 流程定义ID
	StartNodeId  string                 `json:"startNodeId"`  // 开始节点ID
	NodeModelMap map[string]NodeModelBO `json:"nodeModelMap"` // 节点map
	NodeModels   *[]NodeModelBO         `json:"nodeModels"`   // 节点列表
}
