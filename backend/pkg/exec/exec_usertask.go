package exec

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/pkg/errors"
	"github.com/wegoteam/weflow/pkg/common/constant"
	"github.com/wegoteam/weflow/pkg/common/entity"
	"github.com/wegoteam/weflow/pkg/common/utils"
	"github.com/wegoteam/weflow/pkg/model"
	"github.com/wegoteam/weflow/pkg/service"
	"github.com/wegoteam/wepkg/id/snowflake"
)

// Agree
// @Description: 同意
// @param: userTaskID 用户任务ID
// @param: opUserID 操作用户ID
// @param: OpUserName 操作用户名称
// @param: opinionDesc 意见描述
// @param: params 参数
// @return error
func Agree(userTaskID, opUserID, opUserName, opinionDesc string, params map[string]any) error {
	//创建用户任务执行对象
	userTaskExecution, err := NewUserTaskExecution(userTaskID)
	if err != nil {
		return err
	}
	//验证用户任务信息
	verifyErr := userTaskExecution.verify(opUserID)
	if verifyErr != nil {
		return verifyErr
	}
	userTaskExecution.OpUserID = opUserID
	userTaskExecution.OpUserName = opUserName
	userTaskExecution.OpinionDesc = opinionDesc
	userTaskExecution.UserTaskStatus = constant.InstanceUserTaskStatusAgree
	userTaskExecution.Opinion = constant.InstanceUserTaskOpinionAgree
	return userTaskExecution.agree(userTaskID, params)
}

// Save
// @Description: 保存
// @param: userTaskID 用户任务ID
// @param: opUserID 操作用户ID
// @param: OpUserName 操作用户名称
// @param: opinionDesc 意见描述
// @param: params 参数
// @return error
func Save(userTaskID, opUserID, opUserName, opinionDesc string, params map[string]any) error {
	//创建用户任务执行对象
	userTaskExecution, err := NewUserTaskExecution(userTaskID)
	if err != nil {
		return err
	}
	//验证用户任务信息
	verifyErr := userTaskExecution.verify(opUserID)
	if verifyErr != nil {
		return verifyErr
	}
	userTaskExecution.OpUserID = opUserID
	userTaskExecution.OpUserName = opUserName
	userTaskExecution.OpinionDesc = opinionDesc
	userTaskExecution.Opinion = constant.InstanceUserTaskOpinionSave
	return userTaskExecution.save(userTaskID, params)
}

// Disagree
// @Description: 不同意
// @param: userTaskID 用户任务ID
// @param: opUserID 操作用户ID
// @param: OpUserName 操作用户名称
// @param: opinionDesc 意见描述
// @return error
func Disagree(userTaskID, opUserID, opUserName, opinionDesc string) error {
	//创建用户任务执行对象
	userTaskExecution, err := NewUserTaskExecution(userTaskID)
	if err != nil {
		return err
	}
	//验证用户任务信息
	verifyErr := userTaskExecution.verify(opUserID)
	if verifyErr != nil {
		return verifyErr
	}
	userTaskExecution.OpUserID = opUserID
	userTaskExecution.OpUserName = opUserName
	userTaskExecution.OpinionDesc = opinionDesc
	userTaskExecution.UserTaskStatus = constant.InstanceUserTaskStatusDisagree
	userTaskExecution.Opinion = constant.InstanceUserTaskOpinionDisagree
	return userTaskExecution.disagree(userTaskID)
}

// Rollback
// @Description: 回退上节点
// @param: userTaskID 用户任务ID
// @param: opUserID 操作用户ID
// @param: opUserName 操作用户名称
// @param: opinionDesc 意见描述
// @return error
func Rollback(userTaskID, opUserID, opUserName, opinionDesc string) error {
	userTaskExecution, err := NewUserTaskExecution(userTaskID)
	if err != nil {
		return err
	}
	//验证用户任务信息
	verifyErr := userTaskExecution.verify(opUserID)
	if verifyErr != nil {
		return verifyErr
	}
	userTaskExecution.OpUserID = opUserID
	userTaskExecution.OpUserName = opUserName
	userTaskExecution.OpinionDesc = opinionDesc
	userTaskExecution.UserTaskStatus = constant.InstanceUserTaskStatusRollback
	userTaskExecution.Opinion = constant.InstanceUserTaskOpinionRollback
	return userTaskExecution.rollback(userTaskID)
}

