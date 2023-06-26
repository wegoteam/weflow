package usertask

import (
	"github.com/wegoteam/weflow/internal/base"
	"github.com/wegoteam/weflow/internal/biz/entity/bo"
	"github.com/wegoteam/weflow/internal/consts"
	weflowApi "github.com/wegoteam/weflow/pkg/api"
	"github.com/wegoteam/weflow/pkg/common/entity"
)

// GetTodoUserTaskList
// @Description: 获取待办用户任务列表
// @return base.Page[bo.UserTaskResult]
func GetTodoUserTaskList(param *entity.UserTaskQueryBO) *base.Response {
	pageResult, err := weflowApi.PageTodoUserTasks(param)
	if err != nil {
		return base.Fail(consts.ERROR, err.Error())
	}
	usertask := make([]bo.UserTaskTodoResult, len(pageResult.Records))
	for i, val := range pageResult.Records {
		usertask[i] = bo.UserTaskTodoResult{
			InstTaskID:     val.InstTaskID,
			TaskName:       val.TaskName,
			InstStatus:     val.TStatus,
			StartTime:      val.StartTime,
			EndTime:        val.EndTime,
			CreateUserID:   val.CreateUserID,
			CreateUserName: val.CreateUserName,
			NodeTaskID:     val.NodeTaskID,
			NodeID:         val.NodeID,
			ParentID:       val.ParentID,
			NodeModel:      val.NodeModel,
			NodeName:       val.NodeName,
			UserTaskID:     val.UserTaskID,
		}
	}
	page := &base.Page[bo.UserTaskTodoResult]{
		Total:    pageResult.Total,
		Records:  usertask,
		PageNum:  param.PageNum,
		PageSize: param.PageSize,
	}
	return base.OK(page)
}

// GetDoneUserTaskList
// @Description: 获取已办用户任务列表
// @param: param 查询参数
// @return base.Page[bo.UserTaskResult]
func GetDoneUserTaskList(param *entity.UserTaskQueryBO) *base.Response {
	pageResult, err := weflowApi.PageDoneUserTasks(param)
	if err != nil {
		return base.Fail(consts.ERROR, err.Error())
	}
	usertask := make([]bo.UserTaskResult, len(pageResult.Records))
	for i, val := range pageResult.Records {
		usertask[i] = bo.UserTaskResult{
			InstTaskID:     val.InstTaskID,
			TaskName:       val.TaskName,
			InstStatus:     val.TStatus,
			StartTime:      val.StartTime,
			EndTime:        val.EndTime,
			CreateUserID:   val.CreateUserID,
			CreateUserName: val.CreateUserName,
			NodeTaskID:     val.NodeTaskID,
			NodeID:         val.NodeID,
			ParentID:       val.ParentID,
			NodeModel:      val.NodeModel,
			NodeName:       val.NodeName,
			UserTaskID:     val.UserTaskID,
		}
	}
	page := &base.Page[bo.UserTaskResult]{
		Total:    pageResult.Total,
		Records:  usertask,
		PageNum:  param.PageNum,
		PageSize: param.PageSize,
	}
	return base.OK(page)
}

// GetReceivedUserTaskList
// @Description: 获取用户任务列表（我收到的）
// @param: param 查询参数
// @return base.Page[bo.UserTaskResult]
func GetReceivedUserTaskList(param *entity.UserTaskQueryBO) *base.Response {
	pageResult, err := weflowApi.PageReceivedUserTasks(param)
	if err != nil {
		base.Fail(consts.ERROR, err.Error())
	}
	usertask := make([]bo.UserTaskResult, len(pageResult.Records))
	for i, val := range pageResult.Records {
		usertask[i] = bo.UserTaskResult{
			InstTaskID:     val.InstTaskID,
			TaskName:       val.TaskName,
			InstStatus:     val.TStatus,
			StartTime:      val.StartTime,
			EndTime:        val.EndTime,
			CreateUserID:   val.CreateUserID,
			CreateUserName: val.CreateUserName,
			NodeTaskID:     val.NodeTaskID,
			NodeID:         val.NodeID,
			ParentID:       val.ParentID,
			NodeModel:      val.NodeModel,
			NodeName:       val.NodeName,
			UserTaskID:     val.UserTaskID,
		}
	}
	page := &base.Page[bo.UserTaskResult]{
		Total:    pageResult.Total,
		Records:  usertask,
		PageNum:  param.PageNum,
		PageSize: param.PageSize,
	}
	return base.OK(page)
}

// AgreeUserTask
// @Description: 同意用户任务
// @param: param 参数
// @return *base.Response
func AgreeUserTask(param *bo.UserTaskAgreeBO) *base.Response {
	err := weflowApi.Agree(param.UserTaskID, param.OpUserID, param.OpUserName, param.OpinionDesc, param.Params)
	if err != nil {
		base.Fail(consts.ERROR, err.Error())
	}
	return base.OK(true)
}

// DisagreeUserTask
// @Description: 不同意用户任务
// @param: param 参数
// @return *base.Response
func DisagreeUserTask(param *bo.UserTaskDisagreeBO) *base.Response {
	err := weflowApi.Disagree(param.UserTaskID, param.OpUserID, param.OpUserName, param.OpinionDesc)
	if err != nil {
		base.Fail(consts.ERROR, err.Error())
	}
	return base.OK(true)
}

// SaveUserTask
// @Description: 保存用户任务
// @param: param 参数
// @return *base.Response
func SaveUserTask(param *bo.UserTaskSaveBO) *base.Response {
	err := weflowApi.Save(param.UserTaskID, param.OpUserID, param.OpUserName, param.OpinionDesc, param.Params)
	if err != nil {
		base.Fail(consts.ERROR, err.Error())
	}
	return base.OK(true)
}
