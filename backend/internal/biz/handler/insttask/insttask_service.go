package insttask

import (
	"github.com/wegoteam/weflow/internal/base"
	"github.com/wegoteam/weflow/internal/biz/entity/bo"
	"github.com/wegoteam/weflow/internal/consts"
	weflowApi "github.com/wegoteam/weflow/pkg/api"
	"github.com/wegoteam/weflow/pkg/common/entity"
)

// GetInitiateInstTaskList
// @Description: 获取发起中的实例任务列表
// @param: param 查询参数
// @return *base.Response
func GetInitiateInstTaskList(param *entity.InstTaskQueryBO) *base.Response {
	//已发列表；获取发起人的实例任务
	pageResult, err := weflowApi.PageInitiatingInstTasks(param)
	if err != nil {
		return base.Fail(consts.ERROR, err.Error())
	}
	insttask := make([]bo.InstTaskResult, len(pageResult.Records))
	for i, instTask := range pageResult.Records {
		insttask[i] = bo.InstTaskResult{
			ID:             instTask.ID,
			InstTaskID:     instTask.InstTaskID,
			ModelID:        instTask.ModelID,
			VersionID:      instTask.VersionID,
			ProcessDefID:   instTask.ProcessDefID,
			FormDefID:      instTask.FormDefID,
			TaskName:       instTask.TaskName,
			Status:         instTask.Status,
			Remark:         instTask.Remark,
			CreateTime:     instTask.CreateTime,
			CreateUserID:   instTask.CreateUserID,
			CreateUserName: instTask.CreateUserName,
			UpdateTime:     instTask.UpdateTime,
			UpdateUserID:   instTask.UpdateUserID,
			UpdateUserName: instTask.UpdateUserName,
			StartTime:      instTask.StartTime,
			EndTime:        instTask.EndTime,
		}
	}
	page := &base.Page[bo.InstTaskResult]{
		Total:    pageResult.Total,
		Records:  insttask,
		PageNum:  param.PageNum,
		PageSize: param.PageSize,
	}
	return base.OK(page)
}

// Start
// @Description: 发起实例任务
// @param: modelID 模板ID
// @param: userID 发起人ID
// @param: userName 发起人名称
// @param: params 参数
func Start(param *bo.InstTaskStartBO) *base.Response {
	instTaskID, err := weflowApi.Start(param.ModelID, param.UserID, param.UserName, param.Params)
	if err != nil {
		base.Fail(consts.ERROR, err.Error())
	}
	return base.OK(instTaskID)
}

// Stop
// @Description: 停止实例任务
// @param: instTaskID 实例任务ID
// @param: opUserID 操作人ID
// @param: opUserName 操作人名称
// @param: opinionDesc 意见描述
// @return bool
func Stop(param *bo.InstTaskStopBO) *base.Response {
	err := weflowApi.Stop(param.InstTaskID, param.OpUserID, param.OpUserName, param.OpinionDesc)
	if err != nil {
		base.Fail(consts.ERROR, err.Error())
	}
	return base.OK(true)
}

// Suspend
// @Description: 暂停、挂起实例任务
// @param: instTaskID 实例任务ID
// @param: opUserID 操作人ID
// @param: opUserName 操作人名称
// @param: opinionDesc 意见描述
// @return bool
func Suspend(param *bo.InstTaskSuspendBO) *base.Response {
	err := weflowApi.Suspend(param.InstTaskID, param.OpUserID, param.OpUserName, param.OpinionDesc)
	if err != nil {
		base.Fail(consts.ERROR, err.Error())
	}
	return base.OK(true)
}

// Sesume
// @Description: 恢复实例任务
// @param: instTaskID 实例任务ID
// @param: opUserID 操作人ID
// @param: opUserName 操作人名称
// @param: opinionDesc 意见描述
// @return bool
func Sesume(param *bo.InstTaskSesumeBO) *base.Response {
	err := weflowApi.Sesume(param.InstTaskID, param.OpUserID, param.OpUserName, param.OpinionDesc)
	if err != nil {
		base.Fail(consts.ERROR, err.Error())
	}
	return base.OK(true)
}

// Delete
// @Description: 删除实例任务
// @param: instTaskID 实例任务ID
// @param: opUserID 操作人ID
// @param: opUserName 操作人名称
// @param: opinionDesc 意见描述
// @return error
func Delete(param *bo.InstTaskDeleteBO) *base.Response {
	err := weflowApi.Delete(param.InstTaskID, param.OpUserID, param.OpUserName, param.OpinionDesc)
	if err != nil {
		base.Fail(consts.ERROR, err.Error())
	}
	return base.OK(true)
}