// RollbackStartNode
// @Description: 回退到开始节点
// @param: userTaskID 用户任务ID
// @param: opUserID 操作用户ID
// @param: opUserName 操作用户名称
// @param: opinionDesc 意见描述
// @return error
func RollbackStartNode(userTaskID, opUserID, opUserName, opinionDesc string) error {
	userTaskExecution, err := NewUserTaskExecution(userTaskID)
	if err != nil {
		return err
	}
	//验证用户任务信息
	verifyErr := userTaskExecution.verify(opUserID)
	if verifyErr != nil {
		return verifyErr
	}
	userTaskExecution.OpUserID = opUserID
	userTaskExecution.OpUserName = opUserName
	userTaskExecution.OpinionDesc = opinionDesc
	userTaskExecution.UserTaskStatus = constant.InstanceUserTaskStatusRollback
	userTaskExecution.Opinion = constant.InstanceUserTaskOpinionRollback
	return userTaskExecution.rollback(userTaskID)
}

// agree
// @Description: 同意
//保存实例任务参数、修改当前用户任务、修改当前节点任务、添加用户任务评论、保存数据
// @receiver userTaskExecution
// @param: userTaskID 用户任务ID
// @param: params 参数
// @return bool
func (userTaskExecution *UserTaskExecution) agree(userTaskID string, params map[string]any) error {
	execution := userTaskExecution.Execution
	execution.InstTaskParamMap = params
	//执行流转
	execErr := execNodeTask(userTaskExecution)
	if execErr != nil {
		return execErr
	}
	//修改当前用户任务
	userTasks := execution.UserTasks
	editUserTask := &entity.UserTaskBO{
		ExecOpType:  constant.OperationTypeUpdate,
		UserTaskID:  userTaskID,
		Status:      constant.InstanceUserTaskStatusAgree,
		Opinion:     int32(userTaskExecution.Opinion),
		OpinionDesc: userTaskExecution.OpinionDesc,
		UpdateTime:  execution.Now,
	}
	*userTasks = append(*userTasks, *editUserTask)
	//执行数据
	err := userTaskExecution.execInstUserTaskData()
	if err != nil {
		return err
	}
	hlog.Infof("实例任务[%s]的当前节点任务[%s]同意操作执行成功", execution.InstTaskID, userTaskID)
	return nil
}

// save
// @Description: 保存
// @receiver userTaskExecution
// @param: userTaskID 用户任务ID
// @param: params 参数
// @return bool
func (userTaskExecution *UserTaskExecution) save(userTaskID string, params map[string]any) error {
	execution := userTaskExecution.Execution
	execution.InstTaskParamMap = params

	//开启事务
	tx := MysqlDB.Begin()
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
	//转换实例任务参数
	addInstTaskParams := service.TransformInstTaskParam(execution.InstTaskID, execution.InstTaskParamMap, execution.Now)
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
	tx.Commit()
	hlog.Infof("实例任务[%s]的当前节点任务[%s]保存操作执行成功", execution.InstTaskID, userTaskID)
	return nil
}

// disagree
// @Description: 不同意
// @receiver userTaskExecution
// @param: userTaskID
// @param: params
// @return bool
func (userTaskExecution *UserTaskExecution) disagree(userTaskID string) error {
	execution := userTaskExecution.Execution
	execErr := execNodeTask(userTaskExecution)
	if execErr != nil {
		return execErr
	}
	//修改当前用户任务
	userTasks := execution.UserTasks
	editUserTask := &entity.UserTaskBO{
		ExecOpType:  constant.OperationTypeUpdate,
		UserTaskID:  userTaskID,
		Status:      constant.InstanceUserTaskStatusDisagree,
		OpinionDesc: userTaskExecution.OpinionDesc,
		UpdateTime:  execution.Now,
	}
	*userTasks = append(*userTasks, *editUserTask)
	//执行数据
	err := userTaskExecution.execInstUserTaskData()
	if err != nil {
		return err
	}
	hlog.Infof("实例任务[%s]的当前节点任务[%s]不同意操作执行成功", execution.InstTaskID, userTaskID)
	return nil
}

