package expr

import (
	"github.com/antonmedv/expr"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/wegoteam/weflow/pkg/common/utils"
)

/**
执行条件表达式
*/
func ExecExpr(exprData string, env map[string]interface{}) bool {

	//条件为空，默认条件通过
	if utils.IsStrBlank(exprData) {
		return true
	}

	output, err := expr.Eval(exprData, env)
	if err != nil {
		hlog.Infof("执行条件表达式错误：%v", err)
		return false
	}
	switch output.(type) {
	case bool:
		hlog.Infof("执行条件表达式返回结果：%v", output.(bool))
		return output.(bool)
	case int:
		hlog.Infof("执行条件表达式返回结果：%v", output.(int))
	case string:
		hlog.Infof("执行条件表达式返回结果：%v", output.(string))
	default:
		hlog.Infof("执行条件表达式返回结果：%v", output)
	}
	return false
}
