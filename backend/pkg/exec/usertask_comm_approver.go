package exec

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/wegoteam/weflow/pkg/common/constant"
	"github.com/wegoteam/weflow/pkg/common/entity"
	"github.com/wegoteam/wepkg/snowflake"
	"time"
)

// ICommApproverStrategy 常用审批人
// @Description: 常用审批人【指定成员：1；发起人自己：2；发起人自选：3：角色：4；部门：5】
type ICommApproverStrategy interface {
	GenUserTasks() []entity.UserTaskBO
}

func GenCommApproverStrategy(genUserTaskBO *GenUserTaskBO) IExecNodeHandler {

	switch genUserTaskBO.Type {
	case constant.ApprovalUserTypeUser:
		return &ApprovalUserTypeUser{
			InstTaskID:     genUserTaskBO.InstTaskID,
			NodeTaskID:     genUserTaskBO.NodeTaskID,
			NodeID:         genUserTaskBO.NodeID,
			CreateUserID:   genUserTaskBO.CreateUserID,
			CreateUserName: genUserTaskBO.CreateUserName,
			Now:            genUserTaskBO.Now,
			Type:           genUserTaskBO.Type,
			Handlers:       genUserTaskBO.Handler,
			Strategy:       genUserTaskBO.Strategy,
			Obj:            genUserTaskBO.Obj,
		}
	case constant.ApprovalUserTypeInitiator:
		return &ApprovalUserTypeInitiator{
			InstTaskID:     genUserTaskBO.InstTaskID,
			NodeTaskID:     genUserTaskBO.NodeTaskID,
			NodeID:         genUserTaskBO.NodeID,
			CreateUserID:   genUserTaskBO.CreateUserID,
			CreateUserName: genUserTaskBO.CreateUserName,
			Now:            genUserTaskBO.Now,
			Type:           genUserTaskBO.Type,
			Handlers:       genUserTaskBO.Handler,
			Strategy:       genUserTaskBO.Strategy,
			Obj:            genUserTaskBO.Obj,
		}
	case constant.ApprovalUserTypeInitiatorSelect:
		return &ApprovalUserTypeInitiatorSelect{
			InstTaskID:     genUserTaskBO.InstTaskID,
			NodeTaskID:     genUserTaskBO.NodeTaskID,
			NodeID:         genUserTaskBO.NodeID,
			CreateUserID:   genUserTaskBO.CreateUserID,
			CreateUserName: genUserTaskBO.CreateUserName,
			Now:            genUserTaskBO.Now,
			Type:           genUserTaskBO.Type,
			Handlers:       genUserTaskBO.Handler,
			Strategy:       genUserTaskBO.Strategy,
			Obj:            genUserTaskBO.Obj,
		}
	case constant.ApprovalUserTypeRole:
		return &ApprovalUserTypeRole{
			InstTaskID:     genUserTaskBO.InstTaskID,
			NodeTaskID:     genUserTaskBO.NodeTaskID,
			NodeID:         genUserTaskBO.NodeID,
			CreateUserID:   genUserTaskBO.CreateUserID,
			CreateUserName: genUserTaskBO.CreateUserName,
			Now:            genUserTaskBO.Now,
			Type:           genUserTaskBO.Type,
			Handlers:       genUserTaskBO.Handler,
			Strategy:       genUserTaskBO.Strategy,
			Obj:            genUserTaskBO.Obj,
		}
	case constant.ApprovalUserTypeDept:
		return &ApprovalUserTypeDept{
			InstTaskID:     genUserTaskBO.InstTaskID,
			NodeTaskID:     genUserTaskBO.NodeTaskID,
			NodeID:         genUserTaskBO.NodeID,
			CreateUserID:   genUserTaskBO.CreateUserID,
			CreateUserName: genUserTaskBO.CreateUserName,
			Now:            genUserTaskBO.Now,
			Type:           genUserTaskBO.Type,
			Handlers:       genUserTaskBO.Handler,
			Strategy:       genUserTaskBO.Strategy,
			Obj:            genUserTaskBO.Obj,
		}
	default:
		hlog.Error("实例任务[%s]节点[%s]执行常用审批人策略生成用户任务类型设置有误，请检查配置", genUserTaskBO.InstTaskID, genUserTaskBO.NodeID)
		return nil
	}
}

// ApprovalUserTypeUser
// @Description: 指定成员
type ApprovalUserTypeUser struct {
	InstTaskID     string          //实例任务id
	NodeTaskID     string          //节点任务id
	NodeID         string          //节点id
	CreateUserID   string          //创建人ID
	CreateUserName string          //创建人名称
	Now            time.Time       //当前时间
	Type           int             //常用审批人【指定成员：1；发起人自己：2；发起人自选：3：角色：4；部门：5】主管（相对岗位）【直属主管：1；部门主管：2；连续多级主管：3；部门控件对应主管：4】其他【表单人员控件：1；部门控件：2；角色控件：3】
	Handlers       entity.Handlers //处理人列表
	Strategy       int             //处理人策略【常用审批人：1；主管（相对岗位）：2；其他：3】
	Obj            string          //扩展字段，设计中可忽略
	Relative       string          //相对岗位，设计中可忽略
}

