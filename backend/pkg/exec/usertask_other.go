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
	GenUserTasks() []entity.UserTaskBO
}

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

type OtherTypeFormMember struct {
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

func (o OtherTypeFormMember) GenUserTasks() []entity.UserTaskBO {
	//TODO implement me
	panic("implement me")
}

type OtherTypeFormDept struct {
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

func (o OtherTypeFormDept) GenUserTasks() []entity.UserTaskBO {
	//TODO implement me
	panic("implement me")
}

type OtherTypeFormRole struct {
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

func (o OtherTypeFormRole) GenUserTasks() []entity.UserTaskBO {
	//TODO implement me
	panic("implement me")
}
