package service

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/golang-module/carbon/v2"
	"github.com/pkg/errors"
	"github.com/wegoteam/weflow/pkg/common/constant"
	"github.com/wegoteam/weflow/pkg/common/entity"
	"github.com/wegoteam/weflow/pkg/common/utils"
	"github.com/wegoteam/weflow/pkg/model"
	"gorm.io/gorm"
)

// GetInstTask
// @Description: 获取实例任务
// @param: instTaskID 实例任务ID
// @return *entity.InstTaskResult
func GetInstTask(instTaskID string) (*entity.InstTaskResult, error) {
	if utils.IsStrBlank(instTaskID) {
		return nil, errors.New("实例任务id不能为空")
	}
	var instTask = &model.InstTaskDetail{}
	err := MysqlDB.Where("inst_task_id = ?", instTaskID).Find(instTask).Error
	if err != nil {
		hlog.Errorf("查询实例任务失败：%s", err.Error())
		return nil, errors.New("查询实例任务失败")
	}
	if instTask == nil {
		return nil, errors.New("实例任务不存在")
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
	}, nil
}

// GetInsttaskAllDetail
// @Description: 查询实例任务详情（实例任务，节点任务，用户任务信息，流程任务参数）
// @param: instTaskID 实例任务ID
// @param: userTaskID 用户任务ID
// @return *entity.InstTaskAllDetailResult
// @return error
func GetInsttaskAllDetail(instTaskID, userTaskID string) (*entity.InstTaskAllDetailResult, error) {
	instTask, insttaskErr := GetInstTask(instTaskID)
	if insttaskErr != nil {
		hlog.Errorf("查询实例任务失败：%s", insttaskErr.Error())
		return nil, errors.New("查询实例任务详情失败")
	}

	nodeTasks, nodeTaskErr := GetInstNodeTasks(instTaskID)
	if nodeTaskErr != nil {
		hlog.Errorf("查询节点任务失败：%s", nodeTaskErr.Error())
		return nil, errors.New("查询实例任务详情失败")
	}

	userTasks, userTaskErr := GetInstUserTasks(instTaskID)
	if userTaskErr != nil {
		hlog.Errorf("查询用户任务失败：%s", userTaskErr.Error())
		return nil, errors.New("查询实例任务详情失败")
	}

	userOpinions, userOpinionErr := GetInstUserTaskOpinions(instTaskID)
	if userOpinionErr != nil {
		hlog.Errorf("查询用户任务意见失败：%s", userOpinionErr.Error())
		return nil, errors.New("查询实例任务详情失败")
	}

	taskFormpers, taskFormperErr := GetInstTaskFormPers(instTaskID)
	if taskFormperErr != nil {
		hlog.Errorf("查询节点任务表单权限失败：%s", taskFormperErr.Error())
		return nil, errors.New("查询实例任务详情失败")
	}

	taskParams, taskParamErr := GetInstTaskParams(instTaskID)
	if taskParamErr != nil {
		hlog.Errorf("查询节点任务参数失败：%s", taskParamErr.Error())
		return nil, errors.New("查询实例任务详情失败")
	}

	var instTaskAllDetail = &entity.InstTaskAllDetailResult{
		InstTask:             instTask,
		InstNodeTasks:        nodeTasks,
		InstUserTasks:        userTasks,
		InstUserTaskOpinions: userOpinions,
		InstNodeTaskFormpers: taskFormpers,
		InstTaskParams:       taskParams,
	}

	return instTaskAllDetail, nil
}

