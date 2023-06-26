package api

import (
	"github.com/wegoteam/weflow/pkg/common/entity"
	"github.com/wegoteam/weflow/pkg/exec"
	"github.com/wegoteam/weflow/pkg/service"
)

// PageInitiatingInstTasks
// @Description: 已发列表；获取发起人的实例任务
// @param: userID 发起人用户ID
// @param: param 查询参数
// @return []entity.InstTaskResult
func PageInitiatingInstTasks(param *entity.InstTaskQueryBO) (*entity.Page[entity.InstTaskResult], error) {
	return service.PageInitiatingInstTasks(param)
}

// Start
// @Description: 发起实例任务
// @param: modelID 模板ID
// @param: userID 发起人ID
// @param: userName 发起人名称
// @param: params 参数
func Start(modelID, userID, userName string, params map[string]any) (string, error) {
	return exec.Start(modelID, userID, userName, params)
}

// Stop
// @Description: 停止实例任务
// @param: instTaskID 实例任务ID
// @param: opUserID 操作人ID
// @param: opUserName 操作人名称
// @param: opinionDesc 意见描述
// @return bool
func Stop(instTaskID, opUserID, opUserName, opinionDesc string) error {
	return exec.Stop(instTaskID, opUserID, opUserName, opinionDesc)
}

// Suspend
// @Description: 暂停、挂起实例任务
// @param: instTaskID 实例任务ID
// @param: opUserID 操作人ID
// @param: opUserName 操作人名称
// @param: opinionDesc 意见描述
// @return bool
func Suspend(instTaskID, opUserID, opUserName, opinionDesc string) error {
	return exec.Suspend(instTaskID, opUserID, opUserName, opinionDesc)
}

// Sesume
// @Description: 恢复实例任务
// @param: instTaskID 实例任务ID
// @param: opUserID 操作人ID
// @param: opUserName 操作人名称
// @param: opinionDesc 意见描述
// @return bool
func Sesume(instTaskID, opUserID, opUserName, opinionDesc string) error {
	return exec.Sesume(instTaskID, opUserID, opUserName, opinionDesc)
}

// Delete
// @Description: 删除实例任务
// @param: instTaskID 实例任务ID
// @param: opUserID 操作人ID
// @param: opUserName 操作人名称
// @param: opinionDesc 意见描述
// @return error
func Delete(instTaskID, opUserID, opUserName, opinionDesc string) error {
	return exec.Delete(instTaskID, opUserID, opUserName, opinionDesc)
}