// turn
// @Description: 转办任务，将任务交接给他人办理，办理完成后继续下步骤
// @receiver userTaskExecution
// @return bool
func (userTaskExecution *UserTaskExecution) turn(userTaskID string) error {
	execution := userTaskExecution.Execution

	//修改当前用户任务
	userTasks := execution.UserTasks
	editUserTask := &entity.UserTaskBO{
		ExecOpType:  constant.OperationTypeUpdate,
		UserTaskID:  userTaskID,
		Status:      constant.InstanceUserTaskStatusDisagree,
		OpinionDesc: userTaskExecution.OpinionDesc,
		UpdateTime:  execution.Now,
	}
	*userTasks = append(*userTasks, *editUserTask)
	//执行数据
	err := userTaskExecution.execInstUserTaskData()
	if err != nil {
		return err
	}
	//开启事务
	tx := MysqlDB.Begin()

	tx.Commit()
	return nil
}

// delegate
// @Description: 委托任务，将任务委托给他人，他人办理完成后再回到委托人
// @receiver userTaskExecution
// @return bool
func (userTaskExecution *UserTaskExecution) delegate(userTaskID string) error {

	return nil
}

// rollback
// @Description: 回退上节点
// @receiver userTaskExecution
// @return bool
func (userTaskExecution *UserTaskExecution) rollback(userTaskID string) error {
	execution := userTaskExecution.Execution
	//修改当前用户任务
	userTasks := execution.UserTasks
	editUserTask := &entity.UserTaskBO{
		ExecOpType:  constant.OperationTypeUpdate,
		UserTaskID:  userTaskID,
		Status:      constant.InstanceUserTaskStatusRollback,
		OpinionDesc: userTaskExecution.OpinionDesc,
		UpdateTime:  execution.Now,
	}
	*userTasks = append(*userTasks, *editUserTask)
	//执行数据
	err := userTaskExecution.execInstUserTaskData()
	if err != nil {
		return err
	}
	return nil
}

// rollbackStartNode
// @Description: 回退发起节点
// @receiver userTaskExecution
// @return bool
func (userTaskExecution *UserTaskExecution) rollbackStartNode(userTaskID string) error {
	execution := userTaskExecution.Execution
	//修改当前用户任务
	editUserTask := &model.InstUserTask{
		Status:      constant.InstanceUserTaskStatusRollback,
		OpinionDesc: userTaskExecution.OpinionDesc,
		UpdateTime:  execution.Now,
	}

	//开启事务
	tx := MysqlDB.Begin()
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
	editInstUserTaskErr := tx.Where("inst_task_id = ? and user_task_id = ?", execution.InstTaskID, userTaskExecution.UserTaskID).Updates(editUserTask).Error
	if editInstUserTaskErr != nil {
		hlog.Error("实例任务[%v]更新实例用户任务失败", execution.InstTaskID, editInstUserTaskErr)
		tx.Rollback()
		return editInstUserTaskErr
	}
	//修改实例任务状态为终止
	editInstTask := model.InstTaskDetail{
		Status:         constant.InstanceTaskStatusRollback,
		UpdateTime:     execution.Now,
		UpdateUserID:   userTaskExecution.OpUserID,
		UpdateUserName: userTaskExecution.OpUserName,
	}
	editInstTaskErr := tx.Where("inst_task_id = ?", execution.InstTaskID).Updates(editInstTask).Error
	if editInstTaskErr != nil {
		hlog.Error("实例任务[%v]修改实例任务失败", execution.InstTaskID, editInstTaskErr)
		tx.Rollback()
		return editInstTaskErr
	}
	tx.Commit()
	return nil
}

// rollbackAnyNode
// @Description: 回退任意节点
// @receiver userTaskExecution
// @return bool
func (userTaskExecution *UserTaskExecution) rollbackAnyNode(userTaskID, nodeID string) error {
	execution := userTaskExecution.Execution
	//修改当前用户任务
	userTasks := execution.UserTasks
	editUserTask := &entity.UserTaskBO{
		ExecOpType:  constant.OperationTypeUpdate,
		UserTaskID:  userTaskID,
		Status:      constant.InstanceUserTaskStatusRollback,
		OpinionDesc: userTaskExecution.OpinionDesc,
		UpdateTime:  execution.Now,
	}
	*userTasks = append(*userTasks, *editUserTask)
	//执行数据
	err := userTaskExecution.execInstUserTaskData()
	if err != nil {
		return err
	}
	return nil
}

