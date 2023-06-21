package insttask

import (
	"github.com/wegoteam/weflow/internal/base"
	"github.com/wegoteam/weflow/internal/biz/entity/bo"
	"github.com/wegoteam/weflow/internal/consts"
	"github.com/wegoteam/weflow/pkg/common/entity"
	"github.com/wegoteam/weflow/pkg/service"
)

// GetInitiateInstTaskList
// @Description: 获取发起中的实例任务列表
// @param: param 查询参数
// @return *base.Response
func GetInitiateInstTaskList(param *entity.InstTaskQueryBO) *base.Response {
	//已发列表；获取发起人的实例任务
	pageResult, err := service.PageInitiatingInstTasks(param)
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
