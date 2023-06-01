package exec

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/wegoteam/weflow/pkg/common/constant"
	"github.com/wegoteam/weflow/pkg/common/entity"
	"github.com/wegoteam/weflow/pkg/common/utils"
	"github.com/wegoteam/weflow/pkg/model"
	"github.com/wegoteam/weflow/pkg/service"
)

// execInstData
// @Description: 保存实例数据
// @receiver instTaskExecution
func (instTaskExecution *InstTaskExecution) execInstData() {

	execution := instTaskExecution.Execution
	//转换实例任务
	addInstTask := transformInstTaskExecution(instTaskExecution)
	//转换实例节点任务
	addInstNodeTasks, editInstNodeTasks, delInstNodeTasks := transformInstNodeTask(*execution.InstNodeTasks)
	//转换实例用户任务
	addInstUserTasks, editInstUserTasks, delInstUserTasks := transformInstUserTask(*execution.UserTasks)
	//转换实例任务日志
	addInstTaskOpLogs := transformInstTaskOpLog(*execution.InstTaskOpLogs)
	//转换实例任务表单权限
	addInstTaskFormPers, _, _ := transformInstTaskFormPer(*execution.TaskFormPers)
	//转换实例任务参数
	addInstTaskParams := service.TransformInstTaskParam(execution.InstTaskID, execution.InstTaskParamMap, execution.Now)
	//开启事务
	tx := MysqlDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	//保存实例节点任务
	instTaskErr := tx.Create(addInstTask).Error
	if instTaskErr != nil {
		hlog.Error("实例任务[%v]保存实例任务失败", execution.InstTaskID, instTaskErr)
		panic(instTaskErr)
	}
	//保存实例节点任务
	if utils.IsNotEmptySlice(addInstNodeTasks) {
		addInstNodeTaskErr := tx.CreateInBatches(addInstNodeTasks, len(addInstNodeTasks)).Error
		if addInstNodeTaskErr != nil {
			hlog.Error("实例任务[%v]保存实例节点任务失败", execution.InstTaskID, addInstNodeTaskErr)
			panic(addInstNodeTaskErr)
		}
	}
	if utils.IsNotEmptySlice(editInstNodeTasks) {
		for _, editInstNodeTask := range editInstNodeTasks {
			editInstNodeTaskErr := tx.Where("inst_task_id = ? and node_task_id = ?", execution.InstTaskID, editInstNodeTask.NodeTaskID).Updates(editInstNodeTask).Error
			if editInstNodeTaskErr != nil {
				hlog.Error("实例任务[%v]更新实例节点任务失败", execution.InstTaskID, editInstNodeTaskErr)
				panic(editInstNodeTaskErr)
			}
		}
	}
	if utils.IsNotEmptySlice(delInstNodeTasks) {
		delInstNodeTaskIds := make([]string, 0)
		for _, delInstNodeTask := range delInstNodeTasks {
			delInstNodeTaskIds = append(delInstNodeTaskIds, delInstNodeTask.NodeTaskID)
		}
		delInstNodeTaskErr := tx.Where("inst_task_id = ? and node_task_id IN ?", execution.InstTaskID, delInstNodeTaskIds).Delete(&model.InstNodeTask{}).Error
		if delInstNodeTaskErr != nil {
			hlog.Error("实例任务[%v]删除实例节点任务失败", execution.InstTaskID, delInstNodeTaskErr)
			panic(delInstNodeTaskErr)
		}
	}
	//保存实例用户任务
	if utils.IsNotEmptySlice(addInstUserTasks) {
		addInstUserTaskErr := tx.CreateInBatches(addInstUserTasks, len(addInstUserTasks)).Error
		if addInstUserTaskErr != nil {
			hlog.Error("实例任务[%v]保存实例用户任务失败", execution.InstTaskID, addInstUserTaskErr)
			panic(addInstUserTaskErr)
		}
	}
	if utils.IsNotEmptySlice(editInstUserTasks) {
		for _, editInstUserTask := range editInstUserTasks {
			editInstUserTaskErr := tx.Where("inst_task_id = ? user_task_id = ?", execution.InstTaskID, editInstUserTask.UserTaskID).Updates(editInstUserTask).Error
			if editInstUserTaskErr != nil {
				hlog.Error("实例任务[%v]更新实例用户任务失败", execution.InstTaskID, editInstUserTaskErr)
				panic(editInstUserTaskErr)
			}
		}
	}
	if utils.IsNotEmptySlice(delInstUserTasks) {
		delInstUserTaskIds := make([]string, 0)
		for _, delInstUserTask := range delInstUserTasks {
			delInstUserTaskIds = append(delInstUserTaskIds, delInstUserTask.NodeTaskID)
		}
		delInstUserTaskErr := tx.Where("inst_task_id = ? user_task_id IN ?", execution.InstTaskID, delInstUserTaskIds).Delete(&model.InstUserTask{}).Error
		if delInstUserTaskErr != nil {
			hlog.Error("实例任务[%v]删除实例用户任务失败", execution.InstTaskID, delInstUserTaskErr)
			panic(delInstUserTaskErr)
		}
	}
	//保存实例任务日志
	if utils.IsNotEmptySlice(addInstTaskOpLogs) {
		addInstTaskFormPerErr := tx.CreateInBatches(addInstTaskOpLogs, len(addInstTaskOpLogs)).Error
		if addInstTaskFormPerErr != nil {
			hlog.Error("实例任务[%v]保存实例任务日志失败", execution.InstTaskID, addInstTaskFormPerErr)
			panic(addInstTaskFormPerErr)
		}
	}
	//保存实例节点任务表单权限
	if utils.IsNotEmptySlice(addInstTaskFormPers) {
		addInstTaskFormPerErr := tx.CreateInBatches(addInstTaskFormPers, len(addInstTaskFormPers)).Error
		if addInstTaskFormPerErr != nil {
			hlog.Error("实例任务[%v]保存实例任务表单权限失败", execution.InstTaskID, addInstTaskFormPerErr)
			panic(addInstTaskFormPerErr)
		}
	}
	//保存实例任务参数
	if addInstTaskParams != nil && len(addInstTaskParams) > 0 {
		addInstTaskParamErr := tx.CreateInBatches(addInstTaskParams, len(addInstTaskParams)).Error
		if addInstTaskParamErr != nil {
			hlog.Error("实例任务[%v]保存实例任务参数失败", execution.InstTaskID, addInstTaskParamErr)
			panic(addInstTaskParamErr)
		}
	}
	//提交事务
	tx.Commit()
}

