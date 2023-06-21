package usertask

import (
	"github.com/wegoteam/weflow/internal/base"
	"github.com/wegoteam/weflow/internal/entity/bo"
	"github.com/wegoteam/weflow/pkg/common/entity"
	"github.com/wegoteam/weflow/pkg/service"
)

// GetTodoUserTaskList
// @Description: 获取待办用户任务列表
// @return base.Page[bo.UserTaskResult]
func GetTodoUserTaskList(param *entity.UserTaskQueryBO) base.Page[bo.UserTaskTodoResult] {
	pageResult := service.PageTodoUserTasks("547", param)
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
	return *page
}

// GetDoneUserTaskList
// @Description: 获取已办用户任务列表
// @param: param 查询参数
// @return base.Page[bo.UserTaskResult]
func GetDoneUserTaskList(param *entity.UserTaskQueryBO) base.Page[bo.UserTaskResult] {
	pageResult := service.PageDoneUserTasks("547", param)
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
	return *page
}

// GetReceivedUserTaskList
// @Description: 获取用户任务列表（我收到的）
// @param: param 查询参数
// @return base.Page[bo.UserTaskResult]
func GetReceivedUserTaskList(param *entity.UserTaskQueryBO) base.Page[bo.UserTaskResult] {
	pageResult := service.PageTodoUserTasks("547", param)
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
	return *page
}
