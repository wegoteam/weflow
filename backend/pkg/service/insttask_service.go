package service

import (
	"github.com/wegoteam/weflow/pkg/common/entity"
	"github.com/wegoteam/weflow/pkg/common/utils"
	"github.com/wegoteam/weflow/pkg/model"
)

// GetInstTask
// @Description: 获取实例任务
// @param instTaskID 实例任务ID
// @return *entity.InstTaskResult
func GetInstTask(instTaskID string) *entity.InstTaskResult {

	if utils.IsStrBlank(instTaskID) {
		panic("实例任务id不能为空")
	}
	var instTask = &model.InstTaskDetail{}
	MysqlDB.Where("inst_task_id = ?", instTaskID).Find(instTask)

	if instTask == nil {
		return nil
	}
	return &entity.InstTaskResult{
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
