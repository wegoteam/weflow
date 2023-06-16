package exec

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/wegoteam/weflow/pkg/common/constant"
	"github.com/wegoteam/weflow/pkg/common/entity"
	"github.com/wegoteam/weflow/pkg/common/utils"
	"github.com/wegoteam/weflow/pkg/model"
	"github.com/wegoteam/weflow/pkg/service"
	"github.com/wegoteam/wepkg/snowflake"
)

// execStartInstData
// @Description: 发起执行的实例数据，进行数据处理
// @receiver instTaskExecution
func (instTaskExecution *InstTaskExecution) execStartInstData() error {
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
	//保存用户任务评论
	addInstUserTaskOpinion := &model.InstUserTaskOpinion{
		InstTaskID:  execution.InstTaskID,
		OpinionID:   snowflake.GetSnowflakeId(),
		Opinion:     int32(instTaskExecution.Opinion),
		OpinionDesc: instTaskExecution.OpinionDesc,
		OpUserID:    instTaskExecution.OpUserID,
		OpUserName:  instTaskExecution.OpUserName,
		CreateTime:  execution.Now,
		UpdateTime:  execution.Now,
		OpinionTime: execution.Now,
	}
	addInstUserTaskOpinionErr := tx.Create(addInstUserTaskOpinion).Error
	if addInstUserTaskOpinionErr != nil {
		hlog.Error("实例任务[%v]保存用户任务评论失败", execution.InstTaskID, addInstUserTaskOpinionErr)
		tx.Rollback()
		return addInstUserTaskOpinionErr
	}
	//保存实例节点任务
	instTaskErr := tx.Create(addInstTask).Error
	if instTaskErr != nil {
		hlog.Error("实例任务[%v]保存实例任务失败", execution.InstTaskID, instTaskErr)
		tx.Rollback()
		return instTaskErr
	}
	//保存实例节点任务
	if utils.IsNotEmptySlice(addInstNodeTasks) {
		addInstNodeTaskErr := tx.CreateInBatches(addInstNodeTasks, len(addInstNodeTasks)).Error
		if addInstNodeTaskErr != nil {
			hlog.Error("实例任务[%v]保存实例节点任务失败", execution.InstTaskID, addInstNodeTaskErr)
			tx.Rollback()
			return addInstNodeTaskErr
		}
	}
	if utils.IsNotEmptySlice(editInstNodeTasks) {
		for _, editInstNodeTask := range editInstNodeTasks {
			editInstNodeTaskErr := tx.Where("inst_task_id = ? and node_task_id = ?", execution.InstTaskID, editInstNodeTask.NodeTaskID).Updates(editInstNodeTask).Error
			if editInstNodeTaskErr != nil {
				hlog.Error("实例任务[%v]更新实例节点任务失败", execution.InstTaskID, editInstNodeTaskErr)
				tx.Rollback()
				return editInstNodeTaskErr
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
			tx.Rollback()
			return delInstNodeTaskErr
		}
	}
	//保存实例用户任务
	if utils.IsNotEmptySlice(addInstUserTasks) {
		addInstUserTaskErr := tx.CreateInBatches(addInstUserTasks, len(addInstUserTasks)).Error
		if addInstUserTaskErr != nil {
			hlog.Error("实例任务[%v]保存实例用户任务失败", execution.InstTaskID, addInstUserTaskErr)
			tx.Rollback()
			return addInstUserTaskErr
		}
	}
	if utils.IsNotEmptySlice(editInstUserTasks) {
		for _, editInstUserTask := range editInstUserTasks {
			editInstUserTaskErr := tx.Where("inst_task_id = ? and user_task_id = ?", execution.InstTaskID, editInstUserTask.UserTaskID).Updates(editInstUserTask).Error
			if editInstUserTaskErr != nil {
				hlog.Error("实例任务[%v]更新实例用户任务失败", execution.InstTaskID, editInstUserTaskErr)
				tx.Rollback()
				return editInstUserTaskErr
			}
		}
	}
	if utils.IsNotEmptySlice(delInstUserTasks) {
		delInstUserTaskIds := make([]string, 0)
		for _, delInstUserTask := range delInstUserTasks {
			delInstUserTaskIds = append(delInstUserTaskIds, delInstUserTask.NodeTaskID)
		}
		delInstUserTaskErr := tx.Where("inst_task_id = ? and user_task_id IN ?", execution.InstTaskID, delInstUserTaskIds).Delete(&model.InstUserTask{}).Error
		if delInstUserTaskErr != nil {
			hlog.Error("实例任务[%v]删除实例用户任务失败", execution.InstTaskID, delInstUserTaskErr)
			tx.Rollback()
			return delInstUserTaskErr
		}
	}
	//保存实例任务日志
	if utils.IsNotEmptySlice(addInstTaskOpLogs) {
		addInstTaskFormPerErr := tx.CreateInBatches(addInstTaskOpLogs, len(addInstTaskOpLogs)).Error
		if addInstTaskFormPerErr != nil {
			hlog.Error("实例任务[%v]保存实例任务日志失败", execution.InstTaskID, addInstTaskFormPerErr)
			tx.Rollback()
			return addInstTaskFormPerErr
		}
	}
	//保存实例节点任务表单权限
	if utils.IsNotEmptySlice(addInstTaskFormPers) {
		addInstTaskFormPerErr := tx.CreateInBatches(addInstTaskFormPers, len(addInstTaskFormPers)).Error
		if addInstTaskFormPerErr != nil {
			hlog.Error("实例任务[%v]保存实例任务表单权限失败", execution.InstTaskID, addInstTaskFormPerErr)
			tx.Rollback()
			return addInstTaskFormPerErr
		}
	}
	//保存实例任务参数
	if addInstTaskParams != nil && len(addInstTaskParams) > 0 {
		addInstTaskParamErr := tx.CreateInBatches(addInstTaskParams, len(addInstTaskParams)).Error
		if addInstTaskParamErr != nil {
			hlog.Error("实例任务[%v]保存实例任务参数失败", execution.InstTaskID, addInstTaskParamErr)
			tx.Rollback()
			return addInstTaskParamErr
		}
	}
	//提交事务
	tx.Commit()
	return nil
}

// execStopInstData
// @Description: 终止操作执行的实例数据，进行数据处理
// @receiver instTaskExecution
func (instTaskExecution *InstTaskExecution) execStopInstData() error {
	//开启事务
	tx := MysqlDB.Begin()
	execution := instTaskExecution.Execution
	//保存用户任务评论
	addInstUserTaskOpinion := &model.InstUserTaskOpinion{
		InstTaskID:  execution.InstTaskID,
		OpinionID:   snowflake.GetSnowflakeId(),
		Opinion:     int32(instTaskExecution.Opinion),
		OpinionDesc: instTaskExecution.OpinionDesc,
		OpUserID:    instTaskExecution.OpUserID,
		OpUserName:  instTaskExecution.OpUserName,
		CreateTime:  execution.Now,
		UpdateTime:  execution.Now,
		OpinionTime: execution.Now,
	}
	addInstUserTaskOpinionErr := tx.Create(addInstUserTaskOpinion).Error
	if addInstUserTaskOpinionErr != nil {
		hlog.Error("实例任务[%v]保存用户任务评论失败", execution.InstTaskID, addInstUserTaskOpinionErr)
		tx.Rollback()
		return addInstUserTaskOpinionErr
	}
	//修改实例任务状态为终止
	editInstTask := model.InstTaskDetail{
		Status:         constant.InstanceTaskStatusStop,
		EndTime:        execution.Now,
		UpdateTime:     execution.Now,
		UpdateUserID:   instTaskExecution.OpUserID,
		UpdateUserName: instTaskExecution.OpUserName,
	}
	editInstTaskErr := tx.Model(&model.InstTaskDetail{}).Where("inst_task_id = ?", execution.InstTaskID).Updates(editInstTask).Error
	if editInstTaskErr != nil {
		hlog.Error("实例任务[%v]修改实例任务失败", execution.InstTaskID, editInstTaskErr)
		tx.Rollback()
		return editInstTaskErr
	}
	//修改当前正在进行中的节点任务状态为终止
	editInstNodeTask := model.InstNodeTask{
		Status:     constant.InstanceNodeTaskStatusStop,
		UpdateTime: execution.Now,
	}
	editInstNodeTaskErr := tx.Model(&model.InstNodeTask{}).Where("inst_task_id = ? and status = ?", execution.InstTaskID, constant.InstanceNodeTaskStatusDoing).Updates(editInstNodeTask).Error
	if editInstNodeTaskErr != nil {
		hlog.Error("实例任务[%v]修改实例节点任务失败", execution.InstTaskID, editInstNodeTaskErr)
		tx.Rollback()
		return editInstNodeTaskErr
	}
	//修改当前正在进行中的用户任务状态为终止
	editInstUserTask := model.InstUserTask{
		Status:     constant.InstanceUserTaskStatusStop,
		HandleTime: execution.Now,
		UpdateTime: execution.Now,
	}
	editInstUserTaskErr := tx.Model(&model.InstUserTask{}).Where("inst_task_id = ? and status = ?", execution.InstTaskID, constant.InstanceUserTaskStatusDoing).Updates(editInstUserTask).Error
	if editInstUserTaskErr != nil {
		hlog.Error("实例任务[%v]修改实例用户任务失败", execution.InstTaskID, editInstUserTaskErr)
		tx.Rollback()
		return editInstUserTaskErr
	}
	//提交事务
	tx.Commit()
	return nil
}

// execSuspendInstData
// @Description: 挂起实例任务操作执行的实例数据，进行数据处理
// @receiver instTaskExecution
func (instTaskExecution *InstTaskExecution) execSuspendInstData() error {
	//开启事务
	tx := MysqlDB.Begin()
	execution := instTaskExecution.Execution
	//保存用户任务评论
	addInstUserTaskOpinion := &model.InstUserTaskOpinion{
		InstTaskID:  execution.InstTaskID,
		OpinionID:   snowflake.GetSnowflakeId(),
		Opinion:     int32(instTaskExecution.Opinion),
		OpinionDesc: instTaskExecution.OpinionDesc,
		OpUserID:    instTaskExecution.OpUserID,
		OpUserName:  instTaskExecution.OpUserName,
		CreateTime:  execution.Now,
		UpdateTime:  execution.Now,
		OpinionTime: execution.Now,
	}
	addInstUserTaskOpinionErr := tx.Create(addInstUserTaskOpinion).Error
	if addInstUserTaskOpinionErr != nil {
		hlog.Error("实例任务[%v]保存用户任务评论失败", execution.InstTaskID, addInstUserTaskOpinionErr)
		tx.Rollback()
		return addInstUserTaskOpinionErr
	}
	//修改实例任务状态为终止
	editInstTask := model.InstTaskDetail{
		Status:         constant.InstanceTaskStatusHangUp,
		UpdateTime:     execution.Now,
		UpdateUserID:   instTaskExecution.OpUserID,
		UpdateUserName: instTaskExecution.OpUserName,
	}
	editInstTaskErr := tx.Model(&model.InstTaskDetail{}).Where("inst_task_id = ?", execution.InstTaskID).Updates(editInstTask).Error
	if editInstTaskErr != nil {
		hlog.Error("实例任务[%v]修改实例任务失败", execution.InstTaskID, editInstTaskErr)
		tx.Rollback()
		return editInstTaskErr
	}
	//提交事务
	tx.Commit()
	return nil
}

// execResumeInstData
// @Description: 终止操作执行的实例数据，进行数据处理
// @receiver instTaskExecution
func (instTaskExecution *InstTaskExecution) execResumeInstData() error {
	//开启事务
	tx := MysqlDB.Begin()
	execution := instTaskExecution.Execution
	//保存用户任务评论
	addInstUserTaskOpinion := &model.InstUserTaskOpinion{
		InstTaskID:  execution.InstTaskID,
		OpinionID:   snowflake.GetSnowflakeId(),
		Opinion:     int32(instTaskExecution.Opinion),
		OpinionDesc: instTaskExecution.OpinionDesc,
		OpUserID:    instTaskExecution.OpUserID,
		OpUserName:  instTaskExecution.OpUserName,
		CreateTime:  execution.Now,
		UpdateTime:  execution.Now,
		OpinionTime: execution.Now,
	}
	addInstUserTaskOpinionErr := tx.Create(addInstUserTaskOpinion).Error
	if addInstUserTaskOpinionErr != nil {
		hlog.Error("实例任务[%v]保存用户任务评论失败", execution.InstTaskID, addInstUserTaskOpinionErr)
		tx.Rollback()
		return addInstUserTaskOpinionErr
	}
	//修改实例任务状态为终止
	editInstTask := model.InstTaskDetail{
		Status:         constant.InstanceTaskStatusDoing,
		UpdateTime:     execution.Now,
		UpdateUserID:   instTaskExecution.OpUserID,
		UpdateUserName: instTaskExecution.OpUserName,
	}
	editInstTaskErr := tx.Model(&model.InstTaskDetail{}).Where("inst_task_id = ?", execution.InstTaskID).Updates(editInstTask).Error
	if editInstTaskErr != nil {
		hlog.Error("实例任务[%v]修改实例任务失败", execution.InstTaskID, editInstTaskErr)
		tx.Rollback()
		return editInstTaskErr
	}
	//提交事务
	tx.Commit()
	return nil
}

// execDeleteInstData
// @Description: 删除操作执行的实例数据，进行数据处理
// @receiver instTaskExecution
// @return error
func (instTaskExecution *InstTaskExecution) execDeleteInstData() error {
	//开启事务
	tx := MysqlDB.Begin()
	execution := instTaskExecution.Execution
	//删除实例任务
	delInstTaskErr := tx.Where("inst_task_id = ?", execution.InstTaskID).Delete(&model.InstTaskDetail{}).Error
	if delInstTaskErr != nil {
		hlog.Error("实例任务[%v]删除实例任务失败", execution.InstTaskID, delInstTaskErr)
		tx.Rollback()
		return delInstTaskErr
	}
	//删除实例节点任务
	delInstNodeTaskErr := tx.Where("inst_task_id = ?", execution.InstTaskID).Delete(&model.InstNodeTask{}).Error
	if delInstNodeTaskErr != nil {
		hlog.Error("实例任务[%v]删除实例节点任务失败", execution.InstTaskID, delInstNodeTaskErr)
		tx.Rollback()
		return delInstNodeTaskErr
	}
	//删除实例用户任务
	delInstUserTaskErr := tx.Where("inst_task_id = ?", execution.InstTaskID).Delete(&model.InstUserTask{}).Error
	if delInstUserTaskErr != nil {
		hlog.Error("实例任务[%v]删除实例用户任务失败", execution.InstTaskID, delInstUserTaskErr)
		tx.Rollback()
		return delInstUserTaskErr
	}
	//删除实例任务参数
	delInstTaskParamErr := tx.Where("inst_task_id = ?", execution.InstTaskID).Delete(&model.InstTaskParam{}).Error
	if delInstTaskParamErr != nil {
		hlog.Error("实例任务[%v]删除实例任务参数失败", execution.InstTaskID, delInstTaskParamErr)
		tx.Rollback()
		return delInstTaskParamErr
	}
	//删除实例任务参数属性
	delInstTaskParamAttrErr := tx.Where("inst_task_id = ?", execution.InstTaskID).Delete(&model.InstTaskParamAttr{}).Error
	if delInstTaskParamAttrErr != nil {
		hlog.Error("实例任务[%v]删除实例任务参数属性失败", execution.InstTaskID, delInstTaskParamAttrErr)
		tx.Rollback()
		return delInstTaskParamAttrErr
	}
	//删除实例任务表单权限
	delInstTaskFormPerErr := tx.Where("inst_task_id = ?", execution.InstTaskID).Delete(&model.InstNodeTaskFormper{}).Error
	if delInstTaskFormPerErr != nil {
		hlog.Error("实例任务[%v]删除实例任务表单权限失败", execution.InstTaskID, delInstTaskFormPerErr)
		tx.Rollback()
		return delInstTaskFormPerErr
	}
	//删除实例任务日志
	delInstTaskOpLogErr := tx.Where("inst_task_id = ?", execution.InstTaskID).Delete(&model.InstTaskOpLog{}).Error
	if delInstTaskOpLogErr != nil {
		hlog.Error("实例任务[%v]删除实例任务日志失败", execution.InstTaskID, delInstTaskOpLogErr)
		tx.Rollback()
		return delInstTaskOpLogErr
	}
	//删除实例任务评论
	delInstUserTaskOpinionErr := tx.Where("inst_task_id = ?", execution.InstTaskID).Delete(&model.InstUserTaskOpinion{}).Error
	if delInstUserTaskOpinionErr != nil {
		hlog.Error("实例任务[%v]删除实例任务评论失败", execution.InstTaskID, delInstUserTaskOpinionErr)
		tx.Rollback()
		return delInstUserTaskOpinionErr
	}
	//提交事务
	tx.Commit()
	return nil
}

// execStartInstData
// @Description: 保存实例数据
// @receiver instTaskExecution
func (userTaskExecution *UserTaskExecution) execInstUserTaskData() error {
	execution := userTaskExecution.Execution
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
	//修改实例任务
	notifyErr := userTaskExecution.execInstTaskNotice(tx)
	if notifyErr != nil {
		hlog.Error("实例任务[%v]修改实例任务失败", execution.InstTaskID, notifyErr)
		tx.Rollback()
		return notifyErr
	}
	//保存用户任务评论
	addInstUserTaskOpinion := &model.InstUserTaskOpinion{
		InstTaskID:  execution.InstTaskID,
		NodeTaskID:  userTaskExecution.NodeTaskID,
		UserTaskID:  userTaskExecution.UserTaskID,
		NodeID:      userTaskExecution.NodeID,
		OpinionID:   snowflake.GetSnowflakeId(),
		Opinion:     int32(userTaskExecution.Opinion),
		OpinionDesc: userTaskExecution.OpinionDesc,
		OpUserID:    userTaskExecution.OpUserID,
		OpUserName:  userTaskExecution.OpUserName,
		CreateTime:  execution.Now,
		UpdateTime:  execution.Now,
		OpinionTime: execution.Now,
	}
	addInstUserTaskOpinionErr := tx.Create(addInstUserTaskOpinion).Error
	if addInstUserTaskOpinionErr != nil {
		hlog.Error("实例任务[%v]保存用户任务评论失败", execution.InstTaskID, addInstUserTaskOpinionErr)
		tx.Rollback()
		return addInstUserTaskOpinionErr
	}
	//保存实例节点任务
	if utils.IsNotEmptySlice(addInstNodeTasks) {
		addInstNodeTaskErr := tx.CreateInBatches(addInstNodeTasks, len(addInstNodeTasks)).Error
		if addInstNodeTaskErr != nil {
			hlog.Error("实例任务[%v]保存实例节点任务失败", execution.InstTaskID, addInstNodeTaskErr)
			tx.Rollback()
			return addInstNodeTaskErr
		}
	}
	if utils.IsNotEmptySlice(editInstNodeTasks) {
		for _, editInstNodeTask := range editInstNodeTasks {
			editInstNodeTaskErr := tx.Where("inst_task_id = ? and node_task_id = ?", execution.InstTaskID, editInstNodeTask.NodeTaskID).Updates(editInstNodeTask).Error
			if editInstNodeTaskErr != nil {
				hlog.Error("实例任务[%v]更新实例节点任务失败", execution.InstTaskID, editInstNodeTaskErr)
				tx.Rollback()
				return editInstNodeTaskErr
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
			tx.Rollback()
			return delInstNodeTaskErr
		}
	}
	//保存实例用户任务
	if utils.IsNotEmptySlice(addInstUserTasks) {
		addInstUserTaskErr := tx.CreateInBatches(addInstUserTasks, len(addInstUserTasks)).Error
		if addInstUserTaskErr != nil {
			hlog.Error("实例任务[%v]保存实例用户任务失败", execution.InstTaskID, addInstUserTaskErr)
			tx.Rollback()
			return addInstUserTaskErr
		}
	}
	if utils.IsNotEmptySlice(editInstUserTasks) {
		for _, editInstUserTask := range editInstUserTasks {
			editInstUserTaskErr := tx.Where("inst_task_id = ? and user_task_id = ?", execution.InstTaskID, editInstUserTask.UserTaskID).Updates(editInstUserTask).Error
			if editInstUserTaskErr != nil {
				hlog.Error("实例任务[%v]更新实例用户任务失败", execution.InstTaskID, editInstUserTaskErr)
				tx.Rollback()
				return editInstUserTaskErr
			}
		}
	}
	if utils.IsNotEmptySlice(delInstUserTasks) {
		delInstUserTaskIds := make([]string, 0)
		for _, delInstUserTask := range delInstUserTasks {
			delInstUserTaskIds = append(delInstUserTaskIds, delInstUserTask.NodeTaskID)
		}
		delInstUserTaskErr := tx.Where("inst_task_id = ? and user_task_id IN ?", execution.InstTaskID, delInstUserTaskIds).Delete(&model.InstUserTask{}).Error
		if delInstUserTaskErr != nil {
			hlog.Error("实例任务[%v]删除实例用户任务失败", execution.InstTaskID, delInstUserTaskErr)
			tx.Rollback()
			return delInstUserTaskErr
		}
	}
	//保存实例任务日志
	if utils.IsNotEmptySlice(addInstTaskOpLogs) {
		addInstTaskFormPerErr := tx.CreateInBatches(addInstTaskOpLogs, len(addInstTaskOpLogs)).Error
		if addInstTaskFormPerErr != nil {
			hlog.Error("实例任务[%v]保存实例任务日志失败", execution.InstTaskID, addInstTaskFormPerErr)
			tx.Rollback()
			return addInstTaskFormPerErr
		}
	}
	//保存实例节点任务表单权限
	if utils.IsNotEmptySlice(addInstTaskFormPers) {
		addInstTaskFormPerErr := tx.CreateInBatches(addInstTaskFormPers, len(addInstTaskFormPers)).Error
		if addInstTaskFormPerErr != nil {
			hlog.Error("实例任务[%v]保存实例任务表单权限失败", execution.InstTaskID, addInstTaskFormPerErr)
			tx.Rollback()
			return addInstTaskFormPerErr
		}
	}
	//保存实例任务参数
	if addInstTaskParams != nil && len(addInstTaskParams) > 0 {
		//删除实例任务参数
		delInstTaskParamErr := tx.Where("inst_task_id = ?", execution.InstTaskID).Delete(&model.InstTaskParam{}).Error
		if delInstTaskParamErr != nil {
			hlog.Error("实例任务[%v]保存实例任务参数失败", execution.InstTaskID, delInstTaskParamErr)
			tx.Rollback()
			return delInstTaskParamErr
		}
		addInstTaskParamErr := tx.CreateInBatches(addInstTaskParams, len(addInstTaskParams)).Error
		if addInstTaskParamErr != nil {
			hlog.Error("实例任务[%v]保存实例任务参数失败", execution.InstTaskID, addInstTaskParamErr)
			tx.Rollback()
			return addInstTaskParamErr
		}
	}
	//提交事务
	tx.Commit()
	return nil
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
			continue
		case constant.OperationTypeDelete:
			continue
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