// transformInstTaskExecution
// @Description: 转换实例任务
// @param instTaskExecution
// @return *model.InstTaskDetail
func transformInstTaskExecution(instTaskExecution *InstTaskExecution) *model.InstTaskDetail {
	execution := instTaskExecution.Execution
	//保存实例任务
	var instTask = &model.InstTaskDetail{
		InstTaskID:     execution.InstTaskID,
		ModelID:        instTaskExecution.ModelID,
		ProcessDefID:   execution.ProcessDefId,
		FormDefID:      execution.FormDefId,
		VersionID:      instTaskExecution.VersionID,
		TaskName:       execution.InstTaskName,
		Status:         int32(execution.InstTaskStatus),
		Remark:         "",
		CreateTime:     execution.Now,
		CreateUserID:   instTaskExecution.OpUserID,
		CreateUserName: instTaskExecution.OpUserName,
		UpdateTime:     execution.Now,
		UpdateUserID:   instTaskExecution.OpUserID,
		UpdateUserName: instTaskExecution.OpUserName,
		StartTime:      execution.Now,
		EndTime:        execution.Now,
	}
	return instTask
}

// transformInstNodeTask
// @Description: 转换实例节点任务
// @param instNodeTasks
// @return addInstNodeTask 添加的实例节点任务
// @return editInstNodeTask 编辑的实例节点任务
// @return delInstNodeTask 删除的实例节点任务
func transformInstNodeTask(instNodeTasks []entity.InstNodeTaskBO) (addInstNodeTask, editInstNodeTask, delInstNodeTask []model.InstNodeTask) {
	addInstNodeTask = make([]model.InstNodeTask, 0)
	editInstNodeTask = make([]model.InstNodeTask, 0)
	delInstNodeTask = make([]model.InstNodeTask, 0)

	if utils.IsEmptySlice(instNodeTasks) {
		return addInstNodeTask, editInstNodeTask, delInstNodeTask
	}
	for _, instNodeTask := range instNodeTasks {
		switch instNodeTask.ExecOpType {
		case constant.OperationTypeAdd:
			nodeTaskModel := &model.InstNodeTask{
				InstTaskID:     instNodeTask.InstTaskID,
				NodeTaskID:     instNodeTask.NodeTaskID,
				NodeID:         instNodeTask.NodeID,
				ParentID:       instNodeTask.ParentID,
				NodeModel:      instNodeTask.NodeModel,
				NodeName:       instNodeTask.NodeName,
				ApproveType:    instNodeTask.ApproveType,
				NoneHandler:    instNodeTask.NoneHandler,
				AppointHandler: instNodeTask.AppointHandler,
				HandleMode:     instNodeTask.HandleMode,
				FinishMode:     instNodeTask.FinishMode,
				BranchMode:     instNodeTask.BranchMode,
				DefaultBranch:  instNodeTask.DefaultBranch,
				BranchLevel:    instNodeTask.BranchLevel,
				ConditionGroup: instNodeTask.ConditionGroup,
				ConditionExpr:  instNodeTask.ConditionExpr,
				Remark:         instNodeTask.Remark,
				Status:         instNodeTask.Status,
				CreateTime:     instNodeTask.CreateTime,
				UpdateTime:     instNodeTask.UpdateTime,
			}
			addInstNodeTask = append(addInstNodeTask, *nodeTaskModel)
		case constant.OperationTypeUpdate:
			nodeTaskModel := &model.InstNodeTask{
				NodeTaskID: instNodeTask.NodeTaskID,
				Status:     instNodeTask.Status,
				UpdateTime: instNodeTask.UpdateTime,
			}
			editInstNodeTask = append(editInstNodeTask, *nodeTaskModel)
		case constant.OperationTypeDelete:
			nodeTaskModel := &model.InstNodeTask{
				NodeTaskID: instNodeTask.NodeTaskID,
			}
			delInstNodeTask = append(delInstNodeTask, *nodeTaskModel)
		}
	}
	return addInstNodeTask, editInstNodeTask, delInstNodeTask
}