// revoke
// @Description: 撤回，处理人撤回
// @receiver userTaskExecution
// @return bool
func (userTaskExecution *UserTaskExecution) revoke(userTaskID string) error {

	return nil
}

// cancel
// @Description: 撤销：发起人撤销
// @receiver userTaskExecution
// @return bool
func (userTaskExecution *UserTaskExecution) cancel(userTaskID string) error {

	return nil
}

// urge
// @Description: 催办
// @receiver userTaskExecution
// @return bool
func (userTaskExecution *UserTaskExecution) urge(userTaskID string) error {

	return nil
}

// addSign
// @Description: 加签
// @receiver userTaskExecution
// @return bool
func (userTaskExecution *UserTaskExecution) addSign(userTaskID string) error {

	return nil
}

// reduceSign
// @Description: 减签
// @receiver userTaskExecution
// @return bool
func (userTaskExecution *UserTaskExecution) reduceSign(userTaskID string) error {

	return nil
}

// cc
// @Description: 抄送
// @receiver userTaskExecution
// @return bool
func (userTaskExecution *UserTaskExecution) cc(userTaskID string) error {

	return nil
}

// ccReply
// @Description: 抄送回复
// @receiver userTaskExecution
// @return bool
func (userTaskExecution *UserTaskExecution) ccReply(userTaskID string) error {

	return nil
}

// ccRevoke
// @Description: 抄送撤回
// @receiver userTaskExecution
// @return bool
func (userTaskExecution *UserTaskExecution) ccRevoke(userTaskID string) error {

	return nil
}

// verify
// @Description: 验证用户任务信息
// @receiver userTaskExecution
func (userTaskExecution *UserTaskExecution) verify(opUserID string) error {
	//用户任务状态是否是进行中
	if userTaskExecution.UserTaskStatus != constant.InstanceUserTaskStatusDoing {
		hlog.Errorf("当前用户任务[%s]状态不是进行中", userTaskExecution.UserTaskID)
		return errors.New("当前用户任务状态不是进行中")
	}
	//用户任务操作员与提交处理人任务操作员是否一致
	if userTaskExecution.OpUserID != opUserID {
		hlog.Errorf("当前用户[%s]无权限操作该任务[%s]", opUserID, userTaskExecution.UserTaskID)
		return errors.New("当前用户无权限操作该任务")
	}
	if userTaskExecution.NodeModel == constant.NotifyNodeModel {
		return nil
	}
	//实例任务状态是否是进行中
	if userTaskExecution.Execution.InstTaskStatus != constant.InstanceTaskStatusDoing {
		hlog.Errorf("当前实例任务[%s]状态不是进行中", userTaskExecution.Execution.InstTaskID)
		return errors.New("当前实例任务状态不是进行中")
	}
	//节点任务状态是否是进行中、节点任务权限模式是否是知会
	if userTaskExecution.NodeTaskStatus != constant.InstanceNodeTaskStatusDoing {
		hlog.Errorf("当前节点任务[%s]状态不是进行中", userTaskExecution.UserTaskID)
		return errors.New("当前节点任务状态不是进行中")
	}
	return nil
}

