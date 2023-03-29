package parser

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/pquerna/ffjson/ffjson"
	"github.com/wego2023/weflow/pkg/common/entity"
	"github.com/wego2023/weflow/pkg/common/utils"
)

var FlowParserServiceImpl = new(FlowParser)

type FlowParser struct {
}

type FlowParserService interface {
	Parser(data string) []entity.NodeModelEntity
}

func (receiver *FlowParser) Test() {
	hlog.Info("hello world")
}

/*
*
将字符串解析为节点实体
*/
func (receiver *FlowParser) Parser(data string) []entity.NodeModelEntity {

	var nodes = make([]entity.NodeModelEntity, 10)
	if utils.IsStrBlank(data) {
		hlog.Info("节点json数据为空，无需解析")
		return nodes
	}
	err := ffjson.Unmarshal([]byte(data), &nodes)
	if err != nil {
		hlog.Error("解析流程图失败，错误信息：%s", err.Error())
	}
	hlog.Info("解析流程图成功")
	return nodes
}