// transformInstUserTask
// @Description: 转换实例用户任务
// @return addInstUserTask
// @return editInstUserTask
// @return delInstUserTask
func transformInstUserTask(userTasks []entity.UserTaskBO) (addInstUserTask, editInstUserTask, delInstUserTask []model.InstUserTask) {
	addInstUserTask = make([]model.InstUserTask, 0)
	editInstUserTask = make([]model.InstUserTask, 0)
	delInstUserTask = make([]model.InstUserTask, 0)

	if utils.IsEmptySlice(userTasks) {
		return addInstUserTask, editInstUserTask, delInstUserTask
	}
	for _, userTask := range userTasks {
		switch userTask.ExecOpType {
		case constant.OperationTypeAdd:
			userTaskModel := &model.InstUserTask{
				InstTaskID:   userTask.InstTaskID,
				NodeTaskID:   userTask.NodeTaskID,
				NodeID:       userTask.NodeID,
				UserTaskID:   userTask.UserTaskID,
				Type:         userTask.Type,
				Strategy:     userTask.Strategy,
				NodeUserName: userTask.NodeUserName,
				NodeUserID:   userTask.NodeUserID,
				Sort:         userTask.Sort,
				Obj:          userTask.Obj,
				Relative:     userTask.Relative,
				Status:       userTask.Status,
				CreateTime:   userTask.CreateTime,
				UpdateTime:   userTask.UpdateTime,
				HandleTime:   userTask.HandleTime,
				OpUserID:     userTask.OpUserID,
				OpUserName:   userTask.OpUserName,
				Opinion:      userTask.Opinion,
				OpinionDesc:  userTask.OpinionDesc,
			}
			addInstUserTask = append(addInstUserTask, *userTaskModel)
		case constant.OperationTypeUpdate:
			userTaskModel := &model.InstUserTask{
				UserTaskID:  userTask.UserTaskID,
				Status:      userTask.Status,
				UpdateTime:  userTask.UpdateTime,
				HandleTime:  userTask.HandleTime,
				Opinion:     userTask.Opinion,
				OpinionDesc: userTask.OpinionDesc,
			}
			editInstUserTask = append(editInstUserTask, *userTaskModel)
		case constant.OperationTypeDelete:
			userTaskModel := &model.InstUserTask{
				UserTaskID: userTask.UserTaskID,
			}
			delInstUserTask = append(delInstUserTask, *userTaskModel)
		}
	}
	return addInstUserTask, editInstUserTask, delInstUserTask
}