func (approvalUserTypeUser *ApprovalUserTypeUser) GenUserTasks() []entity.UserTaskBO {
	userTasks := make([]entity.UserTaskBO, 0)

	handler := approvalUserTypeUser.Handlers
	var userTask = entity.UserTaskBO{
		InstTaskID:   approvalUserTypeUser.InstTaskID,
		NodeTaskID:   approvalUserTypeUser.NodeTaskID,
		NodeID:       approvalUserTypeUser.NodeID,
		UserTaskID:   snowflake.GetSnowflakeId(),
		Type:         int32(approvalUserTypeUser.Type),
		Strategy:     int32(approvalUserTypeUser.Strategy),
		NodeUserName: handler.Name,
		NodeUserID:   handler.ID,
		Sort:         int32(handler.Sort),
		Obj:          approvalUserTypeUser.Obj,
		Relative:     approvalUserTypeUser.Relative,
		Status:       constant.InstanceUserTaskStatusDoing,
		CreateTime:   approvalUserTypeUser.Now,
		UpdateTime:   approvalUserTypeUser.Now,
		HandleTime:   approvalUserTypeUser.Now,
		OpUserID:     handler.ID,
		OpUserName:   handler.Name,
		Opinion:      constant.InstanceUserTaskOpinionNotPublish,
		OpinionDesc:  "",
	}
	userTasks = append(userTasks, userTask)

	return userTasks
}

// ApprovalUserTypeInitiator
// @Description: 发起人自己
type ApprovalUserTypeInitiator struct {
	InstTaskID     string          //实例任务id
	NodeTaskID     string          //节点任务id
	NodeID         string          //节点id
	CreateUserID   string          //创建人ID
	CreateUserName string          //创建人名称
	Now            time.Time       //当前时间
	Type           int             //常用审批人【指定成员：1；发起人自己：2；发起人自选：3：角色：4；部门：5】主管（相对岗位）【直属主管：1；部门主管：2；连续多级主管：3；部门控件对应主管：4】其他【表单人员控件：1；部门控件：2；角色控件：3】
	Handlers       entity.Handlers //处理人列表
	Strategy       int             //处理人策略【常用审批人：1；主管（相对岗位）：2；其他：3】
	Obj            string          //扩展字段，设计中可忽略
	Relative       string          //相对岗位，设计中可忽略
}

func (a *ApprovalUserTypeInitiator) GenUserTasks() []entity.UserTaskBO {
	//TODO implement me
	panic("implement me")
}

// ApprovalUserTypeInitiatorSelect
// @Description: 发起人自选
type ApprovalUserTypeInitiatorSelect struct {
	InstTaskID     string          //实例任务id
	NodeTaskID     string          //节点任务id
	NodeID         string          //节点id
	CreateUserID   string          //创建人ID
	CreateUserName string          //创建人名称
	Now            time.Time       //当前时间
	Type           int             //常用审批人【指定成员：1；发起人自己：2；发起人自选：3：角色：4；部门：5】主管（相对岗位）【直属主管：1；部门主管：2；连续多级主管：3；部门控件对应主管：4】其他【表单人员控件：1；部门控件：2；角色控件：3】
	Handlers       entity.Handlers //处理人列表
	Strategy       int             //处理人策略【常用审批人：1；主管（相对岗位）：2；其他：3】
	Obj            string          //扩展字段，设计中可忽略
	Relative       string          //相对岗位，设计中可忽略
}

func (a ApprovalUserTypeInitiatorSelect) GenUserTasks() []entity.UserTaskBO {
	//TODO implement me
	panic("implement me")
}

// ApprovalUserTypeRole
// @Description: 角色
type ApprovalUserTypeRole struct {
	InstTaskID     string          //实例任务id
	NodeTaskID     string          //节点任务id
	NodeID         string          //节点id
	CreateUserID   string          //创建人ID
	CreateUserName string          //创建人名称
	Now            time.Time       //当前时间
	Type           int             //常用审批人【指定成员：1；发起人自己：2；发起人自选：3：角色：4；部门：5】主管（相对岗位）【直属主管：1；部门主管：2；连续多级主管：3；部门控件对应主管：4】其他【表单人员控件：1；部门控件：2；角色控件：3】
	Handlers       entity.Handlers //处理人列表
	Strategy       int             //处理人策略【常用审批人：1；主管（相对岗位）：2；其他：3】
	Obj            string          //扩展字段，设计中可忽略
	Relative       string          //相对岗位，设计中可忽略
}

func (a ApprovalUserTypeRole) GenUserTasks() []entity.UserTaskBO {
	//TODO implement me
	panic("implement me")
}

// ApprovalUserTypeDept
// @Description: 部门
type ApprovalUserTypeDept struct {
	InstTaskID     string          //实例任务id
	NodeTaskID     string          //节点任务id
	NodeID         string          //节点id
	CreateUserID   string          //创建人ID
	CreateUserName string          //创建人名称
	Now            time.Time       //当前时间
	Type           int             //常用审批人【指定成员：1；发起人自己：2；发起人自选：3：角色：4；部门：5】主管（相对岗位）【直属主管：1；部门主管：2；连续多级主管：3；部门控件对应主管：4】其他【表单人员控件：1；部门控件：2；角色控件：3】
	Handlers       entity.Handlers //处理人列表
	Strategy       int             //处理人策略【常用审批人：1；主管（相对岗位）：2；其他：3】
	Obj            string          //扩展字段，设计中可忽略
	Relative       string          //相对岗位，设计中可忽略
}

func (a ApprovalUserTypeDept) GenUserTasks() []entity.UserTaskBO {
	//TODO implement me
	panic("implement me")
}
