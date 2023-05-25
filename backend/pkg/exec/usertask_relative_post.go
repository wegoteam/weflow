package exec

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/wegoteam/weflow/pkg/common/constant"
	"github.com/wegoteam/weflow/pkg/common/entity"
	"time"
)

// IRelativeStrategy
// @Description: 主管（相对岗位）【直属主管：1；部门主管：2；连续多级主管：3；部门控件对应主管：4】
type IRelativeStrategy interface {
	genUserTasks() []entity.UserTaskBO
}

func GenRelativePostStrategy(genUserTaskBO *GenUserTaskBO) IExecNodeHandler {

	switch genUserTaskBO.Type {
	case constant.RelativeTypeDeptDirectly:
		return &RelativeTypeDeptDirectly{
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
	case constant.RelativeTypeDeptDept:
		return &RelativeTypeDeptDept{
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
	case constant.RelativeTypeDeptMulti:
		return &RelativeTypeDeptMulti{
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
	case constant.RelativeTypeDeptControl:
		return &RelativeTypeDeptDeptControl{
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
		hlog.Error("实例任务[%s]节点[%s]执行相对岗位策略生成用户任务类型设置有误，请检查配置", genUserTaskBO.InstTaskID, genUserTaskBO.NodeID)
		return nil
	}
}

// RelativeTypeDeptDirectly
// @Description: 相对岗位【直属主管：1；部门主管：2；连续多级主管：3；部门控件对应主管：4】
type RelativeTypeDeptDirectly struct {
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
// @Description: 直属主管生成用户任务
// @receiver relativeTypeDeptDirectly
// @return []entity.UserTaskBO
func (relativeTypeDeptDirectly *RelativeTypeDeptDirectly) genUserTasks() []entity.UserTaskBO {
	userTasks := make([]entity.UserTaskBO, 0)

	return userTasks
}

// RelativeTypeDeptDept
// @Description: 相对岗位【直属主管：1；部门主管：2；连续多级主管：3；部门控件对应主管：4】
type RelativeTypeDeptDept struct {
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
// @Description: 部门主管生成用户任务
// @receiver relativeTypeDeptDept
// @return []entity.UserTaskBO
func (relativeTypeDeptDept *RelativeTypeDeptDept) genUserTasks() []entity.UserTaskBO {
	userTasks := make([]entity.UserTaskBO, 0)

	return userTasks
}

// RelativeTypeDeptMulti
// @Description: 相对岗位【直属主管：1；部门主管：2；连续多级主管：3；部门控件对应主管：4】
type RelativeTypeDeptMulti struct {
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
// @Description: 连续多级主管生成用户任务
// @receiver relativeTypeDeptMulti
// @return []entity.UserTaskBO
func (relativeTypeDeptMulti *RelativeTypeDeptMulti) genUserTasks() []entity.UserTaskBO {
	userTasks := make([]entity.UserTaskBO, 0)

	return userTasks
}

// RelativeTypeDeptDeptControl
// @Description: 相对岗位【直属主管：1；部门主管：2；连续多级主管：3；部门控件对应主管：4】
type RelativeTypeDeptDeptControl struct {
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
// @Description: 部门控件对应主管生成用户任务
// @receiver relativeTypeDeptDeptControl
// @return []entity.UserTaskBO
func (relativeTypeDeptDeptControl *RelativeTypeDeptDeptControl) genUserTasks() []entity.UserTaskBO {
	userTasks := make([]entity.UserTaskBO, 0)

	return userTasks
}
