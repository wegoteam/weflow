package service

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/golang-module/carbon/v2"
	"github.com/pkg/errors"
	"github.com/wegoteam/weflow/pkg/common/entity"
	"github.com/wegoteam/weflow/pkg/common/utils"
	"github.com/wegoteam/weflow/pkg/model"
	"gorm.io/gorm"
)

// GetInstNodeUserTask
// @Description: 获取实例任务、节点任务、用户任务信息
// @param userTaskID 实例用户任务ID
// @return *entity.InstNodeAndUserTaskResult
func GetInstNodeUserTask(userTaskID string) *entity.InstNodeAndUserTaskResult {
	if utils.IsStrBlank(userTaskID) {
		return nil
	}
	var instNodeUserTask entity.InstNodeAndUserTaskResult
	MysqlDB.Raw("select"+
		" ut.id UID, ut.user_task_id UserTaskID, ut.type Type, ut.strategy Strategy, ut.node_user_name NodeUserName, ut.node_user_id NodeUserID, ut.sort Sort, ut.obj Obj, ut.relative Relative, ut.status UStatus, ut.create_time UCreateTime, ut.update_time UUpdateTime, ut.handle_time HandleTime, ut.op_user_id OpUserID, ut.op_user_name OpUserName, ut.opinion Opinion, ut.opinion_desc OpinionDesc,"+
		" nt.id NID, nt.node_task_id NodeTaskID, nt.node_id NodeID, nt.parent_id ParentID, nt.node_model NodeModel, nt.node_name NodeName, nt.approve_type ApproveType, nt.none_handler NoneHandler, nt.appoint_handler AppointHandler, nt.handle_mode HandleMode, nt.finish_mode FinishMode, nt.branch_mode BranchMode, nt.default_branch DefaultBranch, nt.branch_level BranchLevel, nt.condition_group ConditionGroup, nt.condition_expr ConditionExpr, nt.remark NRemark, nt.status NStatus, nt.create_time NCreateTime, nt.update_time NUpdateTime,"+
		" itd.id TID, itd.inst_task_id InstTaskID, itd.model_id ModelID, itd.process_def_id ProcessDefID, itd.form_def_id FormDefID, itd.version_id VersionID, itd.task_name TaskName, itd.status TStatus, itd.remark TRemark, itd.create_time TCreateTime, itd.create_user_id CreateUserID, itd.create_user_name CreateUserName, itd.update_time TUpdateTime, itd.update_user_id UpdateUserID, itd.update_user_name UpdateUserName, itd.start_time StartTime, itd.end_time EndTime"+
		" from inst_user_task ut"+
		" left join inst_node_task nt on ut.node_task_id = nt.node_task_id"+
		" left join inst_task_detail itd on nt.inst_task_id = itd.inst_task_id"+
		" where ut.user_task_id = ?", userTaskID).First(&instNodeUserTask)
	return &instNodeUserTask
}

