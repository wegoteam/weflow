package exec

import (
	"github.com/wegoteam/weflow/pkg/common/entity"
	"github.com/wegoteam/weflow/pkg/parser"
)

type ExecFlow struct {
}

func StartInstTask() {
	processDefModel := parser.GetProcessDefModel("1640993392605401001")

	execution := &entity.Execution{}
	execution.ProcessDefModel = processDefModel
	execution.InstTaskId = "1640993392605400001"
	execution.InstTaskName = "测试流程"
	execution.InstTaskStatus = 1

	startNodeId := processDefModel.StartNodeId

	startNode := (*processDefModel.NodeModelMap)[startNodeId]
	Exec(&startNode, execution)
}
