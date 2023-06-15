package exec

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/pkg/errors"
	"github.com/wegoteam/weflow/pkg/common/constant"
	"github.com/wegoteam/weflow/pkg/common/entity"
	"github.com/wegoteam/weflow/pkg/common/utils"
	"github.com/wegoteam/weflow/pkg/service"
)

// Agree
// @Description: 同意
// @param userTaskID 用户任务ID
// @param opUserID 操作用户ID
// @param OpUserName 操作用户名称
// @param opinionDesc 意见描述
// @param params 参数
// @return bool
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
	userTaskExecution.Opinion = constant.InstanceUserTaskStatusAgree
	return userTaskExecution.agree(userTaskID, params)
}

// Save
// @Description: 保存
// @param userTaskID 用户任务ID
// @param opUserID 操作用户ID
// @param OpUserName 操作用户名称
// @param opinionDesc 意见描述
// @param params 参数
// @return bool
func Save(userTaskID, opUserID, opUserName, opinionDesc string, params map[string]any) error {
	//创建用户任务执行对象
	userTaskExecution, err := NewUserTaskExecution(userTaskID)
	if err != nil {
		return err
	}
	//验证用户任务信息
	userTaskExecution.verify(opUserID)
	userTaskExecution.OpUserID = opUserID
	userTaskExecution.OpUserName = opUserName
	userTaskExecution.OpinionDesc = opinionDesc
	userTaskExecution.Opinion = constant.InstanceUserTaskOpinionSave
	return userTaskExecution.save(userTaskID, params)
}

// Disagree
// @Description: 不同意
// @param userTaskID 用户任务ID
// @param opUserID 操作用户ID
// @param OpUserName 操作用户名称
// @param opinionDesc 意见描述
// @return bool
func Disagree(userTaskID, opUserID, opUserName, opinionDesc string) error {
	//创建用户任务执行对象
	userTaskExecution, err := NewUserTaskExecution(userTaskID)
	if err != nil {
		return err
	}
	//验证用户任务信息
	userTaskExecution.verify(opUserID)
	userTaskExecution.OpUserID = opUserID
	userTaskExecution.OpUserName = opUserName
	userTaskExecution.OpinionDesc = opinionDesc
	userTaskExecution.UserTaskStatus = constant.InstanceUserTaskStatusDisagree
	userTaskExecution.Opinion = constant.InstanceUserTaskStatusDisagree
	return userTaskExecution.disagree(userTaskID)
}

// agree
// @Description: 同意
//保存实例任务参数、修改当前用户任务、修改当前节点任务、添加用户任务评论、保存数据
// @receiver userTaskExecution
// @param userTaskID 用户任务ID
// @param params 参数
// @return bool
func (userTaskExecution *UserTaskExecution) agree(userTaskID string, params map[string]any) error {
	execution := userTaskExecution.Execution
	execution.InstTaskParamMap = params
	//执行流转
	execNodeTask(userTaskExecution)
	//修改当前用户任务
	userTasks := execution.UserTasks
	editUserTask := &entity.UserTaskBO{
		ExecOpType: constant.OperationTypeUpdate,
		UserTaskID: userTaskID,
		Status:     constant.InstanceUserTaskStatusAgree,
		//Opinion:     int32(userTaskExecution.Opinion),
		OpinionDesc: userTaskExecution.OpinionDesc,
		UpdateTime:  execution.Now,
	}
	*userTasks = append(*userTasks, *editUserTask)
	//执行数据
	err := userTaskExecution.execInstUserTaskData()
	if err != nil {
		return err
	}
	hlog.Infof("当前节点任务[%s]同意操作，节点任务已完成", userTaskID)
	return nil
}

// save
// @Description: 保存
// @receiver userTaskExecution
// @param userTaskID 用户任务ID
// @param params 参数
// @return bool
func (userTaskExecution *UserTaskExecution) save(userTaskID string, params map[string]any) error {
	execution := userTaskExecution.Execution
	execution.InstTaskParamMap = params

	//执行数据
	err := userTaskExecution.execInstUserTaskData()
	if err != nil {
		return err
	}
	return nil
}

// disagree
// @Description: 不同意
// @receiver userTaskExecution
// @param userTaskID
// @param params
// @return bool
func (userTaskExecution *UserTaskExecution) disagree(userTaskID string) error {
	execution := userTaskExecution.Execution
	//execution.InstTaskParamMap = params
	//验证表单权限
	execNodeTask(userTaskExecution)
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
	hlog.Infof("当前节点任务[%s]不同意操作，节点任务已完成", userTaskID)
	return nil
}

// turn
// @Description: 转办
// @receiver userTaskExecution
// @return bool
func (userTaskExecution *UserTaskExecution) turn() error {

	return nil
}

// delegate
// @Description: 委托
// @receiver userTaskExecution
// @return bool
func (userTaskExecution *UserTaskExecution) delegate() error {

	return nil
}

