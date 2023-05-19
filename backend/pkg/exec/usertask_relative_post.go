package exec

import "github.com/wegoteam/weflow/pkg/common/entity"

//  IRelativeStrategy
//  @Description: 主管（相对岗位）【直属主管：1；部门主管：2；连续多级主管：3；部门控件对应主管：4】
type IRelativeStrategy interface {
	GetUserTasks(nodeHandler entity.NodeHandler) []entity.UserTaskBO
}
