package service

import (
	"github.com/wegoteam/weflow/pkg/common/constant"
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

// GetInitiatingInstTask
// @Description: 已发列表；获取发起人的实例任务
// @param userID 发起人用户ID
// @return *entity.InstTaskResult
func GetInitiatingInstTask(userID string) []entity.InstTaskResult {

	var instTaskList = make([]entity.InstTaskResult, 0)
	if utils.IsStrBlank(userID) {
		return instTaskList
	}
	var instTasks = []model.InstTaskDetail{}
	MysqlDB.Where("create_user_id = ?", userID).Order("start_time desc").Find(&instTasks)
	if utils.IsEmptySlice(instTasks) {
		return instTaskList
	}
	for _, instTask := range instTasks {
		instTaskBO := &entity.InstTaskResult{
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
		instTaskList = append(instTaskList, *instTaskBO)
	}

	return instTaskList
}

// GetDraftInstTask
// @Description: 草稿列表；获取发起人的实例任务
// @param userID
// @return []entity.InstTaskResult
func GetDraftInstTask(userID string) []entity.InstTaskResult {

	var instTaskList = make([]entity.InstTaskResult, 0)
	if utils.IsStrBlank(userID) {
		return instTaskList
	}
	var instTasks = []model.InstTaskDetail{}
	MysqlDB.Where("create_user_id = ? and status = ?", userID, constant.InstanceTaskStatusDraft).Order("start_time desc").Find(&instTasks)
	if utils.IsEmptySlice(instTasks) {
		return instTaskList
	}
	for _, instTask := range instTasks {
		instTaskBO := &entity.InstTaskResult{
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
		instTaskList = append(instTaskList, *instTaskBO)
	}

	return instTaskList
}