// GetTodoUserTasks
// @Description: 获取待办用户任务
// @param userID
// @return *[]entity.InstNodeAndUserTaskResult
func GetTodoUserTasks(userID string) *[]entity.InstNodeAndUserTaskResult {
	//var userTasks []entity.InstNodeAndUserTaskResult
	//MysqlDB.Raw("select"+
	//	" ut.id UID, ut.user_task_id UserTaskID, ut.type Type, ut.strategy Strategy, ut.node_user_name NodeUserName, ut.node_user_id NodeUserID, ut.sort Sort, ut.obj Obj, ut.relative Relative, ut.status UStatus, ut.create_time UCreateTime, ut.update_time UUpdateTime, ut.handle_time HandleTime, ut.op_user_id OpUserID, ut.op_user_name OpUserName, ut.opinion Opinion, ut.opinion_desc OpinionDesc,"+
	//	" nt.id NID, nt.node_task_id NodeTaskID, nt.node_id NodeID, nt.parent_id ParentID, nt.node_model NodeModel, nt.node_name NodeName, nt.approve_type ApproveType, nt.none_handler NoneHandler, nt.appoint_handler AppointHandler, nt.handle_mode HandleMode, nt.finish_mode FinishMode, nt.branch_mode BranchMode, nt.default_branch DefaultBranch, nt.branch_level BranchLevel, nt.condition_group ConditionGroup, nt.condition_expr ConditionExpr, nt.remark NRemark, nt.status NStatus, nt.create_time NCreateTime, nt.update_time NUpdateTime,"+
	//	" itd.id TID, itd.inst_task_id InstTaskID, itd.model_id ModelID, itd.process_def_id ProcessDefID, itd.form_def_id FormDefID, itd.version_id VersionID, itd.task_name TaskName, itd.status TStatus, itd.remark TRemark, itd.create_time TCreateTime, itd.create_user_id CreateUserID, itd.create_user_name CreateUserName, itd.update_time TUpdateTime, itd.update_user_id UpdateUserID, itd.update_user_name UpdateUserName, itd.start_time StartTime, itd.end_time EndTime"+
	//	" from inst_user_task ut"+
	//	" left join inst_node_task nt on ut.node_task_id = nt.node_task_id"+
	//	" left join inst_task_detail itd on nt.inst_task_id = itd.inst_task_id"+
	//	" where ((ut.status = 1 and nt.status = 2 and itd.status = 2) or (ut.status = 1  and nt.node_model = 4)) and ut.op_user_id = ?"+
	//	" order by ut.create_time desc", userID).Find(&userTasks)
	var userTasks []entity.InstNodeAndUserTaskResult
	tx := MysqlDB.Model(&model.InstUserTask{}).
		Select(" inst_user_task.id UID,inst_user_task.user_task_id UserTaskID, inst_user_task.type Type, inst_user_task.strategy Strategy, inst_user_task.node_user_name NodeUserName, inst_user_task.node_user_id NodeUserID, inst_user_task.sort Sort, inst_user_task.obj Obj, inst_user_task.relative Relative, inst_user_task.status UStatus, inst_user_task.create_time UCreateTime, inst_user_task.update_time UUpdateTime, inst_user_task.handle_time HandleTime, inst_user_task.op_user_id OpUserID, inst_user_task.op_user_name OpUserName, inst_user_task.opinion Opinion, inst_user_task.opinion_desc OpinionDesc," +
			" inst_node_task.id NID, inst_node_task.node_task_id NodeTaskID, inst_node_task.node_id NodeID, inst_node_task.parent_id ParentID, inst_node_task.node_model NodeModel, inst_node_task.node_name NodeName, inst_node_task.approve_type ApproveType, inst_node_task.none_handler NoneHandler, inst_node_task.appoint_handler AppointHandler, inst_node_task.handle_mode HandleMode, inst_node_task.finish_mode FinishMode, inst_node_task.branch_mode BranchMode, inst_node_task.default_branch DefaultBranch, inst_node_task.branch_level BranchLevel, inst_node_task.condition_group ConditionGroup, inst_node_task.condition_expr ConditionExpr, inst_node_task.remark NRemark, inst_node_task.status NStatus, inst_node_task.create_time NCreateTime, inst_node_task.update_time NUpdateTime," +
			" inst_task_detail.id TID, inst_task_detail.inst_task_id InstTaskID, inst_task_detail.model_id ModelID, inst_task_detail.process_def_id ProcessDefID, inst_task_detail.form_def_id FormDefID, inst_task_detail.version_id VersionID, inst_task_detail.task_name TaskName, inst_task_detail.status TStatus, inst_task_detail.remark TRemark, inst_task_detail.create_time TCreateTime, inst_task_detail.create_user_id CreateUserID, inst_task_detail.create_user_name CreateUserName, inst_task_detail.update_time TUpdateTime, inst_task_detail.update_user_id UpdateUserID, inst_task_detail.update_user_name UpdateUserName, inst_task_detail.start_time StartTime, inst_task_detail.end_time EndTime").
		Joins("left join inst_node_task  on inst_user_task.node_task_id = inst_node_task.node_task_id").
		Joins("left join inst_task_detail on inst_node_task.inst_task_id = inst_task_detail.inst_task_id")
	tx.Where("(inst_user_task.status = 1 and inst_node_task.status = 2 and inst_task_detail.status = 2) or (inst_user_task.status = 1  and inst_node_task.node_model = 4)")
	tx.Where("inst_user_task.op_user_id = ?", userID).Order("inst_user_task.create_time desc").Find(&userTasks)
	return &userTasks
}

