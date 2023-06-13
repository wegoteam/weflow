package service

import (
	"github.com/wegoteam/weflow/pkg/common/entity"
	"github.com/wegoteam/weflow/pkg/common/utils"
	"github.com/wegoteam/weflow/pkg/model"
)

// GetInstNodeUserTask
// @Description: 获取实例任务、节点任务、用户任务信息
// @param userTaskID 实例用户任务ID
// @return *entity.InstNodeAndUserTaskResult
func GetInstNodeUserTask(userTaskID string) *entity.InstNodeAndUserTaskResult {
	if utils.IsStrBlank(userTaskID) {
		panic("实例用户任务id不能为空")
	}
	var instNodeUserTask entity.InstNodeAndUserTaskResult
	MysqlDB.Raw("select"+
		" ut.id UID, ut.user_task_id UserTaskId, ut.type Type, ut.strategy Strategy, ut.node_user_name NodeUserName, ut.node_user_id NodeUserID, ut.sort Sort, ut.obj Obj, ut.relative Relative, ut.status UStatus, ut.create_time UCreateTime, ut.update_time UUpdateTime, ut.handle_time HandleTime, ut.op_user_id OpUserID, ut.op_user_name OpUserName, ut.opinion Opinion, ut.opinion_desc OpinionDesc,"+
		" nt.id NID, nt.node_task_id NodeTaskID, nt.node_id NodeID, nt.parent_id ParentID, nt.node_model NodeModel, nt.node_name NodeName, nt.approve_type ApproveType, nt.none_handler NoneHandler, nt.appoint_handler AppointHandler, nt.handle_mode HandleMode, nt.finish_mode FinishMode, nt.branch_mode BranchMode, nt.default_branch DefaultBranch, nt.branch_level BranchLevel, nt.condition_group ConditionGroup, nt.condition_expr ConditionExpr, nt.remark NRemark, nt.status NStatus, nt.create_time NCreateTime, nt.update_time NUpdateTime,"+
		" itd.id TID, itd.inst_task_id InstTaskID, itd.model_id ModelID, itd.process_def_id ProcessDefID, itd.form_def_id FormDefID, itd.version_id VersionID, itd.task_name TaskName, itd.status TStatus, itd.remark TRemark, itd.create_time TCreateTime, itd.create_user_id CreateUserID, itd.create_user_name CreateUserName, itd.update_time TUpdateTime, itd.update_user_id UpdateUserID, itd.update_user_name UpdateUserName, itd.start_time StartTime, itd.end_time EndTime"+
		" from inst_user_task ut"+
		" left join inst_node_task nt on ut.node_task_id = nt.node_task_id"+
		" left join inst_task_detail itd on nt.inst_task_id = itd.inst_task_id"+
		" where ut.user_task_id = ?", userTaskID).First(&instNodeUserTask)
	return &instNodeUserTask
}

// GetExecNodeTaskMap
// @Description: 获取实例任务的执行节点任务信息
// @param instTaskID
func GetExecNodeTaskMap(instTaskID string) map[string]entity.ExecNodeTaskBO {

	execNodeTaskMap := make(map[string]entity.ExecNodeTaskBO)
	if utils.IsStrBlank(instTaskID) {
		return execNodeTaskMap
	}
	var nodeTaskList []model.InstNodeTask
	MysqlDB.Where("inst_task_id = ?", instTaskID).Find(&nodeTaskList)
	if utils.IsEmptySlice(nodeTaskList) {
		return execNodeTaskMap
	}
	for _, nodeTask := range nodeTaskList {

		execNodeTaskMap[nodeTask.NodeID] = entity.ExecNodeTaskBO{
			NodeTaskID: nodeTask.NodeTaskID,
			NodeID:     nodeTask.NodeID,
			Status:     int8(nodeTask.Status),
			NodeModel:  int8(nodeTask.NodeModel),
		}
	}

	return execNodeTaskMap
}

// GetOpUserTasks
// @Description: 获取实例任务的执行节点任务信息
// @param instTaskID
// @param nodeTaskID
func GetOpUserTasks(instTaskID, nodeTaskID string) []entity.InstUserTaskResult {
	var userTaskResults = make([]entity.InstUserTaskResult, 0)
	userTasks := []model.InstUserTask{}
	MysqlDB.Where("inst_task_id = ? and node_task_id = ?", instTaskID, nodeTaskID).Find(&userTasks)
	if utils.IsEmptySlice(userTasks) {
		return userTaskResults
	}
	for _, userTask := range userTasks {
		instUserTaskResult := &entity.InstUserTaskResult{
			ID:           userTask.ID,
			InstTaskID:   userTask.InstTaskID,
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
		userTaskResults = append(userTaskResults, *instUserTaskResult)
	}
	return userTaskResults
}

// GetOpSortUserTasks
// @Description: 获取实例任务的执行节点任务信息
// @param instTaskID
// @param nodeTaskID
// @param opSort
// @return []entity.InstUserTaskResult
func GetOpSortUserTasks(instTaskID, nodeTaskID string, opSort int) []entity.InstUserTaskResult {
	var userTaskResults = make([]entity.InstUserTaskResult, 0)
	userTasks := []model.InstUserTask{}
	MysqlDB.Where("inst_task_id = ? and node_task_id = ? and sort = ? ", instTaskID, nodeTaskID, opSort).Find(&userTasks)
	if utils.IsEmptySlice(userTasks) {
		return userTaskResults
	}
	for _, userTask := range userTasks {
		instUserTaskResult := &entity.InstUserTaskResult{
			ID:           userTask.ID,
			InstTaskID:   userTask.InstTaskID,
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
		userTaskResults = append(userTaskResults, *instUserTaskResult)
	}
	return userTaskResults
}

// GetUserTaskMaxOpSort
// @Description: 当前节点任务依次审批的用户任务最大处理顺序
// @param currNodeModelBO
// @return int
func GetUserTaskMaxOpSort(currNodeModelBO *entity.NodeModelBO) int {
	if currNodeModelBO == nil {
		panic("当前的节点模型为空")
	}
	var maxOpSort int
	handlers := currNodeModelBO.NodeHandler.Handlers
	if utils.IsEmptySlice(handlers) {
		return maxOpSort
	}
	for _, handler := range handlers {
		if handler.Sort >= maxOpSort {
			maxOpSort = handler.Sort
		}
	}
	return maxOpSort
}
