package parser

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"testing"
)

func Test(t *testing.T) {
	FlowParserServiceImpl.Parser("")
}

func TestParser(t *testing.T) {
	var data = "[{\"nodeModel\":1,\"nodeName\":\"发起人\",\"nodeId\":\"1640993392605401088\",\"completeConn\":\"0\",\"forwardMode\":\"1\",\"timeLimit\":0,\"perData\":\"\",\"handlerList\":\"\",\"msgConfigList\":\"\"},{\"nodeModel\":2,\"nodeName\":\"审批\",\"nodeId\":\"1640993449224310784\",\"completeConn\":\"0\",\"forwardMode\":\"1\",\"allowAdd\":\"1\",\"notifyState\":\"1\",\"timeLimit\":0,\"perData\":\"{\\\\\\\"1640993433239818240\\\\\\\":\\\\\\\"2\\\\\\\",\\\\\\\"1640993434883985408\\\\\\\":\\\\\\\"2\\\\\\\"}\",\"handlerList\":\"[{\\\\\\\"handlerId\\\\\\\":\\\\\\\"547\\\\\\\",\\\\\\\"handlerName\\\\\\\":\\\\\\\"xuch01\\\\\\\",\\\\\\\"handlerType\\\\\\\":1,\\\\\\\"handlerSort\\\\\\\":1}]\",\"msgConfigList\":\"\"},{\"nodeModel\":7,\"nodeId\":\"1640993508049424384\",\"children\":[[{\"nodeModel\":6,\"nodeName\":\"条件1\",\"nodeId\":\"1640993508049424385\",\"connData\":\"\",\"conditions\":\"\",\"parentId\":\"1640993508049424384\"},{\"nodeModel\":2,\"nodeName\":\"审批\",\"nodeId\":\"1640993526328201216\",\"parentId\":\"1640993508049424384\",\"completeConn\":\"0\",\"forwardMode\":\"1\",\"allowAdd\":\"1\",\"notifyState\":\"1\",\"timeLimit\":0,\"perData\":\"{\\\\\\\"1640993433239818240\\\\\\\":\\\\\\\"2\\\\\\\",\\\\\\\"1640993434883985408\\\\\\\":\\\\\\\"2\\\\\\\"}\",\"handlerList\":\"[{\\\\\\\"handlerId\\\\\\\":\\\\\\\"547\\\\\\\",\\\\\\\"handlerName\\\\\\\":\\\\\\\"xuch01\\\\\\\",\\\\\\\"handlerType\\\\\\\":1,\\\\\\\"handlerSort\\\\\\\":1}]\",\"msgConfigList\":\"\"}],[{\"nodeModel\":6,\"nodeName\":\"条件2\",\"nodeId\":\"1640993508049424386\",\"connData\":\"\",\"conditions\":\"\",\"parentId\":\"1640993508049424384\"},{\"nodeModel\":4,\"nodeName\":\"知会\",\"nodeId\":\"1640993535555670016\",\"parentId\":\"1640993508049424384\",\"completeConn\":\"0\",\"forwardMode\":\"1\",\"processMode\":\"2\",\"allowAdd\":\"1\",\"notifyState\":\"1\",\"timeLimit\":0,\"perData\":\"{\\\\\\\"1640993433239818240\\\\\\\":\\\\\\\"1\\\\\\\",\\\\\\\"1640993434883985408\\\\\\\":\\\\\\\"1\\\\\\\"}\",\"handlerList\":\"[{\\\\\\\"handlerId\\\\\\\":\\\\\\\"547\\\\\\\",\\\\\\\"handlerName\\\\\\\":\\\\\\\"xuch01\\\\\\\",\\\\\\\"handlerType\\\\\\\":1,\\\\\\\"handlerSort\\\\\\\":1}]\",\"msgConfigList\":\"\"}]]},{\"nodeModel\":9,\"nodeName\":\"分支汇聚\",\"nodeId\":\"1640993508053618688\"},{\"nodeModel\":8,\"nodeName\":\"流程结束\",\"nodeId\":\"1640993392605401089\"}]"
	nodes := FlowParserServiceImpl.Parser(data)
	hlog.Info("数据为%v", nodes)
}