// GetInitiatingInstTasks
// @Description: 已发列表；获取发起人的实例任务
// @param: userID 发起人用户ID
// @return *entity.InstTaskResult
func GetInitiatingInstTasks(userID string) []entity.InstTaskResult {
	var instTaskList = make([]entity.InstTaskResult, 0)
	if utils.IsStrBlank(userID) {
		return instTaskList
	}
	var instTasks = []model.InstTaskDetail{}
	MysqlDB.Where("inst_task_detail.create_user_id = ? and inst_task_detail.status in (2,3,4,5,6,7,8)", userID).Order("inst_task_detail.start_time desc").Find(&instTasks)
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

// QueryInitiatingInstTasks
// @Description: 已发列表；获取发起人的实例任务
// @param: userID 发起人用户ID
// @param: param 查询参数
// @return []entity.InstTaskResult
func QueryInitiatingInstTasks(param *entity.InstTaskQueryBO) []entity.InstTaskResult {
	var instTaskList = make([]entity.InstTaskResult, 0)
	if utils.IsStrBlank(param.UserID) {
		return instTaskList
	}
	var instTasks = []model.InstTaskDetail{}
	MysqlDB.Model(&model.InstTaskDetail{}).Scopes(BuildInstTaskQuery(param)).Where("inst_task_detail.create_user_id = ? and inst_task_detail.status in (2,3,4,5,6,7,8)", param.UserID).Order("inst_task_detail.start_time desc").Find(&instTasks)
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

// PageInitiatingInstTasks
// @Description: 已发列表；获取发起人的实例任务
// @param: userID 发起人用户ID
// @param: param 查询参数
// @return []entity.InstTaskResult
func PageInitiatingInstTasks(param *entity.InstTaskQueryBO) (*entity.Page[entity.InstTaskResult], error) {
	var instTaskList = make([]entity.InstTaskResult, 0)
	if utils.IsStrBlank(param.UserID) {
		return &entity.Page[entity.InstTaskResult]{
			Records:  instTaskList,
			Total:    0,
			PageSize: param.PageSize,
			PageNum:  param.PageNum,
		}, nil
	}
	var total int64
	err := MysqlDB.Model(&model.InstTaskDetail{}).Scopes(BuildInstTaskQuery(param)).Where("inst_task_detail.create_user_id = ? and inst_task_detail.status in (2,3,4,5,6,7,8)", param.UserID).Order("inst_task_detail.start_time desc").Count(&total).Error
	if err != nil {
		hlog.Errorf("查询发起人的实例任务失败 error: %v", err)
		return nil, errors.New("查询发起人的实例任务失败")
	}
	if total == 0 {
		return &entity.Page[entity.InstTaskResult]{
			Records:  instTaskList,
			Total:    total,
			PageSize: param.PageSize,
			PageNum:  param.PageNum,
		}, nil
	}
	var instTasks = []model.InstTaskDetail{}
	err2 := MysqlDB.Model(&model.InstTaskDetail{}).Scopes(entity.Paginate(param.PageNum, param.PageSize), BuildInstTaskQuery(param)).
		Where("inst_task_detail.create_user_id = ? and inst_task_detail.status in (2,3,4,5,6,7,8)", param.UserID).
		Order("inst_task_detail.start_time desc").Find(&instTasks).Error
	if err2 != nil {
		hlog.Errorf("查询发起人的实例任务失败 error: %v", err2)
		return nil, errors.New("查询发起人的实例任务失败")
	}
	if utils.IsEmptySlice(instTasks) {
		return &entity.Page[entity.InstTaskResult]{
			Records:  instTaskList,
			Total:    0,
			PageSize: param.PageSize,
			PageNum:  param.PageNum,
		}, nil
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
	return &entity.Page[entity.InstTaskResult]{
		Records:  instTaskList,
		Total:    total,
		PageSize: param.PageSize,
		PageNum:  param.PageNum,
	}, nil
}

// BuildInstTaskQuery
// @Description: 实例任务查询条件
// @param: param *entity.InstTaskQueryBO
// @return func(db *gorm.DB) *gorm.DB
func BuildInstTaskQuery(param *entity.InstTaskQueryBO) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		tx := db
		if param.InstStatus != 0 {
			tx = db.Where("inst_task_detail.status = ?", param.InstStatus)
		}
		if utils.IsStrNotBlank(param.TaskName) {
			tx = db.Where("inst_task_detail.task_name like ?", "%"+param.TaskName+"%")
		}
		if utils.IsStrNotBlank(param.CreateTimeStart) && utils.IsStrNotBlank(param.CreateTimeEnd) {
			carbon.Parse(param.CreateTimeStart).ToStdTime()
			tx = db.Where("inst_task_detail.create_time BETWEEN ? AND ?", carbon.Parse(param.CreateTimeStart).ToStdTime(), carbon.Parse(param.CreateTimeEnd).ToStdTime())
		}
		if utils.IsStrNotBlank(param.FinishTimeStart) && utils.IsStrNotBlank(param.FinishTimeEnd) {
			tx = db.Where("inst_task_detail.end_time BETWEEN ? AND ?", carbon.Parse(param.FinishTimeStart).ToStdTime(), carbon.Parse(param.FinishTimeEnd).ToStdTime())
		}
		return tx
	}
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