// rollback
// @Description: 回退上节点
// @receiver userTaskExecution
// @return bool
func (userTaskExecution *UserTaskExecution) rollback() error {

	return nil
}

// rollbackStartNode
// @Description: 回退发起节点
// @receiver userTaskExecution
// @return bool
func (userTaskExecution *UserTaskExecution) rollbackStartNode() error {

	return nil
}

// rollbackAnyNode
// @Description: 回退任意节点
// @receiver userTaskExecution
// @return bool
func (userTaskExecution *UserTaskExecution) rollbackAnyNode() error {

	return nil
}

// revoke
// @Description: 撤回
// @receiver userTaskExecution
// @return bool
func (userTaskExecution *UserTaskExecution) revoke() error {

	return nil
}

// cancel
// @Description: 取消
// @receiver userTaskExecution
// @return bool
func (userTaskExecution *UserTaskExecution) cancel() error {

	return nil
}

// urge
// @Description: 催办
// @receiver userTaskExecution
// @return bool
func (userTaskExecution *UserTaskExecution) urge() error {

	return nil
}

// addSign
// @Description: 加签
// @receiver userTaskExecution
// @return bool
func (userTaskExecution *UserTaskExecution) addSign() error {

	return nil
}

// reduceSign
// @Description: 减签
// @receiver userTaskExecution
// @return bool
func (userTaskExecution *UserTaskExecution) reduceSign() error {

	return nil
}

// cc
// @Description: 抄送
// @receiver userTaskExecution
// @return bool
func (userTaskExecution *UserTaskExecution) cc() error {

	return nil
}

// ccReply
// @Description: 抄送回复
// @receiver userTaskExecution
// @return bool
func (userTaskExecution *UserTaskExecution) ccReply() error {

	return nil
}

// ccRevoke
// @Description: 抄送撤回
// @receiver userTaskExecution
// @return bool
func (userTaskExecution *UserTaskExecution) ccRevoke() error {

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
// @param userTaskExecution
func execNodeTask(userTaskExecution *UserTaskExecution) {
	//判断当前的节点任务是否完成
	finishFlag := isFinish(userTaskExecution)
	if !finishFlag {
		//执行数据
		userTaskExecution.execInstUserTaskData()
		hlog.Infof("当前节点任务[%s]同意操作，节点任务未完成", userTaskExecution.UserTaskID)
		return
	}
	execution := userTaskExecution.Execution
	//验证节点任务：依次审批、会签、或签；取决于流转用户任务还是流转节点任务
	//执行流转节点获取是流转任务
	processDefModel := execution.ProcessDefModel
	currNodeModelBO := processDefModel.NodeModelMap[userTaskExecution.NodeID]
	//依次审批
	if userTaskExecution.HandleMode == constant.ApprovalWayOrder {
		//判断当前节点任务依次审批的用户任务最大处理顺序
		maxOpSort := service.GetUserTaskMaxOpSort(&currNodeModelBO)
		if userTaskExecution.OpSort != maxOpSort {
			//执行下任务
			execNextTask(&currNodeModelBO, userTaskExecution)
		}
	}
	//修改当前节点任务
	nodeTasks := execution.InstNodeTasks
	editNodeTask := &entity.InstNodeTaskBO{
		ExecOpType: constant.OperationTypeUpdate,
		NodeTaskID: userTaskExecution.NodeTaskID,
		Status:     int32(userTaskExecution.NodeTaskStatus),
		UpdateTime: execution.Now,
	}
	*nodeTasks = append(*nodeTasks, *editNodeTask)
	//抄送任务不继续流转
	if userTaskExecution.NodeModel == constant.NotifyNodeModel {
		return
	}
	//执行下一个节点
	execNextNode(&currNodeModelBO, execution)
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
	if userTaskExecution.UserTaskStatus == constant.InstanceUserTaskStatusAgree {
		finishCount++
	}
	//当前节点完成
	if finishCount == len(userTasks) {
		return true
	}
	return false
}

// execApprovalWayCount
// @Description: 会签默认0所有人（可选人大于0），需要所有人同意
func execApprovalWayCount(userTaskExecution *UserTaskExecution) bool {
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
	if userTaskExecution.UserTaskStatus == constant.InstanceUserTaskStatusAgree {
		finishCount++
	}
	//当前节点完成
	if finishCount == len(userTasks) {
		return true
	}
	if finishModeCount == 0 {
		return true
	}
	return finishCount >= finishModeCount
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
	if userTaskExecution.UserTaskStatus == constant.InstanceUserTaskStatusAgree {
		agreeCount++
		if finishModeCount != 0 && agreeCount > finishModeCount {
			userTaskExecution.NodeTaskStatus = constant.InstanceNodeTaskStatusComplete
			return true
		}
	}
	if userTaskExecution.UserTaskStatus == constant.InstanceUserTaskStatusDisagree {
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
