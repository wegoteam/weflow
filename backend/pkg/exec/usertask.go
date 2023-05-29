package exec

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/wegoteam/weflow/pkg/common/constant"
	"github.com/wegoteam/weflow/pkg/common/entity"
	"time"
)

// ExecUserTask
// @Description: 生成用户任务
//处理人策略【常用审批人：1；主管（相对岗位）：2；其他：3】
//常用审批人【指定成员：1；发起人自己：2；发起人自选：3：角色：4；部门：5】主管（相对岗位）【直属主管：1；部门主管：2；连续多级主管：3；部门控件对应主管：4】其他【表单人员控件：1；部门控件：2；角色控件：3】
//  @param instNodeTask
//  @param nodeHandler
//  @return []entity.UserTaskBO
func ExecUserTask(execution Execution, instNodeTask entity.InstNodeTaskBO, nodeHandler entity.NodeHandler) []entity.UserTaskBO {
	userTasks := make([]entity.UserTaskBO, 0)
	//生成用户任务
	var genUserTaskBO = GenUserTaskBO{
		InstTaskID:     execution.InstTaskID,
		NodeTaskID:     instNodeTask.NodeTaskID,
		ApproveType:    instNodeTask.ApproveType,
		NoneHandler:    instNodeTask.NoneHandler,
		AppointHandler: instNodeTask.AppointHandler,
		HandleMode:     instNodeTask.HandleMode,
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
	if genUserTasks == nil || len(genUserTasks) == 0 {
		hlog.Warnf("实例任务[%s]的流程定义[%s]执行节点[%s]节点名称[%s]获取用户任务为空，请检查节点处理人策略配置", execution.InstTaskID, execution.ProcessDefId, instNodeTask.NodeID, instNodeTask.NodeName)
		return userTasks
	}
	userTasks = append(userTasks, genUserTasks...)
	return userTasks
}

// IExecNodeHandler
// @Description: 执行节点处理人接口
type IExecNodeHandler interface {
	// genUserTasks
	// @Description: 生成用户任务
	// @return []entity.UserTaskBO
	genUserTasks() []entity.UserTaskBO
}

// GenUserTaskBO
// @Description: 生成用户任务BO
type GenUserTaskBO struct {
	InstTaskID     string            //实例任务id
	NodeTaskID     string            //节点任务id
	ApproveType    int32             // 审批类型【人工审批：1；自动通过：2；自动拒绝】默认人工审批1
	NoneHandler    int32             // 审批人为空时【自动通过：1；自动转交管理员：2；指定审批人：3】默认自动通过1
	AppointHandler string            // 审批人为空时指定审批人ID
	HandleMode     int32             // 审批方式【依次审批：1、会签（需要完成人数的审批人同意或拒绝才可完成节点）：2、或签（其中一名审批人同意或拒绝即可）：3】默认会签2
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

// ExecHandlerStrategy
// @Description: 执行生成用户任务策略
// @receiver genUserTaskBO
// @return []entity.UserTaskBO
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
	return nodeHandlerStrategy.genUserTasks()
}
