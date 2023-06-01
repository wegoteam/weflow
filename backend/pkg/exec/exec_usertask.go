package exec

// agree
// @Description: 同意
// @receiver userTaskExecution
// @param userTaskID
// @param params
// @return bool
func (userTaskExecution *UserTaskExecution) agree(userTaskID string, params map[string]any) bool {

	return true
}

// disagree
// @Description: 不同意
// @receiver userTaskExecution
// @param userTaskID
// @param params
// @return bool
func (userTaskExecution *UserTaskExecution) disagree(userTaskID string, params map[string]any) bool {

	return true
}

// turn
// @Description: 转办
// @receiver userTaskExecution
// @return bool
func (userTaskExecution *UserTaskExecution) turn() bool {

	return true
}

// delegate
// @Description: 委托
// @receiver userTaskExecution
// @return bool
func (userTaskExecution *UserTaskExecution) delegate() bool {

	return true
}

// rollback
// @Description: 回退上节点
// @receiver userTaskExecution
// @return bool
func (userTaskExecution *UserTaskExecution) rollback() bool {

	return true
}

// rollbackStartNode
// @Description: 回退发起节点
// @receiver userTaskExecution
// @return bool
func (userTaskExecution *UserTaskExecution) rollbackStartNode() bool {

	return true
}

// rollbackAnyNode
// @Description: 回退任意节点
// @receiver userTaskExecution
// @return bool
func (userTaskExecution *UserTaskExecution) rollbackAnyNode() bool {

	return true
}

// revoke
// @Description: 撤回
// @receiver userTaskExecution
// @return bool
func (userTaskExecution *UserTaskExecution) revoke() bool {

	return true
}

// cancel
// @Description: 取消
// @receiver userTaskExecution
// @return bool
func (userTaskExecution *UserTaskExecution) cancel() bool {

	return true
}

// urge
// @Description: 催办
// @receiver userTaskExecution
// @return bool
func (userTaskExecution *UserTaskExecution) urge() bool {

	return true
}

// save
// @Description: 保存
// @receiver userTaskExecution
// @return bool
func (userTaskExecution *UserTaskExecution) save() bool {

	return true
}

// addSign
// @Description: 加签
// @receiver userTaskExecution
// @return bool
func (userTaskExecution *UserTaskExecution) addSign() bool {

	return true
}

// reduceSign
// @Description: 减签
// @receiver userTaskExecution
// @return bool
func (userTaskExecution *UserTaskExecution) reduceSign() bool {

	return true
}

// cc
// @Description: 抄送
// @receiver userTaskExecution
// @return bool
func (userTaskExecution *UserTaskExecution) cc() bool {

	return true
}

// ccReply
// @Description: 抄送回复
// @receiver userTaskExecution
// @return bool
func (userTaskExecution *UserTaskExecution) ccReply() bool {

	return true
}

// ccRevoke
// @Description: 抄送撤回
// @receiver userTaskExecution
// @return bool
func (userTaskExecution *UserTaskExecution) ccRevoke() bool {

	return true
}