// QueryTodoUserTasks
// @Description: 查询待办任务
// @param: userID
// @return *[]entity.InstNodeAndUserTaskResult
func QueryTodoUserTasks(param *entity.UserTaskQueryBO) (*[]entity.InstNodeAndUserTaskResult, error) {
	var userTasks []entity.InstNodeAndUserTaskResult
	tx := MysqlDB.Model(&model.InstUserTask{}).
		Select(" inst_user_task.id UID,inst_user_task.user_task_id UserTaskID, inst_user_task.type Type, inst_user_task.strategy Strategy, inst_user_task.node_user_name NodeUserName, inst_user_task.node_user_id NodeUserID, inst_user_task.sort Sort, inst_user_task.obj Obj, inst_user_task.relative Relative, inst_user_task.status UStatus, inst_user_task.create_time UCreateTime, inst_user_task.update_time UUpdateTime, inst_user_task.handle_time HandleTime, inst_user_task.op_user_id OpUserID, inst_user_task.op_user_name OpUserName, inst_user_task.opinion Opinion, inst_user_task.opinion_desc OpinionDesc," +
			" inst_node_task.id NID, inst_node_task.node_task_id NodeTaskID, inst_node_task.node_id NodeID, inst_node_task.parent_id ParentID, inst_node_task.node_model NodeModel, inst_node_task.node_name NodeName, inst_node_task.approve_type ApproveType, inst_node_task.none_handler NoneHandler, inst_node_task.appoint_handler AppointHandler, inst_node_task.handle_mode HandleMode, inst_node_task.finish_mode FinishMode, inst_node_task.branch_mode BranchMode, inst_node_task.default_branch DefaultBranch, inst_node_task.branch_level BranchLevel, inst_node_task.condition_group ConditionGroup, inst_node_task.condition_expr ConditionExpr, inst_node_task.remark NRemark, inst_node_task.status NStatus, inst_node_task.create_time NCreateTime, inst_node_task.update_time NUpdateTime," +
			" inst_task_detail.id TID, inst_task_detail.inst_task_id InstTaskID, inst_task_detail.model_id ModelID, inst_task_detail.process_def_id ProcessDefID, inst_task_detail.form_def_id FormDefID, inst_task_detail.version_id VersionID, inst_task_detail.task_name TaskName, inst_task_detail.status TStatus, inst_task_detail.remark TRemark, inst_task_detail.create_time TCreateTime, inst_task_detail.create_user_id CreateUserID, inst_task_detail.create_user_name CreateUserName, inst_task_detail.update_time TUpdateTime, inst_task_detail.update_user_id UpdateUserID, inst_task_detail.update_user_name UpdateUserName, inst_task_detail.start_time StartTime, inst_task_detail.end_time EndTime").
		Joins("left join inst_node_task  on inst_user_task.node_task_id = inst_node_task.node_task_id").
		Joins("left join inst_task_detail on inst_node_task.inst_task_id = inst_task_detail.inst_task_id").Scopes(BuildUserTaskQuery(param))
	tx.Where("(inst_user_task.status = 1 and inst_node_task.status = 2 and inst_task_detail.status = 2) or (inst_user_task.status = 1  and inst_node_task.node_model = 4)")
	err := tx.Where("inst_user_task.op_user_id = ?", param.UserID).Order("inst_user_task.create_time desc").Find(&userTasks).Error
	if err != nil {
		hlog.Errorf("查询待办用户任务失败:%s", err.Error())
		return nil, errors.New("查询待办用户任务失败")
	}
	return &userTasks, nil
}

// PageTodoUserTasks
// @Description: 分页待办用户任务
// @param: userID
// @return *entity.Page[entity.InstNodeAndUserTaskResult]
func PageTodoUserTasks(param *entity.UserTaskQueryBO) (*entity.Page[entity.InstNodeAndUserTaskResult], error) {
	//查询总数
	var total int64
	tx := MysqlDB.Model(&model.InstUserTask{}).
		Joins("left join inst_node_task  on inst_user_task.node_task_id = inst_node_task.node_task_id").
		Joins("left join inst_task_detail on inst_node_task.inst_task_id = inst_task_detail.inst_task_id").Scopes(BuildUserTaskQuery(param))
	tx.Where("(inst_user_task.status = 1 and inst_node_task.status = 2 and inst_task_detail.status = 2) or (inst_user_task.status = 1  and inst_node_task.node_model = 4)")
	err := tx.Where("inst_user_task.op_user_id = ?", param.UserID).Count(&total).Error
	if err != nil {
		hlog.Errorf("查询待办用户任务失败:%s", err.Error())
		return nil, errors.New("查询待办用户任务失败")
	}
	if total == 0 {
		return &entity.Page[entity.InstNodeAndUserTaskResult]{
			Total:    total,
			Records:  []entity.InstNodeAndUserTaskResult{},
			PageNum:  param.PageNum,
			PageSize: param.PageSize,
		}, nil
	}
	//分页查询
	var userTasks []entity.InstNodeAndUserTaskResult
	tx2 := MysqlDB.Model(&model.InstUserTask{}).
		Select(" inst_user_task.id UID,inst_user_task.user_task_id UserTaskID, inst_user_task.type Type, inst_user_task.strategy Strategy, inst_user_task.node_user_name NodeUserName, inst_user_task.node_user_id NodeUserID, inst_user_task.sort Sort, inst_user_task.obj Obj, inst_user_task.relative Relative, inst_user_task.status UStatus, inst_user_task.create_time UCreateTime, inst_user_task.update_time UUpdateTime, inst_user_task.handle_time HandleTime, inst_user_task.op_user_id OpUserID, inst_user_task.op_user_name OpUserName, inst_user_task.opinion Opinion, inst_user_task.opinion_desc OpinionDesc,"+
			" inst_node_task.id NID, inst_node_task.node_task_id NodeTaskID, inst_node_task.node_id NodeID, inst_node_task.parent_id ParentID, inst_node_task.node_model NodeModel, inst_node_task.node_name NodeName, inst_node_task.approve_type ApproveType, inst_node_task.none_handler NoneHandler, inst_node_task.appoint_handler AppointHandler, inst_node_task.handle_mode HandleMode, inst_node_task.finish_mode FinishMode, inst_node_task.branch_mode BranchMode, inst_node_task.default_branch DefaultBranch, inst_node_task.branch_level BranchLevel, inst_node_task.condition_group ConditionGroup, inst_node_task.condition_expr ConditionExpr, inst_node_task.remark NRemark, inst_node_task.status NStatus, inst_node_task.create_time NCreateTime, inst_node_task.update_time NUpdateTime,"+
			" inst_task_detail.id TID, inst_task_detail.inst_task_id InstTaskID, inst_task_detail.model_id ModelID, inst_task_detail.process_def_id ProcessDefID, inst_task_detail.form_def_id FormDefID, inst_task_detail.version_id VersionID, inst_task_detail.task_name TaskName, inst_task_detail.status TStatus, inst_task_detail.remark TRemark, inst_task_detail.create_time TCreateTime, inst_task_detail.create_user_id CreateUserID, inst_task_detail.create_user_name CreateUserName, inst_task_detail.update_time TUpdateTime, inst_task_detail.update_user_id UpdateUserID, inst_task_detail.update_user_name UpdateUserName, inst_task_detail.start_time StartTime, inst_task_detail.end_time EndTime").
		Joins("left join inst_node_task  on inst_user_task.node_task_id = inst_node_task.node_task_id").
		Joins("left join inst_task_detail on inst_node_task.inst_task_id = inst_task_detail.inst_task_id").Scopes(entity.Paginate(param.PageNum, param.PageSize), BuildUserTaskQuery(param))
	tx2.Where("(inst_user_task.status = 1 and inst_node_task.status = 2 and inst_task_detail.status = 2) or (inst_user_task.status = 1  and inst_node_task.node_model = 4)")
	err2 := tx2.Where("inst_user_task.op_user_id = ?", param.UserID).Order("inst_user_task.create_time desc").Find(&userTasks).Error
	if err2 != nil {
		hlog.Errorf("查询待办用户任务失败:%s", err2.Error())
		return nil, errors.New("查询待办用户任务失败")
	}
	//返回分页数据
	return &entity.Page[entity.InstNodeAndUserTaskResult]{
		Total:    total,
		Records:  userTasks,
		PageNum:  param.PageNum,
		PageSize: param.PageSize,
	}, nil
}

