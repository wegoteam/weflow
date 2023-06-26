package exec

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/wegoteam/weflow/pkg/common/constant"
	"github.com/wegoteam/weflow/pkg/model"
	"gorm.io/gorm"
)

//实例任务监听：发起、完成、挂起、恢复、终止、回滚

// execInstTaskNotice
// @Description: 执行实例任务监听
// @receiver userTaskExecution
// @param: db
// @return error
func (userTaskExecution *UserTaskExecution) execInstTaskNotice(tx *gorm.DB) error {
	execution := userTaskExecution.Execution
	if execution.InstTaskStatus == constant.InstanceTaskStatusDoing {
		return nil
	}
	//保存实例任务
	var editInstTask = &model.InstTaskDetail{
		Status:         int32(execution.InstTaskStatus),
		UpdateTime:     execution.Now,
		UpdateUserID:   userTaskExecution.OpUserID,
		UpdateUserName: userTaskExecution.OpUserName,
		EndTime:        execution.Now,
	}
	editInstTaskErr := tx.Where("inst_task_id = ?", execution.InstTaskID).Updates(editInstTask).Error
	if editInstTaskErr != nil {
		hlog.Error("实例任务[%v]更新实例任务失败", execution.InstTaskID, editInstTaskErr)
		return editInstTaskErr
	}
	return nil
}
