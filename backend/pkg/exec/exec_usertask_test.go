package exec

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"testing"
)

func TestUserTaskExecution_agree(t *testing.T) {
	var instTaskParamMap = make(map[string]interface{})
	instTaskParamMap["testparam1"] = "testparam1"
	instTaskParamMap["testparam2"] = "testparam22222"
	instTaskParamMap["testparam3"] = "testparam33333"
	instTaskParamMap["testparam3"] = "testparam4"

	agree := Agree("425987729969223", "547", "xuch01", "测试", instTaskParamMap)
	hlog.Info(agree)
}

func TestUserTaskExecution_disagree(t *testing.T) {
	var instTaskParamMap = make(map[string]interface{})
	instTaskParamMap["testparam1"] = "testparam1"
	instTaskParamMap["testparam2"] = "testparam22222"
	instTaskParamMap["testparam3"] = "testparam33333"
	instTaskParamMap["testparam3"] = "testparam4"

	agree := Disagree("425987729969223", "547", "xuch01", "测试")
	hlog.Info(agree)
}

func TestSave(t *testing.T) {
	var instTaskParamMap = make(map[string]interface{})
	instTaskParamMap["testparam1"] = "testparam1"
	instTaskParamMap["testparam2"] = "testparam22222"
	instTaskParamMap["testparam3"] = "testparam33333"
	instTaskParamMap["testparam3"] = "testparam4"

	save := Save("425987729969223", "547", "xuch01", "测试", instTaskParamMap)
	hlog.Info(save)
}
