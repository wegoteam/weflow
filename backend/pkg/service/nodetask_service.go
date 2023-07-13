package service

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/pkg/errors"
	"github.com/wegoteam/weflow/pkg/common/entity"
	"github.com/wegoteam/weflow/pkg/common/utils"
	"github.com/wegoteam/weflow/pkg/model"
)

// GetInstNodeTasks
// @Description: 获取实例节点任务
// @param: instTaskID 实例任务ID
// @return []entity.InstNodeTaskResult
// @return error
func GetInstNodeTasks(instTaskID string) ([]entity.InstNodeTaskResult, error) {
	var nodeTasks = []model.InstNodeTask{}
	nodeTaskErr := MysqlDB.Where("inst_task_id = ?", instTaskID).Find(&nodeTasks).Error
	if nodeTaskErr != nil {
		hlog.Errorf("查询节点任务失败：%s", nodeTaskErr.Error())
		return nil, errors.New("查询节点任务失败")
	}
	var nodeTaskResults = make([]entity.InstNodeTaskResult, 0)
	if utils.IsEmptySlice(nodeTasks) {
		return nodeTaskResults, nil
	}
	for _, nodeTask := range nodeTasks {
		var nodeTaskResult = entity.InstNodeTaskResult{
			ID:             nodeTask.ID,
			InstTaskID:     nodeTask.InstTaskID,
			NodeTaskID:     nodeTask.NodeTaskID,
			NodeID:         nodeTask.NodeID,
			ParentID:       nodeTask.ParentID,
			NodeModel:      nodeTask.NodeModel,
			NodeName:       nodeTask.NodeName,
			ApproveType:    nodeTask.ApproveType,
			NoneHandler:    nodeTask.NoneHandler,
			AppointHandler: nodeTask.AppointHandler,
			HandleMode:     nodeTask.HandleMode,
			FinishMode:     nodeTask.FinishMode,
			BranchMode:     nodeTask.BranchMode,
			DefaultBranch:  nodeTask.DefaultBranch,
			BranchLevel:    nodeTask.BranchLevel,
			ConditionGroup: nodeTask.ConditionGroup,
			ConditionExpr:  nodeTask.ConditionExpr,
			Remark:         nodeTask.Remark,
			Status:         nodeTask.Status,
			CreateTime:     nodeTask.CreateTime,
			UpdateTime:     nodeTask.UpdateTime,
		}
		nodeTaskResults = append(nodeTaskResults, nodeTaskResult)
	}
	return nodeTaskResults, nil
}

// GetInstTaskFormPers
// @Description: 获取节点任务表单权限
// @param: instTaskID 实例任务ID
// @param: nodeID 节点ID
// @return []entity.InstNodeTaskFormperResult
// @return error
func GetInstTaskFormPers(instTaskID string) ([]entity.InstNodeTaskFormperResult, error) {
	var taskFormpers = []model.InstNodeTaskFormper{}
	taskFormperErr := MysqlDB.Where("inst_task_id = ?", instTaskID).Find(&taskFormpers).Error
	if taskFormperErr != nil {
		hlog.Errorf("查询节点任务表单权限失败：%s", taskFormperErr.Error())
		return nil, errors.New("查询节点任务表单权限失败")
	}
	var taskFormperResults = make([]entity.InstNodeTaskFormperResult, 0)
	if utils.IsEmptySlice(taskFormpers) {
		return taskFormperResults, nil
	}
	for _, taskFormper := range taskFormpers {
		var taskFormperResult = entity.InstNodeTaskFormperResult{
			ID:         taskFormper.ID,
			InstTaskID: taskFormper.InstTaskID,
			NodeTaskID: taskFormper.NodeTaskID,
			NodeID:     taskFormper.NodeID,
			ElemID:     taskFormper.ElemID,
			ElemPID:    taskFormper.ElemPID,
			Per:        taskFormper.Per,
		}
		taskFormperResults = append(taskFormperResults, taskFormperResult)
	}
	return taskFormperResults, nil
}

// GetInstNodeTaskFormPers
// @Description: 获取节点任务表单权限
// @param: instTaskID 实例任务ID
// @param: nodeID 节点ID
// @return []entity.InstNodeTaskFormperResult
// @return error
func GetInstNodeTaskFormPers(instTaskID, nodeID string) ([]entity.InstNodeTaskFormperResult, error) {
	var taskFormpers = []model.InstNodeTaskFormper{}
	taskFormperErr := MysqlDB.Where("inst_task_id = ? and node_id = ?", instTaskID, nodeID).Find(&taskFormpers).Error
	if taskFormperErr != nil {
		hlog.Errorf("查询节点任务表单权限失败：%s", taskFormperErr.Error())
		return nil, errors.New("查询节点任务表单权限失败")
	}
	var taskFormperResults = make([]entity.InstNodeTaskFormperResult, 0)
	if utils.IsEmptySlice(taskFormpers) {
		return taskFormperResults, nil
	}
	for _, taskFormper := range taskFormpers {
		var taskFormperResult = entity.InstNodeTaskFormperResult{
			ID:         taskFormper.ID,
			InstTaskID: taskFormper.InstTaskID,
			NodeTaskID: taskFormper.NodeTaskID,
			NodeID:     taskFormper.NodeID,
			ElemID:     taskFormper.ElemID,
			ElemPID:    taskFormper.ElemPID,
			Per:        taskFormper.Per,
		}
		taskFormperResults = append(taskFormperResults, taskFormperResult)
	}
	return taskFormperResults, nil
}
