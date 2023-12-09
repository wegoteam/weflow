package expr

import (
	"github.com/antonmedv/expr"
	"github.com/wegoteam/weflow/pkg/common/entity"
	"github.com/wegoteam/weflow/pkg/common/utils"
	"github.com/wegoteam/wepkg/io/json"
	"github.com/wegoteam/wepkg/log/wlog"
)

// ExecExpr
// @Description: 执行条件表达式
// @param: exprData
// @param: env
// @return bool
func ExecExpr(exprData string, env map[string]interface{}) bool {
	//条件为空，默认条件通过
	if utils.IsStrBlank(exprData) {
		return true
	}
	output, err := expr.Eval(exprData, env)
	if err != nil {
		wlog.Infof("执行条件表达式错误：%v", err)
		return false
	}
	switch output.(type) {
	case bool:
		wlog.Infof("执行条件表达式返回结果：%v", output.(bool))
		return output.(bool)
	case int:
		wlog.Infof("执行条件表达式返回结果：%v", output.(int))
	case string:
		wlog.Infof("执行条件表达式返回结果：%v", output.(string))
	default:
		wlog.Infof("执行条件表达式返回结果：%v", output)
	}
	return false
}

// AnalyticalConditions
// @Description: 解析条件
// @param: conditins 条件
// 解析条件
// 流程的不同分支，根据流程连线的配置的条件项流转不同节点
// 条件类型单条件和组合条件
// 一层数组的条件是或关系
// 二层数组的条件是且关系
func AnalyticalConditions(conditins [][]entity.ProcessConditions) (string, error) {
	if conditins == nil || len(conditins) == 0 {
		return "", nil
	}
	var result entity.ConditionResult

	//var orConditions = make([]string, 0)
	for _, conditionGroups := range conditins {
		if conditionGroups == nil || len(conditionGroups) == 0 {
			continue
		}
		//var andConditions = make([]string, 0)

	}
	return json.Marshal(result)
}
