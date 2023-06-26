package exec

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/wegoteam/weflow/pkg/common/constant"
	"github.com/wegoteam/weflow/pkg/common/entity"
	"time"
)

// IOtherStrategy
// @Description: 其他【表单人员控件：1；部门控件：2；角色控件：3】
type IOtherStrategy interface {
	genUserTasks() []entity.UserTaskBO
}

// GenOtherStrategy
// @Description: 生成其他策略
// @param: genUserTaskBO
// @return IExecNodeHandler
func GenOtherStrategy(genUserTaskBO *GenUserTaskBO) IExecNodeHandler {

	switch genUserTaskBO.Type {
	case constant.OtherTypeFormMember:
		return &OtherTypeFormMember{
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
	case constant.OtherTypeFormDept:
		return &OtherTypeFormDept{
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
	case constant.OtherTypeFormRole:
		return &OtherTypeFormRole{
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
		hlog.Error("实例任务[%s]节点[%s]执行其他策略生成用户任务类型设置有误，请检查配置", genUserTaskBO.InstTaskID, genUserTaskBO.NodeID)
		return nil
	}
}

// OtherTypeFormMember
// @Description: 表单人员控件
type OtherTypeFormMember struct {
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
// @Description: 表单人员控件生成用户任务
// @receiver otherTypeFormMember
// @return []entity.UserTaskBO
func (otherTypeFormMember *OtherTypeFormMember) genUserTasks() []entity.UserTaskBO {
	userTasks := make([]entity.UserTaskBO, 0)

	return userTasks
}

// OtherTypeFormDept
// @Description: 部门控件
type OtherTypeFormDept struct {
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
// @Description: 部门控件生成用户任务
// @receiver otherTypeFormDept
// @return []entity.UserTaskBO
func (otherTypeFormDept *OtherTypeFormDept) genUserTasks() []entity.UserTaskBO {
	userTasks := make([]entity.UserTaskBO, 0)

	return userTasks
}

// OtherTypeFormRole
// @Description: 角色控件
type OtherTypeFormRole struct {
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
// @Description: 角色控件生成用户任务
// @receiver otherTypeFormRole
// @return []entity.UserTaskBO
func (otherTypeFormRole *OtherTypeFormRole) genUserTasks() []entity.UserTaskBO {
	userTasks := make([]entity.UserTaskBO, 0)

	return userTasks
}
