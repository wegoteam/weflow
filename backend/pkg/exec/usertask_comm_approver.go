package exec

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/elliotchance/pie/v2"
	"github.com/wegoteam/weflow/pkg/common/constant"
	"github.com/wegoteam/weflow/pkg/common/entity"
	"github.com/wegoteam/weflow/pkg/service"
	"github.com/wegoteam/wepkg/snowflake"
	"time"
)

// ICommApproverStrategy 常用审批人
// @Description: 常用审批人【指定成员：1；发起人自己：2；发起人自选：3：角色：4；部门：5】
type ICommApproverStrategy interface {
	genUserTasks() []entity.UserTaskBO
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

// GenUserTasks
// @Description: 指定成员生成用户任务
// @receiver approvalUserTypeUser
// @return []entity.UserTaskBO
func (approvalUserTypeUser *ApprovalUserTypeUser) genUserTasks() []entity.UserTaskBO {
	userTasks := make([]entity.UserTaskBO, 0)

	handlers := approvalUserTypeUser.Handlers
	if handlers == nil || len(handlers) == 0 {
		return userTasks
	}
	for _, handler := range handlers {
		var userTask = entity.UserTaskBO{
			ExecOpType:   constant.OperationTypeAdd,
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
	}

	return userTasks
}

// ApprovalUserTypeInitiator
// @Description: 发起人自己
type ApprovalUserTypeInitiator struct {
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

// GenUserTasks
// @Description: 发起人自己生成用户任务
// @receiver approvalUserTypeInitiator
// @return []entity.UserTaskBO
func (approvalUserTypeInitiator *ApprovalUserTypeInitiator) genUserTasks() []entity.UserTaskBO {
	userTasks := make([]entity.UserTaskBO, 0)

	var userTask = entity.UserTaskBO{
		ExecOpType:   constant.OperationTypeAdd,
		InstTaskID:   approvalUserTypeInitiator.InstTaskID,
		NodeTaskID:   approvalUserTypeInitiator.NodeTaskID,
		NodeID:       approvalUserTypeInitiator.NodeID,
		UserTaskID:   snowflake.GetSnowflakeId(),
		Type:         int32(approvalUserTypeInitiator.Type),
		Strategy:     int32(approvalUserTypeInitiator.Strategy),
		NodeUserName: approvalUserTypeInitiator.CreateUserName,
		NodeUserID:   approvalUserTypeInitiator.CreateUserID,
		Sort:         int32(1),
		Obj:          approvalUserTypeInitiator.Obj,
		Relative:     approvalUserTypeInitiator.Relative,
		Status:       constant.InstanceUserTaskStatusDoing,
		CreateTime:   approvalUserTypeInitiator.Now,
		UpdateTime:   approvalUserTypeInitiator.Now,
		HandleTime:   approvalUserTypeInitiator.Now,
		OpUserID:     approvalUserTypeInitiator.CreateUserID,
		OpUserName:   approvalUserTypeInitiator.CreateUserName,
		Opinion:      constant.InstanceUserTaskOpinionNotPublish,
		OpinionDesc:  "",
	}
	userTasks = append(userTasks, userTask)
	return userTasks
}

// ApprovalUserTypeInitiatorSelect
// @Description: 发起人自选
type ApprovalUserTypeInitiatorSelect struct {
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

// GenUserTasks
// @Description: 发起人自选生成用户任务
// @receiver approvalUserTypeInitiatorSelect
// @return []entity.UserTaskBO
func (approvalUserTypeInitiatorSelect *ApprovalUserTypeInitiatorSelect) genUserTasks() []entity.UserTaskBO {
	userTasks := make([]entity.UserTaskBO, 0)

	return userTasks
}

// ApprovalUserTypeRole
// @Description: 角色
type ApprovalUserTypeRole struct {
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

// GenUserTasks
// @Description: 角色生成用户任务
// @receiver a
// @return []entity.UserTaskBO
func (approvalUserTypeRole *ApprovalUserTypeRole) genUserTasks() []entity.UserTaskBO {
	userTasks := make([]entity.UserTaskBO, 0)
	handlers := approvalUserTypeRole.Handlers
	if handlers == nil || len(handlers) == 0 {
		return userTasks
	}
	roleIds := make([]string, 0)
	pie.Each(handlers, func(handler entity.Handlers) {
		roleIds = append(roleIds, handler.ID)
	})
	//获取角色下的用户
	userInfos := service.GetRoleUserInfo(roleIds)
	if userInfos == nil || len(userInfos) == 0 {
		return userTasks
	}
	var sort = 1
	pie.Each(userInfos, func(userInfo entity.UserInfoResult) {
		var userTask = entity.UserTaskBO{
			ExecOpType:   constant.OperationTypeAdd,
			InstTaskID:   approvalUserTypeRole.InstTaskID,
			NodeTaskID:   approvalUserTypeRole.NodeTaskID,
			NodeID:       approvalUserTypeRole.NodeID,
			UserTaskID:   snowflake.GetSnowflakeId(),
			Type:         int32(approvalUserTypeRole.Type),
			Strategy:     int32(approvalUserTypeRole.Strategy),
			NodeUserName: userInfo.UserName,
			NodeUserID:   userInfo.UserID,
			Sort:         int32(sort),
			Obj:          approvalUserTypeRole.Obj,
			Relative:     approvalUserTypeRole.Relative,
			Status:       constant.InstanceUserTaskStatusDoing,
			CreateTime:   approvalUserTypeRole.Now,
			UpdateTime:   approvalUserTypeRole.Now,
			HandleTime:   approvalUserTypeRole.Now,
			OpUserID:     userInfo.UserID,
			OpUserName:   userInfo.UserName,
			Opinion:      constant.InstanceUserTaskOpinionNotPublish,
			OpinionDesc:  "",
		}
		userTasks = append(userTasks, userTask)
		sort = sort + 1
	})
	return userTasks
}

// ApprovalUserTypeDept
// @Description: 部门
type ApprovalUserTypeDept struct {
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

// GenUserTasks
// @Description: 部门生成用户任务
// @receiver approvalUserTypeDept
// @return []entity.UserTaskBO
func (approvalUserTypeDept *ApprovalUserTypeDept) genUserTasks() []entity.UserTaskBO {
	userTasks := make([]entity.UserTaskBO, 0)
	handlers := approvalUserTypeDept.Handlers
	if handlers == nil || len(handlers) == 0 {
		return userTasks
	}
	orgIds := make([]string, 0)
	pie.Each(handlers, func(handler entity.Handlers) {
		orgIds = append(orgIds, handler.ID)
	})
	//获取组织下的用户
	userInfos := service.GetOrgUserInfo(orgIds)
	if userInfos == nil || len(userInfos) == 0 {
		return userTasks
	}
	var sort = 1
	pie.Each(userInfos, func(userInfo entity.UserInfoResult) {
		var userTask = entity.UserTaskBO{
			ExecOpType:   constant.OperationTypeAdd,
			InstTaskID:   approvalUserTypeDept.InstTaskID,
			NodeTaskID:   approvalUserTypeDept.NodeTaskID,
			NodeID:       approvalUserTypeDept.NodeID,
			UserTaskID:   snowflake.GetSnowflakeId(),
			Type:         int32(approvalUserTypeDept.Type),
			Strategy:     int32(approvalUserTypeDept.Strategy),
			NodeUserName: userInfo.UserName,
			NodeUserID:   userInfo.UserID,
			Sort:         int32(sort),
			Obj:          approvalUserTypeDept.Obj,
			Relative:     approvalUserTypeDept.Relative,
			Status:       constant.InstanceUserTaskStatusDoing,
			CreateTime:   approvalUserTypeDept.Now,
			UpdateTime:   approvalUserTypeDept.Now,
			HandleTime:   approvalUserTypeDept.Now,
			OpUserID:     userInfo.UserID,
			OpUserName:   userInfo.UserName,
			Opinion:      constant.InstanceUserTaskOpinionNotPublish,
			OpinionDesc:  "",
		}
		userTasks = append(userTasks, userTask)
		sort = sort + 1
	})
	return userTasks
}
