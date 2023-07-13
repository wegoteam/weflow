package insttask

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/wegoteam/weflow/internal/base"
	"github.com/wegoteam/weflow/internal/biz/entity/bo"
	"github.com/wegoteam/weflow/internal/consts"
	weflowApi "github.com/wegoteam/weflow/pkg/api"
	"github.com/wegoteam/weflow/pkg/common/entity"
	"github.com/wegoteam/weflow/pkg/common/utils"
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
			CreateTime:     utils.TimeToString(instTask.CreateTime),
			CreateUserID:   instTask.CreateUserID,
			CreateUserName: instTask.CreateUserName,
			UpdateTime:     utils.TimeToString(instTask.UpdateTime),
			UpdateUserID:   instTask.UpdateUserID,
			UpdateUserName: instTask.UpdateUserName,
			StartTime:      utils.TimeToString(instTask.StartTime),
			EndTime:        utils.TimeToString(instTask.EndTime),
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
		return base.Fail(consts.ERROR, err.Error())
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
		return base.Fail(consts.ERROR, err.Error())
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
		return base.Fail(consts.ERROR, err.Error())
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
		return base.Fail(consts.ERROR, err.Error())
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
		return base.Fail(consts.ERROR, err.Error())
	}
	return base.OK(true)
}

// GetInsttaskModelDetail
// @Description: 查询实例任务模板详情
// @param: instTaskID 实例任务ID
// @return *base.Response
func GetInsttaskModelDetail(instTaskID string) *base.Response {
	instTask, err := weflowApi.GetInstTask(instTaskID)
	if err != nil {
		hlog.Errorf("获取实例任务模板详情 error: %v", err)
		return base.Fail(consts.ERROR, "查询实例任务模板详情失败")
	}
	modelDetail, err := weflowApi.GetModelAndVersionInfo(instTask.ModelID, instTask.VersionID)
	if err != nil {
		hlog.Errorf("获取实例任务模板详情 error: %v", err)
		return base.Fail(consts.ERROR, "查询实例任务模板详情失败")
	}
	modelResult := &bo.ModelAndVersionInfoResult{
		ModelID:      modelDetail.ModelID,
		ModelTitle:   modelDetail.ModelTitle,
		ProcessDefID: modelDetail.ProcessDefID,
		FormDefID:    modelDetail.FormDefID,
		ModelGroupID: modelDetail.ModelGroupID,
		IconURL:      modelDetail.IconURL,
		Status:       modelDetail.Status,
		Remark:       modelDetail.Remark,
		CreateTime:   modelDetail.CreateTime,
		CreateUser:   modelDetail.CreateUser,
		UpdateTime:   modelDetail.UpdateTime,
		UpdateUser:   modelDetail.UpdateUser,
		FlowContent:  modelDetail.FlowContent,
		FormContent:  modelDetail.FormContent,
	}
	return base.OK(modelResult)
}

// GetInsttaskAllDetail
// @Description: 查询实例任务详情
// @param: instTaskID 实例任务ID
// @param: userTaskID 用户任务ID
// @return *base.Response
func GetInsttaskAllDetail(instTaskID, userTaskID string) *base.Response {
	instTask, err := weflowApi.GetInsttaskAllDetail(instTaskID, userTaskID)
	if err != nil {
		hlog.Errorf("查询实例任务详情 error: %v", err)
		return base.Fail(consts.ERROR, "查询实例任务详情失败")
	}
	return base.OK(instTask)
}