// execUserTask
// @Description: 执行用户任务，串行任务，并行任务，会签、或签
//审批方式【依次审批：1；会签（需要完成人数的审批人同意或拒绝才可完成节点）：2；或签（其中一名审批人同意或拒绝即可）：3】默认会签2
// @param: userTaskExecution
func execNodeTask(userTaskExecution *UserTaskExecution) error {
	//判断当前的节点任务是否完成
	finishFlag := isFinish(userTaskExecution)
	if !finishFlag {
		//执行数据
		err := userTaskExecution.execInstUserTaskData()
		if err != nil {
			return err
		}
		hlog.Infof("当前节点任务[%s]同意操作，节点任务未完成", userTaskExecution.UserTaskID)
		return nil
	}
	execution := userTaskExecution.Execution
	//验证节点任务：依次审批、会签、或签；取决于流转用户任务还是流转节点任务
	//执行流转节点获取是流转任务
	processDefModel := execution.ProcessDefModel
	currNodeModelBO, ok := processDefModel.NodeModelMap[userTaskExecution.NodeID]
	if !ok {
		hlog.Errorf("当前节点[%s]不存在", userTaskExecution.NodeID)
		return errors.New("当前节点不存在")
	}
	//修改当前节点任务的执行
	execNodeTaskBO := entity.ExecNodeTaskBO{
		NodeTaskID: userTaskExecution.NodeTaskID,
		NodeID:     userTaskExecution.NodeID,
		NodeModel:  int8(userTaskExecution.NodeModel),
		Status:     int8(userTaskExecution.NodeTaskStatus),
	}
	execution.ExecNodeTaskMap[userTaskExecution.NodeID] = execNodeTaskBO
	//修改当前节点任务
	nodeTasks := execution.InstNodeTasks
	editNodeTask := &entity.InstNodeTaskBO{
		ExecOpType: constant.OperationTypeUpdate,
		NodeTaskID: userTaskExecution.NodeTaskID,
		Status:     int32(userTaskExecution.NodeTaskStatus),
		UpdateTime: execution.Now,
	}
	*nodeTasks = append(*nodeTasks, *editNodeTask)
	//依次审批
	if userTaskExecution.HandleMode == constant.ApprovalWayOrder {
		//判断当前节点任务依次审批的用户任务最大处理顺序
		maxOpSort := service.GetUserTaskMaxOpSort(&currNodeModelBO)
		if userTaskExecution.OpSort != maxOpSort {
			//执行下任务
			execNextTask(&currNodeModelBO, userTaskExecution)
		}
	}
	//抄送任务不继续流转
	if userTaskExecution.NodeModel == constant.NotifyNodeModel {
		return nil
	}
	//当前节点任务不通过，且是顶层节点的时候，流程结束；否则流转至分支节点决策
	if userTaskExecution.NodeTaskStatus == constant.InstanceNodeTaskStatusNotPass {
		if isParent(userTaskExecution.ParentID) {
			execution.InstTaskStatus = constant.InstanceTaskStatusStop
		} else {
			pNodeModelBO, pok := processDefModel.NodeModelMap[userTaskExecution.ParentID]
			if !pok {
				hlog.Errorf("当前节点[%s]的父节点不存在", userTaskExecution.NodeID)
				return errors.New("当前节点的父节点不存在")
			}
			execNode(&pNodeModelBO, execution)
		}
		return nil
	}
	//执行下一个节点
	execNextNode(&currNodeModelBO, execution)
	return nil
}

// isFinish
// @Description: 判断当前的节点任务是否完成
//审批方式【依次审批：1；会签（需要完成人数的审批人同意或拒绝才可完成节点）：2；或签（其中一名审批人同意或拒绝即可）：3】默认会签2
// @return bool
func isFinish(userTaskExecution *UserTaskExecution) bool {
	switch userTaskExecution.HandleMode {
	case constant.ApprovalWayOrder:
		return execApprovalWayOrder(userTaskExecution)
	case constant.ApprovalWayCount:
		return execApprovalWayCount(userTaskExecution)
	case constant.ApprovalWayOr:
		return execApprovalWayOr(userTaskExecution)
	default:
		return false
	}
}

// execApprovalWayOrder
// @Description: 依次审批默认0所有人不可选人，所有人依次审批
func execApprovalWayOrder(userTaskExecution *UserTaskExecution) bool {
	if userTaskExecution.Opinion == constant.InstanceUserTaskOpinionDisagree {
		userTaskExecution.NodeTaskStatus = constant.InstanceNodeTaskStatusNotPass
		return true
	}
	userTasks := service.GetOpSortUserTasks(userTaskExecution.Execution.InstTaskID, userTaskExecution.NodeTaskID, userTaskExecution.OpSort)
	if utils.IsEmptySlice(userTasks) {
		return false
	}
	var finishCount = 0
	for _, userTask := range userTasks {
		if userTask.UserTaskID == userTaskExecution.UserTaskID {
			continue
		}
		if userTask.Status == constant.InstanceUserTaskStatusAgree {
			finishCount++
		}
	}
	if userTaskExecution.Opinion == constant.InstanceUserTaskOpinionAgree {
		finishCount++
	}
	//当前节点完成
	if finishCount == len(userTasks) {
		userTaskExecution.NodeTaskStatus = constant.InstanceNodeTaskStatusComplete
		return true
	}
	return false
}