// BuildUserTaskQuery
// @Description: 用户任务查询条件
// @param: param
// @return func(db *gorm.DB) *gorm.DB
func BuildUserTaskQuery(param *entity.UserTaskQueryBO) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		tx := db
		if param.InstStatus != 0 {
			tx = db.Where("inst_node_task.status = ?", param.InstStatus)
		}
		if utils.IsStrNotBlank(param.TaskName) {
			tx = db.Where("inst_task_detail.task_name = ?", param.TaskName)
		}
		if utils.IsStrNotBlank(param.CreateUserId) {
			tx = db.Where("inst_task_detail.create_user_id = ?", param.CreateUserId)
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

// GetDoneUserTasks
// @Description: 获取已办用户任务
// @param userID 用户ID
// @return *[]entity.InstNodeAndUserTaskResult
func GetDoneUserTasks(userID string) (*[]entity.InstNodeAndUserTaskResult, error) {
	var userTasks []entity.InstNodeAndUserTaskResult
	tx := MysqlDB.Model(&model.InstUserTask{}).
		Select(" inst_user_task.id UID,inst_user_task.user_task_id UserTaskID, inst_user_task.type Type, inst_user_task.strategy Strategy, inst_user_task.node_user_name NodeUserName, inst_user_task.node_user_id NodeUserID, inst_user_task.sort Sort, inst_user_task.obj Obj, inst_user_task.relative Relative, inst_user_task.status UStatus, inst_user_task.create_time UCreateTime, inst_user_task.update_time UUpdateTime, inst_user_task.handle_time HandleTime, inst_user_task.op_user_id OpUserID, inst_user_task.op_user_name OpUserName, inst_user_task.opinion Opinion, inst_user_task.opinion_desc OpinionDesc," +
			" inst_node_task.id NID, inst_node_task.node_task_id NodeTaskID, inst_node_task.node_id NodeID, inst_node_task.parent_id ParentID, inst_node_task.node_model NodeModel, inst_node_task.node_name NodeName, inst_node_task.approve_type ApproveType, inst_node_task.none_handler NoneHandler, inst_node_task.appoint_handler AppointHandler, inst_node_task.handle_mode HandleMode, inst_node_task.finish_mode FinishMode, inst_node_task.branch_mode BranchMode, inst_node_task.default_branch DefaultBranch, inst_node_task.branch_level BranchLevel, inst_node_task.condition_group ConditionGroup, inst_node_task.condition_expr ConditionExpr, inst_node_task.remark NRemark, inst_node_task.status NStatus, inst_node_task.create_time NCreateTime, inst_node_task.update_time NUpdateTime," +
			" inst_task_detail.id TID, inst_task_detail.inst_task_id InstTaskID, inst_task_detail.model_id ModelID, inst_task_detail.process_def_id ProcessDefID, inst_task_detail.form_def_id FormDefID, inst_task_detail.version_id VersionID, inst_task_detail.task_name TaskName, inst_task_detail.status TStatus, inst_task_detail.remark TRemark, inst_task_detail.create_time TCreateTime, inst_task_detail.create_user_id CreateUserID, inst_task_detail.create_user_name CreateUserName, inst_task_detail.update_time TUpdateTime, inst_task_detail.update_user_id UpdateUserID, inst_task_detail.update_user_name UpdateUserName, inst_task_detail.start_time StartTime, inst_task_detail.end_time EndTime").
		Joins("left join inst_node_task  on inst_user_task.node_task_id = inst_node_task.node_task_id").
		Joins("left join inst_task_detail on inst_node_task.inst_task_id = inst_task_detail.inst_task_id")
	tx.Where("inst_user_task.status in (2,3,4,5)")
	err := tx.Where("inst_user_task.op_user_id = ?", userID).Order("inst_user_task.create_time desc").Find(&userTasks).Error
	if err != nil {
		hlog.Errorf("查询已办用户任务失败 error：%s", err.Error())
		return nil, errors.New("分页查询已办用户任务失败")
	}
	return &userTasks, nil
}

// QueryDoneUserTasks
// @Description: 查询已办用户任务
// @param: userID 用户ID
// @param: param 查询参数
// @return *[]entity.InstNodeAndUserTaskResult
func QueryDoneUserTasks(param *entity.UserTaskQueryBO) (*[]entity.InstNodeAndUserTaskResult, error) {
	var userTasks []entity.InstNodeAndUserTaskResult
	tx := MysqlDB.Model(&model.InstUserTask{}).
		Select(" inst_user_task.id UID,inst_user_task.user_task_id UserTaskID, inst_user_task.type Type, inst_user_task.strategy Strategy, inst_user_task.node_user_name NodeUserName, inst_user_task.node_user_id NodeUserID, inst_user_task.sort Sort, inst_user_task.obj Obj, inst_user_task.relative Relative, inst_user_task.status UStatus, inst_user_task.create_time UCreateTime, inst_user_task.update_time UUpdateTime, inst_user_task.handle_time HandleTime, inst_user_task.op_user_id OpUserID, inst_user_task.op_user_name OpUserName, inst_user_task.opinion Opinion, inst_user_task.opinion_desc OpinionDesc," +
			" inst_node_task.id NID, inst_node_task.node_task_id NodeTaskID, inst_node_task.node_id NodeID, inst_node_task.parent_id ParentID, inst_node_task.node_model NodeModel, inst_node_task.node_name NodeName, inst_node_task.approve_type ApproveType, inst_node_task.none_handler NoneHandler, inst_node_task.appoint_handler AppointHandler, inst_node_task.handle_mode HandleMode, inst_node_task.finish_mode FinishMode, inst_node_task.branch_mode BranchMode, inst_node_task.default_branch DefaultBranch, inst_node_task.branch_level BranchLevel, inst_node_task.condition_group ConditionGroup, inst_node_task.condition_expr ConditionExpr, inst_node_task.remark NRemark, inst_node_task.status NStatus, inst_node_task.create_time NCreateTime, inst_node_task.update_time NUpdateTime," +
			" inst_task_detail.id TID, inst_task_detail.inst_task_id InstTaskID, inst_task_detail.model_id ModelID, inst_task_detail.process_def_id ProcessDefID, inst_task_detail.form_def_id FormDefID, inst_task_detail.version_id VersionID, inst_task_detail.task_name TaskName, inst_task_detail.status TStatus, inst_task_detail.remark TRemark, inst_task_detail.create_time TCreateTime, inst_task_detail.create_user_id CreateUserID, inst_task_detail.create_user_name CreateUserName, inst_task_detail.update_time TUpdateTime, inst_task_detail.update_user_id UpdateUserID, inst_task_detail.update_user_name UpdateUserName, inst_task_detail.start_time StartTime, inst_task_detail.end_time EndTime").
		Joins("left join inst_node_task  on inst_user_task.node_task_id = inst_node_task.node_task_id").
		Joins("left join inst_task_detail on inst_node_task.inst_task_id = inst_task_detail.inst_task_id").Scopes(BuildUserTaskQuery(param))
	tx.Where("inst_user_task.status in (2,3,4,5)")
	err := tx.Where("inst_user_task.op_user_id = ?", param.UserID).Order("inst_user_task.create_time desc").Find(&userTasks).Error
	if err != nil {
		hlog.Errorf("查询已办用户任务失败 error：%s", err.Error())
		return nil, errors.New("分页查询已办用户任务失败")
	}
	return &userTasks, nil
}

// PageDoneUserTasks
// @Description: 分页查询已办用户任务
// @param: userID 用户ID
// @param: param 查询参数
// @return *[]entity.InstNodeAndUserTaskResult
func PageDoneUserTasks(param *entity.UserTaskQueryBO) (*entity.Page[entity.InstNodeAndUserTaskResult], error) {
	//查询总数
	var total int64
	tx := MysqlDB.Model(&model.InstUserTask{}).
		Joins("left join inst_node_task  on inst_user_task.node_task_id = inst_node_task.node_task_id").
		Joins("left join inst_task_detail on inst_node_task.inst_task_id = inst_task_detail.inst_task_id").Scopes(BuildUserTaskQuery(param))
	tx.Where("inst_user_task.status in (2,3,4,5)")
	err := tx.Where("inst_user_task.op_user_id = ?", param.UserID).Order("inst_user_task.create_time desc").Count(&total).Error
	if err != nil {
		hlog.Errorf("查询已办用户任务失败 error：%s", err.Error())
		return nil, errors.New("分页查询已办用户任务失败")
	}
	if total == 0 {
		return &entity.Page[entity.InstNodeAndUserTaskResult]{
			Total:    total,
			Records:  []entity.InstNodeAndUserTaskResult{},
			PageNum:  param.PageNum,
			PageSize: param.PageSize,
		}, nil
	}
	var userTasks []entity.InstNodeAndUserTaskResult
	tx2 := MysqlDB.Model(&model.InstUserTask{}).
		Select(" inst_user_task.id UID,inst_user_task.user_task_id UserTaskID, inst_user_task.type Type, inst_user_task.strategy Strategy, inst_user_task.node_user_name NodeUserName, inst_user_task.node_user_id NodeUserID, inst_user_task.sort Sort, inst_user_task.obj Obj, inst_user_task.relative Relative, inst_user_task.status UStatus, inst_user_task.create_time UCreateTime, inst_user_task.update_time UUpdateTime, inst_user_task.handle_time HandleTime, inst_user_task.op_user_id OpUserID, inst_user_task.op_user_name OpUserName, inst_user_task.opinion Opinion, inst_user_task.opinion_desc OpinionDesc,"+
			" inst_node_task.id NID, inst_node_task.node_task_id NodeTaskID, inst_node_task.node_id NodeID, inst_node_task.parent_id ParentID, inst_node_task.node_model NodeModel, inst_node_task.node_name NodeName, inst_node_task.approve_type ApproveType, inst_node_task.none_handler NoneHandler, inst_node_task.appoint_handler AppointHandler, inst_node_task.handle_mode HandleMode, inst_node_task.finish_mode FinishMode, inst_node_task.branch_mode BranchMode, inst_node_task.default_branch DefaultBranch, inst_node_task.branch_level BranchLevel, inst_node_task.condition_group ConditionGroup, inst_node_task.condition_expr ConditionExpr, inst_node_task.remark NRemark, inst_node_task.status NStatus, inst_node_task.create_time NCreateTime, inst_node_task.update_time NUpdateTime,"+
			" inst_task_detail.id TID, inst_task_detail.inst_task_id InstTaskID, inst_task_detail.model_id ModelID, inst_task_detail.process_def_id ProcessDefID, inst_task_detail.form_def_id FormDefID, inst_task_detail.version_id VersionID, inst_task_detail.task_name TaskName, inst_task_detail.status TStatus, inst_task_detail.remark TRemark, inst_task_detail.create_time TCreateTime, inst_task_detail.create_user_id CreateUserID, inst_task_detail.create_user_name CreateUserName, inst_task_detail.update_time TUpdateTime, inst_task_detail.update_user_id UpdateUserID, inst_task_detail.update_user_name UpdateUserName, inst_task_detail.start_time StartTime, inst_task_detail.end_time EndTime").
		Joins("left join inst_node_task  on inst_user_task.node_task_id = inst_node_task.node_task_id").
		Joins("left join inst_task_detail on inst_node_task.inst_task_id = inst_task_detail.inst_task_id").Scopes(entity.Paginate(param.PageNum, param.PageSize), BuildUserTaskQuery(param))
	tx2.Where("inst_user_task.status in (2,3,4,5)")
	err2 := tx2.Where("inst_user_task.op_user_id = ?", param.UserID).Order("inst_user_task.create_time desc").Find(&userTasks).Error
	if err2 != nil {
		hlog.Errorf("查询已办用户任务失败 error：%s", err2.Error())
		return nil, errors.New("分页查询已办用户任务失败")
	}
	//返回分页数据
	page := &entity.Page[entity.InstNodeAndUserTaskResult]{
		Total:    total,
		Records:  userTasks,
		PageNum:  param.PageNum,
		PageSize: param.PageSize,
	}
	return page, nil
}

// GetReceivedUserTasks
// @Description: 查询已办用户任务
// @param: userID 用户ID
// @return *[]entity.InstNodeAndUserTaskResult
func GetReceivedUserTasks(userID string) (*[]entity.InstNodeAndUserTaskResult, error) {
	var userTasks []entity.InstNodeAndUserTaskResult
	tx := MysqlDB.Model(&model.InstUserTask{}).
		Select(" inst_user_task.id UID,inst_user_task.user_task_id UserTaskID, inst_user_task.type Type, inst_user_task.strategy Strategy, inst_user_task.node_user_name NodeUserName, inst_user_task.node_user_id NodeUserID, inst_user_task.sort Sort, inst_user_task.obj Obj, inst_user_task.relative Relative, inst_user_task.status UStatus, inst_user_task.create_time UCreateTime, inst_user_task.update_time UUpdateTime, inst_user_task.handle_time HandleTime, inst_user_task.op_user_id OpUserID, inst_user_task.op_user_name OpUserName, inst_user_task.opinion Opinion, inst_user_task.opinion_desc OpinionDesc," +
			" inst_node_task.id NID, inst_node_task.node_task_id NodeTaskID, inst_node_task.node_id NodeID, inst_node_task.parent_id ParentID, inst_node_task.node_model NodeModel, inst_node_task.node_name NodeName, inst_node_task.approve_type ApproveType, inst_node_task.none_handler NoneHandler, inst_node_task.appoint_handler AppointHandler, inst_node_task.handle_mode HandleMode, inst_node_task.finish_mode FinishMode, inst_node_task.branch_mode BranchMode, inst_node_task.default_branch DefaultBranch, inst_node_task.branch_level BranchLevel, inst_node_task.condition_group ConditionGroup, inst_node_task.condition_expr ConditionExpr, inst_node_task.remark NRemark, inst_node_task.status NStatus, inst_node_task.create_time NCreateTime, inst_node_task.update_time NUpdateTime," +
			" inst_task_detail.id TID, inst_task_detail.inst_task_id InstTaskID, inst_task_detail.model_id ModelID, inst_task_detail.process_def_id ProcessDefID, inst_task_detail.form_def_id FormDefID, inst_task_detail.version_id VersionID, inst_task_detail.task_name TaskName, inst_task_detail.status TStatus, inst_task_detail.remark TRemark, inst_task_detail.create_time TCreateTime, inst_task_detail.create_user_id CreateUserID, inst_task_detail.create_user_name CreateUserName, inst_task_detail.update_time TUpdateTime, inst_task_detail.update_user_id UpdateUserID, inst_task_detail.update_user_name UpdateUserName, inst_task_detail.start_time StartTime, inst_task_detail.end_time EndTime").
		Joins("left join inst_node_task  on inst_user_task.node_task_id = inst_node_task.node_task_id").
		Joins("left join inst_task_detail on inst_node_task.inst_task_id = inst_task_detail.inst_task_id")
	err := tx.Where("inst_user_task.op_user_id = ?", userID).Order("inst_user_task.create_time desc").Find(&userTasks).Error
	if err != nil {
		hlog.Errorf("查询我收到的用户任务失败 error=%v", err.Error())
		return nil, err
	}
	return &userTasks, nil
}

// QueryReceivedUserTasks
// @Description: 查询我收到的用户任务
// @param: userID 用户ID
// @param: param 查询参数
// @return *[]entity.InstNodeAndUserTaskResult
func QueryReceivedUserTasks(param *entity.UserTaskQueryBO) (*[]entity.InstNodeAndUserTaskResult, error) {
	var userTasks []entity.InstNodeAndUserTaskResult
	tx := MysqlDB.Model(&model.InstUserTask{}).
		Select(" inst_user_task.id UID,inst_user_task.user_task_id UserTaskID, inst_user_task.type Type, inst_user_task.strategy Strategy, inst_user_task.node_user_name NodeUserName, inst_user_task.node_user_id NodeUserID, inst_user_task.sort Sort, inst_user_task.obj Obj, inst_user_task.relative Relative, inst_user_task.status UStatus, inst_user_task.create_time UCreateTime, inst_user_task.update_time UUpdateTime, inst_user_task.handle_time HandleTime, inst_user_task.op_user_id OpUserID, inst_user_task.op_user_name OpUserName, inst_user_task.opinion Opinion, inst_user_task.opinion_desc OpinionDesc," +
			" inst_node_task.id NID, inst_node_task.node_task_id NodeTaskID, inst_node_task.node_id NodeID, inst_node_task.parent_id ParentID, inst_node_task.node_model NodeModel, inst_node_task.node_name NodeName, inst_node_task.approve_type ApproveType, inst_node_task.none_handler NoneHandler, inst_node_task.appoint_handler AppointHandler, inst_node_task.handle_mode HandleMode, inst_node_task.finish_mode FinishMode, inst_node_task.branch_mode BranchMode, inst_node_task.default_branch DefaultBranch, inst_node_task.branch_level BranchLevel, inst_node_task.condition_group ConditionGroup, inst_node_task.condition_expr ConditionExpr, inst_node_task.remark NRemark, inst_node_task.status NStatus, inst_node_task.create_time NCreateTime, inst_node_task.update_time NUpdateTime," +
			" inst_task_detail.id TID, inst_task_detail.inst_task_id InstTaskID, inst_task_detail.model_id ModelID, inst_task_detail.process_def_id ProcessDefID, inst_task_detail.form_def_id FormDefID, inst_task_detail.version_id VersionID, inst_task_detail.task_name TaskName, inst_task_detail.status TStatus, inst_task_detail.remark TRemark, inst_task_detail.create_time TCreateTime, inst_task_detail.create_user_id CreateUserID, inst_task_detail.create_user_name CreateUserName, inst_task_detail.update_time TUpdateTime, inst_task_detail.update_user_id UpdateUserID, inst_task_detail.update_user_name UpdateUserName, inst_task_detail.start_time StartTime, inst_task_detail.end_time EndTime").
		Joins("left join inst_node_task  on inst_user_task.node_task_id = inst_node_task.node_task_id").
		Joins("left join inst_task_detail on inst_node_task.inst_task_id = inst_task_detail.inst_task_id").Scopes(BuildUserTaskQuery(param))
	err := tx.Where("inst_user_task.op_user_id = ?", param.UserID).Order("inst_user_task.create_time desc").Find(&userTasks).Error
	if err != nil {
		hlog.Errorf("查询我收到的用户任务失败 error=%v", err.Error())
		return nil, err
	}
	return &userTasks, nil
}

// PageReceivedUserTasks
// @Description: 分页查询我收到的用户任务
// @param: userID 用户ID
// @param: param 查询参数
// @return *entity.Page[entity.InstNodeAndUserTaskResult]
func PageReceivedUserTasks(param *entity.UserTaskQueryBO) (*entity.Page[entity.InstNodeAndUserTaskResult], error) {
	//查询总数
	var total int64
	tx := MysqlDB.Model(&model.InstUserTask{}).
		Joins("left join inst_node_task  on inst_user_task.node_task_id = inst_node_task.node_task_id").
		Joins("left join inst_task_detail on inst_node_task.inst_task_id = inst_task_detail.inst_task_id").Scopes(BuildUserTaskQuery(param))
	err := tx.Where("inst_user_task.op_user_id = ?", param.UserID).Order("inst_user_task.create_time desc").Count(&total).Error
	if err != nil {
		hlog.Errorf("查询我收到的用户任务失败 error=%v", err.Error())
		return nil, errors.New("查询我收到的用户任务失败")
	}
	if total == 0 {
		return &entity.Page[entity.InstNodeAndUserTaskResult]{
			Total:    total,
			Records:  []entity.InstNodeAndUserTaskResult{},
			PageNum:  param.PageNum,
			PageSize: param.PageSize,
		}, nil
	}
	var userTasks []entity.InstNodeAndUserTaskResult
	tx2 := MysqlDB.Model(&model.InstUserTask{}).
		Select(" inst_user_task.id UID,inst_user_task.user_task_id UserTaskID, inst_user_task.type Type, inst_user_task.strategy Strategy, inst_user_task.node_user_name NodeUserName, inst_user_task.node_user_id NodeUserID, inst_user_task.sort Sort, inst_user_task.obj Obj, inst_user_task.relative Relative, inst_user_task.status UStatus, inst_user_task.create_time UCreateTime, inst_user_task.update_time UUpdateTime, inst_user_task.handle_time HandleTime, inst_user_task.op_user_id OpUserID, inst_user_task.op_user_name OpUserName, inst_user_task.opinion Opinion, inst_user_task.opinion_desc OpinionDesc,"+
			" inst_node_task.id NID, inst_node_task.node_task_id NodeTaskID, inst_node_task.node_id NodeID, inst_node_task.parent_id ParentID, inst_node_task.node_model NodeModel, inst_node_task.node_name NodeName, inst_node_task.approve_type ApproveType, inst_node_task.none_handler NoneHandler, inst_node_task.appoint_handler AppointHandler, inst_node_task.handle_mode HandleMode, inst_node_task.finish_mode FinishMode, inst_node_task.branch_mode BranchMode, inst_node_task.default_branch DefaultBranch, inst_node_task.branch_level BranchLevel, inst_node_task.condition_group ConditionGroup, inst_node_task.condition_expr ConditionExpr, inst_node_task.remark NRemark, inst_node_task.status NStatus, inst_node_task.create_time NCreateTime, inst_node_task.update_time NUpdateTime,"+
			" inst_task_detail.id TID, inst_task_detail.inst_task_id InstTaskID, inst_task_detail.model_id ModelID, inst_task_detail.process_def_id ProcessDefID, inst_task_detail.form_def_id FormDefID, inst_task_detail.version_id VersionID, inst_task_detail.task_name TaskName, inst_task_detail.status TStatus, inst_task_detail.remark TRemark, inst_task_detail.create_time TCreateTime, inst_task_detail.create_user_id CreateUserID, inst_task_detail.create_user_name CreateUserName, inst_task_detail.update_time TUpdateTime, inst_task_detail.update_user_id UpdateUserID, inst_task_detail.update_user_name UpdateUserName, inst_task_detail.start_time StartTime, inst_task_detail.end_time EndTime").
		Joins("left join inst_node_task  on inst_user_task.node_task_id = inst_node_task.node_task_id").
		Joins("left join inst_task_detail on inst_node_task.inst_task_id = inst_task_detail.inst_task_id").Scopes(entity.Paginate(param.PageNum, param.PageSize), BuildUserTaskQuery(param))
	err2 := tx2.Where("inst_user_task.op_user_id = ?", param.UserID).Order("inst_user_task.create_time desc").Find(&userTasks).Error
	if err2 != nil {
		hlog.Errorf("查询我收到的用户任务失败 error=%v", err2.Error())
		return nil, errors.New("查询我收到的用户任务失败")
	}
	//返回分页数据
	page := &entity.Page[entity.InstNodeAndUserTaskResult]{
		Total:    total,
		Records:  userTasks,
		PageNum:  param.PageNum,
		PageSize: param.PageSize,
	}
	return page, nil
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
