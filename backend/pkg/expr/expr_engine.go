package expr

import (
	"fmt"
	"github.com/antonmedv/expr"
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
		fmt.Printf("err: %v", err)
		return false
	}
	switch output.(type) {
	case bool:
		fmt.Printf("%v\n", output.(bool))
		return output.(bool)
	case int:
		fmt.Printf("%v\n", output.(int))
	case string:
		fmt.Printf("%v\n", output.(string))
	default:
		fmt.Printf("%v\n", output)
	}
	return true
}