// execApprovalWayCount
// @Description: 会签默认0所有人（可选人大于0），需要所有人同意
func execApprovalWayCount(userTaskExecution *UserTaskExecution) bool {
	if userTaskExecution.Opinion == constant.InstanceUserTaskOpinionDisagree {
		userTaskExecution.NodeTaskStatus = constant.InstanceNodeTaskStatusNotPass
		return true
	}
	userTasks := service.GetOpUserTasks(userTaskExecution.Execution.InstTaskID, userTaskExecution.NodeTaskID)
	if utils.IsEmptySlice(userTasks) {
		return false
	}
	//完成人数：依次审批默认0所有人不可选人，会签默认0所有人（可选人大于0），或签默认1一个人（可选人大于0），为0时候比较同意和拒绝的人数判断
	finishModeCount := userTaskExecution.FinishMode
	var finishCount = 0
	for _, userTask := range userTasks {
		if userTask.UserTaskID == userTaskExecution.UserTaskID {
			continue
		}
		if userTask.Status == constant.InstanceUserTaskStatusAgree {
			finishCount++
		}
	}
	if userTaskExecution.Opinion == constant.InstanceUserTaskOpinionAgree {
		finishCount++
	}
	//当前节点完成
	if finishCount == len(userTasks) {
		userTaskExecution.NodeTaskStatus = constant.InstanceNodeTaskStatusComplete
		return true
	}
	if finishModeCount == 0 {
		userTaskExecution.NodeTaskStatus = constant.InstanceNodeTaskStatusComplete
		return true
	}
	if finishCount >= finishModeCount {
		userTaskExecution.NodeTaskStatus = constant.InstanceNodeTaskStatusComplete
		return true
	}
	return false
}

// execApprovalWayOr
// @Description: 或签默认1一个人（可选人大于0），为0所有人时候比较同意和拒绝的人数判断，不为0不是所有人判断同意或者拒绝的人数大于等于完成人数
func execApprovalWayOr(userTaskExecution *UserTaskExecution) bool {
	userTasks := service.GetOpUserTasks(userTaskExecution.Execution.InstTaskID, userTaskExecution.NodeTaskID)
	if utils.IsEmptySlice(userTasks) {
		return false
	}
	//完成人数：依次审批默认0所有人不可选人，会签默认0所有人（可选人大于0），或签默认1一个人（可选人大于0），为0时候比较同意和拒绝的人数判断
	finishModeCount := userTaskExecution.FinishMode
	var agreeCount = 0
	var disagreeCount = 0
	for _, userTask := range userTasks {
		if userTask.UserTaskID == userTaskExecution.UserTaskID {
			continue
		}
		if userTask.Status == constant.InstanceUserTaskStatusAgree {
			agreeCount++
		}
		if userTask.Status == constant.InstanceUserTaskStatusDisagree {
			disagreeCount++
		}
	}
	//为0所有人时候比较同意和拒绝的人数判断，不为0不是所有人判断同意或者拒绝的人数大于等于完成人数
	if userTaskExecution.Opinion == constant.InstanceUserTaskOpinionAgree {
		agreeCount++
		if finishModeCount != 0 && agreeCount > finishModeCount {
			userTaskExecution.NodeTaskStatus = constant.InstanceNodeTaskStatusComplete
			return true
		}
	}
	if userTaskExecution.Opinion == constant.InstanceUserTaskOpinionDisagree {
		disagreeCount++
		if finishModeCount != 0 && disagreeCount > finishModeCount {
			userTaskExecution.NodeTaskStatus = constant.InstanceNodeTaskStatusNotPass
			return true
		}
	}
	if finishModeCount != 0 {
		return false
	}
	count := len(userTasks)
	if count%2 == 1 {
		if agreeCount >= (count/2 + 1) {
			userTaskExecution.NodeTaskStatus = constant.InstanceNodeTaskStatusComplete
			return true
		}
		if disagreeCount >= (count/2 + 1) {
			userTaskExecution.NodeTaskStatus = constant.InstanceNodeTaskStatusNotPass
			return true
		}
	} else {
		if agreeCount >= (count / 2) {
			userTaskExecution.NodeTaskStatus = constant.InstanceNodeTaskStatusComplete
			return true
		}
		if disagreeCount >= (count / 2) {
			userTaskExecution.NodeTaskStatus = constant.InstanceNodeTaskStatusNotPass
			return true
		}
	}
	return false
}
