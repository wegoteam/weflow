package exec

import (
	"github.com/wegoteam/weflow/pkg/common/entity"
	"time"
)

// ICommApproverStrategy 常用审批人
// @Description: 常用审批人【指定成员：1；发起人自己：2；发起人自选：3：角色：4；部门：5】
type ICommApproverStrategy interface {
	GetUserTasks(nodeHandler entity.NodeHandler) []entity.UserTaskBO
}

// ApprovalUserTypeUser
// @Description: 指定成员
type ApprovalUserTypeUser struct {
	InstTaskID     string            //实例任务id
	NodeTaskID     string            //节点任务id
	NodeID         string            //节点id
	CreateUserID   string            //创建人ID
	CreateUserName string            //创建人名称
	Now            time.Time         //当前时间
	Type           int               //常用审批人【指定成员：1；发起人自己：2；发起人自选：3：角色：4；部门：5】主管（相对岗位）【直属主管：1；部门主管：2；连续多级主管：3；部门控件对应主管：4】其他【表单人员控件：1；部门控件：2；角色控件：3】
	Handlers       []entity.Handlers //处理人列表
	Strategy       int               //处理人策略【常用审批人：1；主管（相对岗位）：2；其他：3】
	Obj            string            //扩展字段，设计中可忽略
	Relative       string            //相对岗位，设计中可忽略
}

// ApprovalUserTypeInitiator
// @Description: 发起人自己
type ApprovalUserTypeInitiator struct {
}

// ApprovalUserTypeInitiatorSelect
// @Description: 发起人自选
type ApprovalUserTypeInitiatorSelect struct {
}

// ApprovalUserTypeRole
// @Description: 角色
type ApprovalUserTypeRole struct {
}

// ApprovalUserTypeDept
// @Description: 部门
type ApprovalUserTypeDept struct {
}

func (receiver *ApprovalUserTypeUser) GetUserTasks(nodeHandler entity.NodeHandler) []entity.UserTaskBO {

	return nil
}

func (receiver *ApprovalUserTypeInitiator) GetUserTasks(nodeHandler entity.NodeHandler) []entity.UserTaskBO {

	return nil
}
