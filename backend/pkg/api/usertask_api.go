package api

import (
	"github.com/wegoteam/weflow/pkg/common/entity"
	"github.com/wegoteam/weflow/pkg/exec"
	"github.com/wegoteam/weflow/pkg/service"
)

// PageTodoUserTasks
// @Description: 分页待办用户任务
// @param: userID
// @return *entity.Page[entity.InstNodeAndUserTaskResult]
func PageTodoUserTasks(param *entity.UserTaskQueryBO) (*entity.Page[entity.InstNodeAndUserTaskResult], error) {
	return service.PageTodoUserTasks(param)
}

// PageDoneUserTasks
// @Description: 分页查询已办用户任务
// @param: userID 用户ID
// @param: param 查询参数
// @return *[]entity.InstNodeAndUserTaskResult
func PageDoneUserTasks(param *entity.UserTaskQueryBO) (*entity.Page[entity.InstNodeAndUserTaskResult], error) {
	return service.PageDoneUserTasks(param)
}

// PageReceivedUserTasks
// @Description: 分页查询我收到的用户任务
// @param: userID 用户ID
// @param: param 查询参数
// @return *entity.Page[entity.InstNodeAndUserTaskResult]
func PageReceivedUserTasks(param *entity.UserTaskQueryBO) (*entity.Page[entity.InstNodeAndUserTaskResult], error) {
	return service.PageReceivedUserTasks(param)
}

// Agree
// @Description: 同意
// @param: userTaskID 用户任务ID
// @param: opUserID 操作用户ID
// @param: OpUserName 操作用户名称
// @param: opinionDesc 意见描述
// @param: params 参数
// @return bool
func Agree(userTaskID, opUserID, opUserName, opinionDesc string, params map[string]any) error {
	return exec.Agree(userTaskID, opUserID, opUserName, opinionDesc, params)
}

// Save
// @Description: 保存
// @param: userTaskID 用户任务ID
// @param: opUserID 操作用户ID
// @param: OpUserName 操作用户名称
// @param: opinionDesc 意见描述
// @param: params 参数
// @return bool
func Save(userTaskID, opUserID, opUserName, opinionDesc string, params map[string]any) error {
	return exec.Save(userTaskID, opUserID, opUserName, opinionDesc, params)
}

// Disagree
// @Description: 不同意
// @param: userTaskID 用户任务ID
// @param: opUserID 操作用户ID
// @param: OpUserName 操作用户名称
// @param: opinionDesc 意见描述
// @return bool
func Disagree(userTaskID, opUserID, opUserName, opinionDesc string) error {
	return exec.Disagree(userTaskID, opUserID, opUserName, opinionDesc)
}
