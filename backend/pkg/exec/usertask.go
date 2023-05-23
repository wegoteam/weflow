package exec

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/wegoteam/weflow/pkg/common/constant"
	"github.com/wegoteam/weflow/pkg/common/entity"
	"github.com/wegoteam/wepkg/snowflake"
	"time"
)

// GetUserTask
//  @Description: 生成用户任务
//处理人策略【常用审批人：1；主管（相对岗位）：2；其他：3】
//常用审批人【指定成员：1；发起人自己：2；发起人自选：3：角色：4；部门：5】主管（相对岗位）【直属主管：1；部门主管：2；连续多级主管：3；部门控件对应主管：4】其他【表单人员控件：1；部门控件：2；角色控件：3】
//  @param instNodeTask
//  @param nodeHandler
//  @return []entity.UserTaskBO
func GetUserTask(instNodeTask entity.InstNodeTaskBO, nodeHandler entity.NodeHandler) []entity.UserTaskBO {
	userTasks := make([]entity.UserTaskBO, 0)
	//生成用户任务
	handlers := nodeHandler.Handlers
	if handlers == nil || len(handlers) == 0 {
		return userTasks
	}
	for _, handler := range handlers {
		var userTask = entity.UserTaskBO{
			InstTaskID:   instNodeTask.InstTaskID,
			NodeTaskID:   instNodeTask.NodeTaskID,
			NodeID:       instNodeTask.NodeID,
			UserTaskID:   snowflake.GetSnowflakeId(),
			Type:         int32(nodeHandler.Type),
			Strategy:     int32(nodeHandler.Strategy),
			NodeUserName: handler.Name,
			NodeUserID:   handler.ID,
			Sort:         int32(handler.Sort),
			Obj:          nodeHandler.Obj,
			Relative:     nodeHandler.Relative,
			Status:       constant.InstanceUserTaskStatusDoing,
			CreateTime:   instNodeTask.CreateTime,
			UpdateTime:   instNodeTask.UpdateTime,
			HandleTime:   instNodeTask.CreateTime,
			OpUserID:     handler.ID,
			OpUserName:   handler.Name,
			Opinion:      constant.InstanceUserTaskOpinionNotPublish,
			OpinionDesc:  "",
		}
		userTasks = append(userTasks, userTask)
	}
	return userTasks
}

// ExecUserTask
// @Description: 生成用户任务
//处理人策略【常用审批人：1；主管（相对岗位）：2；其他：3】
//常用审批人【指定成员：1；发起人自己：2；发起人自选：3：角色：4；部门：5】主管（相对岗位）【直属主管：1；部门主管：2；连续多级主管：3；部门控件对应主管：4】其他【表单人员控件：1；部门控件：2；角色控件：3】
//  @param instNodeTask
//  @param nodeHandler
//  @return []entity.UserTaskBO
func ExecUserTask(execution entity.Execution, instNodeTask entity.InstNodeTaskBO, nodeHandler entity.NodeHandler) []entity.UserTaskBO {
	userTasks := make([]entity.UserTaskBO, 0)
	//生成用户任务
	var genUserTaskBO = GenUserTaskBO{
		InstTaskID:     execution.InstTaskID,
		NodeTaskID:     instNodeTask.NodeTaskID,
		NodeID:         instNodeTask.NodeID,
		CreateUserID:   execution.CreateUserID,
		CreateUserName: execution.CreateUserName,
		Type:           nodeHandler.Type,
		Strategy:       nodeHandler.Strategy,
		Obj:            nodeHandler.Obj,
		Relative:       nodeHandler.Relative,
		Handler:        nodeHandler.Handlers,
		Now:            execution.Now,
	}
	genUserTasks := genUserTaskBO.ExecHandlerStrategy()
	userTasks = append(userTasks, genUserTasks...)
	return userTasks
}

type IExecNodeHandler interface {
	GenUserTasks() []entity.UserTaskBO
}

type GenUserTaskBO struct {
	InstTaskID     string            //实例任务id
	NodeTaskID     string            //节点任务id
	NodeID         string            //节点id
	CreateUserID   string            //创建人ID
	CreateUserName string            //创建人名称
	Now            time.Time         //当前时间
	Type           int               //常用审批人【指定成员：1；发起人自己：2；发起人自选：3：角色：4；部门：5】主管（相对岗位）【直属主管：1；部门主管：2；连续多级主管：3；部门控件对应主管：4】其他【表单人员控件：1；部门控件：2；角色控件：3】
	Handler        []entity.Handlers //处理人列表
	Strategy       int               //处理人策略【常用审批人：1；主管（相对岗位）：2；其他：3】
	Obj            string            //扩展字段，设计中可忽略
	Relative       string            //相对岗位，设计中可忽略
}

func (genUserTaskBO *GenUserTaskBO) ExecHandlerStrategy() []entity.UserTaskBO {
	var nodeHandlerStrategy IExecNodeHandler
	switch genUserTaskBO.Strategy {
	case constant.ApprovalUserStrategyCommon:
		nodeHandlerStrategy = GenCommApproverStrategy(genUserTaskBO)
	case constant.ApprovalUserStrategyRelativePost:
		nodeHandlerStrategy = GenRelativePostStrategy(genUserTaskBO)
	case constant.ApprovalUserStrategyOther:
		nodeHandlerStrategy = GenOtherStrategy(genUserTaskBO)
	default:
		hlog.Error("实例任务[%s]节点[%s]执行生成用户任务策略设置有误，请检查配置", genUserTaskBO.InstTaskID, genUserTaskBO.NodeID)
		return nil
	}
	return nodeHandlerStrategy.GenUserTasks()
}
