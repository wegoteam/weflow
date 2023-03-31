package parser

import (
	"encoding/json"
	"github.com/cloudwego/hertz/pkg/common/hlog"
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
	hlog.Errorf("parse yaml error - %v", "ces")
}

/*
*
将字符串解析为节点实体
*/
func (receiver *FlowParser) Parser(data string) []*entity.NodeModelBO {

	var nodes = make([]*entity.NodeModelEntity, 0)
	var datas = make([]*entity.NodeModelBO, 0)
	if utils.IsStrBlank(data) {
		hlog.Info("节点json数据为空，无需解析")
		return datas
	}
	err := json.Unmarshal([]byte(data), &nodes)
	if err != nil {
		hlog.Warnf("解析流程图失败，错误信息：%s", err.Error())
	}
	//遍历解析节点
	for _, node := range nodes {
		parserNodeBo := ParserNode(node)

		datas = append(datas, parserNodeBo)
	}
	hlog.Info("解析流程图成功")
	return datas
}

//type INodeParser interface {
//	ParserNode(node *entity.NodeModelEntity) *entity.NodeModelBO
//}
//
//type NodeParser struct {
//}

func ParserNode(node *entity.NodeModelEntity) *entity.NodeModelBO {
	var bo = new(entity.NodeModelBO)
	err := utils.BeanCopy(bo, node)
	if err != nil {
		hlog.Errorf("节点属性转换失败%v\n", err)
	}

	return bo
}
