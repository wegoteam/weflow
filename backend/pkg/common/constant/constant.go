package constant

/*
	节点模型【1：开始节点；2：审批节点；3：办理节点；4：抄送节点；5：自定义节点；6：条件节点；7：分支节点；8：汇聚节点；9：结束节点】
*/
const (
	StartNodeModel       = 1 //开始节点
	ApprovalNodeModel    = 2 //审批节点
	TransactNodeModel    = 3 //办理节点
	NotifyNodeModel      = 4 //抄送节点
	CustomNodeModel      = 5 //自定义节点
	ConditionNodeModel   = 6 //条件节点
	BranchNodeModel      = 7 //分支节点
	ConvergenceNodeModel = 8 //汇聚节点
	EndNodeModel         = 9 //结束节点
)

/**
Redis key的前缀
*/
const (
	RedisProcessDefModel = "weflow:proess-def:" //流程定义Redis的前缀
)

const (
	HasRedisProcessDefModel = 1 //存在流程定义标志
)

/**
实例任务状态【0：草稿；1：创建中；2：进行中； 3：终止； 4：完成； 5：挂起；6：回退】
*/
const (
	InstanceTaskStatusDraft    = 0 //草稿
	InstanceTaskStatusCreate   = 1 //创建中(草稿)
	InstanceTaskStatusDoing    = 2 //进行中
	InstanceTaskStatusStop     = 3 //终止
	InstanceTaskStatusComplete = 4 //完成
	InstanceTaskStatusHangUp   = 5 //挂起
	InstanceTaskStatusRollback = 6 //回退
)

/**
实例节点任务状态【1：未开始；2：处理中；3：完成；4：回退；5：终止；6：不通过】
*/
const (
	InstanceNodeTaskStatusNotStart = 1 //未开始
	InstanceNodeTaskStatusDoing    = 2 //处理中
	InstanceNodeTaskStatusComplete = 3 //完成
	InstanceNodeTaskStatusRollback = 4 //回退
	InstanceNodeTaskStatusStop     = 5 //终止
	InstanceNodeTaskStatusNotPass  = 6 //不通过
)

/**
实例用户任务状态【1：处理中；2：完成（同意）；3：不通过（不同意）；4：回退；5：终止】
*/
const (
	InstanceUserTaskStatusDoing    = 1 //处理中
	InstanceUserTaskStatusAgree    = 2 //完成(同意)
	InstanceUserTaskStatusDisagree = 3 //不通过（不同意）
	InstanceUserTaskStatusRollback = 4 //回退
	InstanceUserTaskStatusStop     = 5 //终止
	InstanceUserTaskStatusSave     = 6 //保存
)

/**
实例用户任务处理意见【1：未发表；2：已阅；3：同意；4：不同意】
*/
const (
	InstanceUserTaskOpinionNotPublish = 1 //未发表
	InstanceUserTaskOpinionAgree      = 2 //同意
	InstanceUserTaskOpinionDisagree   = 3 //不同意
	InstanceUserTaskOpinionRollback   = 4 //回退
	InstanceUserTaskOpinionStop       = 5 //终止
	InstanceUserTaskOpinionSave       = 6 //保存
	InstanceUserTaskOpinionStart      = 7 //发起
	InstanceUserTaskOpinionSuspend    = 8 //挂起
	InstanceUserTaskOpinionSesume     = 9 //恢复
)

/**
常用审批人【指定成员：1；发起人自己：2；发起人自选：3：角色：4；部门：5】主管（相对岗位）【直属主管：1；部门主管：2；连续多级主管：3；部门控件对应主管：4】其他【表单人员控件：1；部门控件：2；角色控件：3】
*/
const (
	//常用审批人
	ApprovalUserTypeUser            = 1 //指定成员
	ApprovalUserTypeInitiator       = 2 //发起人自己
	ApprovalUserTypeInitiatorSelect = 3 //发起人自选
	ApprovalUserTypeRole            = 4 //角色
	ApprovalUserTypeDept            = 5 //部门
	//主管（相对岗位）
	RelativeTypeDeptDirectly = 1 //直属主管
	RelativeTypeDeptDept     = 2 //部门主管
	RelativeTypeDeptMulti    = 3 //连续多级主管
	RelativeTypeDeptControl  = 4 //部门控件对应主管
	//其他
	OtherTypeFormMember = 1 //表单人员控件
	OtherTypeFormDept   = 2 //部门控件
	OtherTypeFormRole   = 3 //角色控件
)

/**
处理人策略【常用审批人：1；主管（相对岗位）：2；其他：3】
*/
const (
	ApprovalUserStrategyCommon       = 1 //常用审批人
	ApprovalUserStrategyRelativePost = 2 //主管（相对岗位）
	ApprovalUserStrategyOther        = 3 //其他
)

/**
审批类型【人工审批：1；自动通过：2；自动拒绝】默认人工审批1
*/
const (
	ApprovalTypeManual = 1 //人工审批
	ApprovalTypePass   = 2 //自动通过
	ApprovalTypeRefuse = 3 //自动拒绝
)

/**
审批人为空时【自动通过：1；自动转交管理员：2；指定审批人：3】默认自动通过1
*/
const (
	ApprovalEmptyTypePass     = 1 //自动通过
	ApprovalEmptyTypeTransfer = 2 //自动转交管理员
	ApprovalEmptyTypeAssign   = 3 //指定审批人
)

/**
审批方式【依次审批：1、会签（需要完成人数的审批人同意或拒绝才可完成节点）：2、或签（其中一名审批人同意或拒绝即可）：3】默认会签2
*/
const (
	ApprovalWayOrder = 1 //依次审批
	ApprovalWayCount = 2 //会签
	ApprovalWayOr    = 3 //或签
)

/**
分支执行方式【单分支：1；多分支：2】默认多分支2
*/
const (
	BranchWaySingle = 1 //单分支
	BranchWayMulti  = 2 //多分支
)

/**
表单权限【可编辑：1；只读：2；隐藏：3；必填：4】默认只读2
*/
const (
	FormPermissionEdit   = 1 //可编辑
	FormPermissionRead   = 2 //只读
	FormPermissionHidden = 3 //隐藏
	FormPermissionMust   = 4 //必填
)

/**
执行操作类型【添加：1；修改：2；删除：3】
*/
const (
	OperationTypeAdd    = 1 //添加
	OperationTypeUpdate = 2 //修改
	OperationTypeDelete = 3 //删除
)

/**
分支节点完成标志【1：分支节点未完成；2：分支节点完成且存在出口；3：分支节点完成无分支出口】
*/
const (
	BranchNodeStatusNotComplete = 1 //分支节点未完成
	BranchNodeStatusComplete    = 2 //分支节点完成且存在出口
	BranchNodeStatusNoBranch    = 3 //分支节点完成无分支出口
)

/**
实例任务操作日志类型【1：节点；2：任务；3：其他】
*/
const (
	InstTaskOpLogNode  = 1
	InstTaskOpLogTask  = 2
	InstTaskOpLogOther = 3
)

/**
模板状态【1：草稿；2：发布；3：停用】默认草稿
*/
const (
	ModelStatusDraft    = 1 //草稿
	ModelStatusDeployed = 2 //已发布
	ModelStatusInvalid  = 3 //已停用
)

/**
使用状态【1：使用；2：未使用】
*/
const (
	ModelVersionUseStatusUse   = 1 //使用
	ModelVersionUseStatusUnUse = 2 //未使用
)
