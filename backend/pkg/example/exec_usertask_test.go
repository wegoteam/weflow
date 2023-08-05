package example

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/wegoteam/weflow/pkg/exec"
	"github.com/wegoteam/weflow/pkg/service"
	"testing"
)

func TestUserTaskExecution_agree(t *testing.T) {
	var instTaskParamMap = make(map[string]interface{})
	instTaskParamMap["testparam1"] = "testparam1"
	instTaskParamMap["testparam2"] = "testparam22222"
	instTaskParamMap["testparam3"] = "testparam33333"
	instTaskParamMap["testparam4"] = "testparam4"

	userID := "547"
	userName := "xuch01"
	desc := "测试"

	userTasks := service.GetTodoUserTasks(userID)
	if userTasks == nil || len(*userTasks) == 0 {
		hlog.Info("当前待办任务为空")
		return
	}
	var userTask = (*userTasks)[0]
	err := exec.Agree(userTask.UserTaskID, userID, userName, desc, instTaskParamMap)
	hlog.Info(err)
}

func TestUserTaskExecution_disagree(t *testing.T) {
	var instTaskParamMap = make(map[string]interface{})
	instTaskParamMap["testparam1"] = "testparam1"
	instTaskParamMap["testparam2"] = "testparam22222"
	instTaskParamMap["testparam3"] = "testparam33333"
	instTaskParamMap["testparam4"] = "testparam4"

	userID := "547"
	userName := "xuch01"
	desc := "测试"

	userTasks := service.GetTodoUserTasks(userID)
	if userTasks == nil || len(*userTasks) == 0 {
		hlog.Info("当前待办任务为空")
		return
	}
	var userTask = (*userTasks)[0]
	err := exec.Disagree(userTask.UserTaskID, userID, userName, desc)
	hlog.Info(err)
}

func TestSave(t *testing.T) {
	var instTaskParamMap = make(map[string]interface{})
	instTaskParamMap["testparam1"] = "testparam1"
	instTaskParamMap["testparam2"] = "testparam22222"
	instTaskParamMap["testparam3"] = "testparam33333"
	instTaskParamMap["testparam4"] = "testparam4"

	userID := "547"
	userName := "xuch01"
	desc := "测试"

	userTasks := service.GetTodoUserTasks(userID)
	if userTasks == nil || len(*userTasks) == 0 {
		hlog.Info("当前待办任务为空")
		return
	}
	var userTask = (*userTasks)[0]
	err := exec.Save(userTask.UserTaskID, userID, userName, desc, instTaskParamMap)
	hlog.Info(err)
}
