package exec

import (
	"github.com/wegoteam/weflow/pkg/common/entity"
	"github.com/wegoteam/wepkg/snowflake"
)

/**
生成用户任务
处理人策略【常用审批人：1；主管（相对岗位）：2；其他：3】
常用审批人【指定成员：1；发起人自己：2；发起人自选：3：角色：4；部门：5】主管（相对岗位）【直属主管：1；部门主管：2；连续多级主管：3；部门控件对应主管：4】其他【表单人员控件：1；部门控件：2；角色控件：3】
*/
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
			Status:       1,
			CreateTime:   instNodeTask.CreateTime,
			UpdateTime:   instNodeTask.UpdateTime,
			HandleTime:   instNodeTask.CreateTime,
			OpUserID:     handler.ID,
			OpUserName:   handler.Name,
			Opinion:      1,
			OpinionDesc:  "",
		}
		for i := 0; i < 10; i++ {

			userTasks = append(userTasks, userTask)
		}
	}

	return userTasks
}

type IExecNodeHandler interface {
	GetUserTasks(nodeHandler entity.NodeHandler) []entity.UserTaskBO
}

type ICommApproverStrategy interface {
	GetUserTasks(nodeHandler entity.NodeHandler) []entity.UserTaskBO
}

type IRelativeStrategy interface {
	GetUserTasks(nodeHandler entity.NodeHandler) []entity.UserTaskBO
}

type IOtherStrategy interface {
	GetUserTasks(nodeHandler entity.NodeHandler) []entity.UserTaskBO
}