// transformInstTaskFormPer
// @Description: 转换实例任务表单权限
// @param taskFormPers
// @return addInstTaskFormPers
// @return editInstTaskFormPers
// @return delInstTaskFormPers
func transformInstTaskFormPer(taskFormPers []entity.TaskFormPerBO) (addInstTaskFormPers, editInstTaskFormPers, delInstTaskFormPers []model.InstNodeTaskFormper) {
	addInstTaskFormPers = make([]model.InstNodeTaskFormper, 0)
	editInstTaskFormPers = make([]model.InstNodeTaskFormper, 0)
	delInstTaskFormPers = make([]model.InstNodeTaskFormper, 0)

	if utils.IsEmptySlice(taskFormPers) {
		return addInstTaskFormPers, editInstTaskFormPers, delInstTaskFormPers
	}
	for _, taskFormPer := range taskFormPers {
		switch taskFormPer.ExecOpType {
		case constant.OperationTypeAdd:
			var addInstTaskFormPer = &model.InstNodeTaskFormper{
				InstTaskID: taskFormPer.InstTaskID,
				NodeTaskID: taskFormPer.NodeTaskID,
				NodeID:     taskFormPer.NodeID,
				ElemID:     taskFormPer.ElemID,
				ElemPID:    taskFormPer.ElemPID,
				Per:        taskFormPer.Per,
			}
			addInstTaskFormPers = append(addInstTaskFormPers, *addInstTaskFormPer)
		case constant.OperationTypeUpdate:

		case constant.OperationTypeDelete:

		}
	}
	return addInstTaskFormPers, editInstTaskFormPers, delInstTaskFormPers
}

// transformInstTaskOpLog
// @Description: 转换实例任务日志
// @param instTaskOpLogs
// @return []model.InstTaskOpLog
func transformInstTaskOpLog(instTaskOpLogs []entity.InstTaskOpLogBO) []model.InstTaskOpLog {
	var addInstTaskOpLogs = make([]model.InstTaskOpLog, 0)
	if utils.IsEmptySlice(instTaskOpLogs) {
		return addInstTaskOpLogs
	}

	for _, instTaskOpLog := range instTaskOpLogs {
		addInstTaskOpLog := &model.InstTaskOpLog{
			InstTaskID: instTaskOpLog.InstTaskID,
			NodeID:     instTaskOpLog.NodeID,
			NodeName:   instTaskOpLog.NodeName,
			CreateTime: instTaskOpLog.CreateTime,
			UpdateTime: instTaskOpLog.UpdateTime,
			Type:       instTaskOpLog.Type,
			Remark:     instTaskOpLog.Remark,
		}
		addInstTaskOpLogs = append(addInstTaskOpLogs, *addInstTaskOpLog)
	}
	return addInstTaskOpLogs
}
