package exec

import "github.com/wegoteam/weflow/pkg/common/entity"

type IOtherStrategy interface {
	GetUserTasks(nodeHandler entity.NodeHandler) []entity.UserTaskBO
}
