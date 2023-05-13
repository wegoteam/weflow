package exec

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/gookit/slog"
	"github.com/wegoteam/weflow/pkg/common/entity"
	"github.com/wegoteam/wepkg/snowflake"
	"time"
)

// ExecBranchNode 分支节点
type ExecBranchNode struct {
	NodeModel     int8       `json:"nodeModel"`             // 节点类型
	NodeName      string     `json:"nodeName"`              // 节点名称
	NodeID        string     `json:"nodeId"`                // 节点ID
	ParentID      string     `json:"parentId"`              // 父节点ID
	BranchMode    int        `json:"branchMode"`            // 分支执行方式【单分支：1；多分支：2】默认多分支2
	DefaultBranch int        `json:"defaultBranch"`         // 单分支处理需要默认分支，在条件优先级无法处理时候执行默认分支，取值分支下标
	ChildrenIds   [][]string `json:"childrenIds,omitempty"` // 子节点ID
	//实例节点位置信息
	PreNodes    []string `json:"preNodes,omitempty"`    //上节点ID
	NextNodes   []string `json:"nextNodes,omitempty"`   //下节点ID
	LastNodes   []string `json:"lastNodes,omitempty"`   //分支节点尾节点ID
	Index       int      `json:"index,omitempty"`       // 下标
	BranchIndex int      `json:"branchIndex,omitempty"` // 分支下标
}

/**
实例化执行节点对象
*/
func NewBranchNode(node *entity.NodeModelBO) *ExecBranchNode {

	return &ExecBranchNode{
		NodeModel:     node.NodeModel,
		NodeName:      node.NodeName,
		NodeID:        node.NodeID,
		ParentID:      node.ParentID,
		BranchMode:    node.BranchMode,
		DefaultBranch: node.DefaultBranch,
		ChildrenIds:   node.ChildrenIDs,
		PreNodes:      node.PreNodes,
		NextNodes:     node.NextNodes,
		LastNodes:     node.LastNodes,
		Index:         node.Index,
		BranchIndex:   node.BranchIndex,
	}
}

/**
执行分支节点
分支节点三个状态：1：分支节点未完成；2：分支节点完成且存在出口；3：分支节点完成无分支出口
生成实例节点任务
执行任务
下节点
*/
func (execBranchNode *ExecBranchNode) ExecCurrNodeModel(execution *entity.Execution) ExecResult {
	hlog.Infof("实例任务[%s]的流程定义[%s]执行分支节点[%s]", execution.InstTaskID, execution.ProcessDefId, execBranchNode.NodeID)
	var branchNodes = make([]entity.NodeModelBO, 0)
	nodeTaskId := snowflake.GetSnowflakeId()

	execNodeTaskMap := execution.ExecNodeTaskMap
	processDefModel := execution.ProcessDefModel
	nodeModelMap := processDefModel.NodeModelMap
	_, ok := execNodeTaskMap[execBranchNode.NodeID]
	if !ok {
		hlog.Infof("实例任务[%s]的流程定义[%s]的分支节点[%s]未执行，生成新的实例节点任务", execution.InstTaskID, execution.ProcessDefId, execBranchNode.NodeID)
		//生成执行节点任务
		var execNodeTask = &entity.ExecNodeTaskBO{
			NodeTaskID: nodeTaskId,
			NodeModel:  execBranchNode.NodeModel,
			NodeID:     execBranchNode.NodeID,
			Status:     1,
		}
		execution.ExecNodeTaskMap[execBranchNode.NodeID] = *execNodeTask

		//生成实例节点任务
		instNodeTasks := execution.InstNodeTasks
		var instNodeTask = execBranchNode.GetInstNodeTask(execution.InstTaskID, nodeTaskId, execution.Now)
		*instNodeTasks = append(*instNodeTasks, instNodeTask)

		for _, childBranchs := range execBranchNode.ChildrenIds {
			if childBranchs == nil {
				continue
			}
			bo, hasNode := nodeModelMap[childBranchs[0]]
			if !hasNode {
				slog.Infof("节点[%v]的分支节点不存在", execBranchNode.NodeID)
			}
			branchNodes = append(branchNodes, bo)
		}

		return ExecResult{
			NextNodes: &branchNodes,
		}
	}

	//判断内存信息分支节点未执行

	return ExecResult{
		NextNodes: &branchNodes,
	}
}

/**
获取实例节点任务
*/
func (execBranchNode *ExecBranchNode) GetInstNodeTask(instTaskID, nodeTaskID string, now time.Time) entity.InstNodeTaskBO {
	//生成实例节点任务
	var instNodeTask = entity.InstNodeTaskBO{
		InstTaskID:    instTaskID,
		NodeTaskID:    nodeTaskID,
		ParentID:      execBranchNode.ParentID,
		NodeModel:     int32(execBranchNode.NodeModel),
		NodeName:      execBranchNode.NodeName,
		BranchMode:    int32(execBranchNode.BranchMode),
		DefaultBranch: int32(execBranchNode.DefaultBranch),
		Status:        1,
		CreateTime:    now,
		UpdateTime:    now,
	}

	return instNodeTask
}

func (execBranchNode *ExecBranchNode) ExecPreNodeModels(nodeModelMap map[string]entity.NodeModelBO) *[]entity.NodeModelBO {
	var preNodes = make([]entity.NodeModelBO, 0)
	if execBranchNode.PreNodes == nil {
		return &preNodes
	}
	for _, val := range execBranchNode.PreNodes {
		pre, ok := nodeModelMap[val]
		if !ok {
			slog.Infof("节点[%v]的上节点不存在", execBranchNode.NodeID)
		}
		preNodes = append(preNodes, pre)
	}
	return &preNodes
}

func (execBranchNode *ExecBranchNode) ExecNextNodeModels(nodeModelMap map[string]entity.NodeModelBO) *[]entity.NodeModelBO {
	var nextNodes = make([]entity.NodeModelBO, 0)
	if execBranchNode.NextNodes == nil {
		return &nextNodes
	}
	for _, val := range execBranchNode.NextNodes {
		next, ok := nodeModelMap[val]
		if !ok {
			slog.Infof("节点[%v]的下节点不存在", execBranchNode.NodeID)
		}
		nextNodes = append(nextNodes, next)
	}
	return &nextNodes
}
