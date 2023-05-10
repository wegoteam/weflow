package entity

// NodeModelEntity 节点类型实体
type NodeModelEntity struct {
	NodeModel int8   `json:"nodeModel"` // 节点类型
	NodeName  string `json:"nodeName"`  // 节点名称
	NodeID    string `json:"nodeId"`    // 节点ID
	ParentID  string `json:"parentId"`  // 父节点ID

	//审批、办理、抄送节点
	ApproveType    int         `json:"approveType"`    //审批类型【人工审批：1；自动通过：2；自动拒绝】默认人工审批1
	FormPer        []FormPer   `json:"formPer"`        // 表单权限
	NodeSetting    NodeSetting `json:"nodeSetting"`    // 节点设置
	NodeHandler    NodeHandler `json:"nodeHandler"`    // 节点处理人
	NoneHandler    int         `json:"noneHandler"`    //审批人为空时【自动通过：1；自动转交管理员：2；指定审批人：3】默认自动通过1--数字
	AppointHandler string      `json:"appointHandler"` //审批人为空时指定审批人ID
	HandleMode     int         `json:"handleMode"`     //审批方式【依次审批：1、会签（需要完成人数的审批人同意或拒绝才可完成节点）：2、或签（其中一名审批人同意或拒绝即可）：3】默认会签2
	FinishMode     int         `json:"finishMode"`     //完成人数：依次审批默认0所有人不可选人，会签默认0所有人（可选人大于0），或签默认1一个人（可选人大于0）

	//条件节点
	Level          int    `json:"level"`          //优先级，分支执行方式为多分支处理方式无优先级应为0
	ConditionGroup string `json:"conditionGroup"` //条件组前端描述展示条件组
	ConditionExpr  string `json:"conditionExpr"`  //条件组解析后的表达式

	//分支节点
	BranchMode    int                 `json:"branchMode"`    // 分支执行方式【单分支：1；多分支：2】默认多分支2
	DefaultBranch int                 `json:"defaultBranch"` // 单分支处理需要默认分支，在条件优先级无法处理时候执行默认分支，取值分支下标
	Children      [][]NodeModelEntity `json:"children"`      // 子节点
}

type NodeHandler struct {
	Type     string     `json:"type"`     //常用审批人【指定成员：1；发起人自己：2；发起人自选：3：角色：4；部门：5】主管（相对岗位）【直属主管：1；部门主管：2；连续多级主管：3；部门控件对应主管：4】其他【表单人员控件：1；部门控件：2；角色控件：3】
	Handlers []Handlers `json:"handlers"` // 处理人列表
	Strategy int        `json:"strategy"` // 处理人策略【常用审批人：1；主管（相对岗位）：2；其他：3】
	Obj      string     `json:"obj"`      //扩展字段，设计中可忽略
	Relative string     `json:"relative"`
}

type Handlers struct {
	ID   string `json:"id"`   //处理人ID
	Name string `json:"name"` //处理人名称
	Sort int    `json:"sort"` //排序
}

type NodeSetting struct {
	ExecCheck   string `json:"execCheck"`   //审批执行校验，设计中可忽略
	Timeout     string `json:"timeout"`     // 审批限时处理，设计中可忽略
	AllowNotify string `json:"allowNotify"` // 允许发起人添加抄送人，先忽略，设计中可忽略
}

type FormPer struct {
	ElemID  string `json:"elemId"`  //表单元素ID
	ElemPID string `json:"elemPId"` //表单元素父ID
	Per     int    `json:"per"`     // 表单权限【可编辑：1；只读：2；隐藏：3】默认只读2
}

type NodeModelBO struct {
	NodeModel int8   `json:"nodeModel"` // 节点类型
	NodeName  string `json:"nodeName"`  // 节点名称
	NodeID    string `json:"nodeId"`    // 节点ID
	ParentID  string `json:"parentId"`  // 父节点ID

	//审批、办理、抄送节点
	ApproveType    int         `json:"approveType"`    //审批类型【人工审批：1；自动通过：2；自动拒绝】默认人工审批1
	FormPer        []FormPer   `json:"formPer"`        // 表单权限
	NodeSetting    NodeSetting `json:"nodeSetting"`    // 节点设置
	NodeHandler    NodeHandler `json:"nodeHandler"`    // 节点处理人
	NoneHandler    int         `json:"noneHandler"`    //审批人为空时【自动通过：1；自动转交管理员：2；指定审批人：3】默认自动通过1--数字
	AppointHandler string      `json:"appointHandler"` //审批人为空时指定审批人ID
	HandleMode     int         `json:"handleMode"`     //审批方式【依次审批：1、会签（需要完成人数的审批人同意或拒绝才可完成节点）：2、或签（其中一名审批人同意或拒绝即可）：3】默认会签2
	FinishMode     int         `json:"finishMode"`     //完成人数：依次审批默认0所有人不可选人，会签默认0所有人（可选人大于0），或签默认1一个人（可选人大于0）

	//条件节点
	Level          int    `json:"level"`          //优先级，分支执行方式为多分支处理方式无优先级应为0
	ConditionGroup string `json:"conditionGroup"` //条件组前端描述展示条件组
	ConditionExpr  string `json:"conditionExpr"`  //条件组解析后的表达式

	//分支节点
	BranchMode    int        `json:"branchMode"`            // 分支执行方式【单分支：1；多分支：2】默认多分支2
	DefaultBranch int        `json:"defaultBranch"`         // 单分支处理需要默认分支，在条件优先级无法处理时候执行默认分支，取值分支下标
	ChildrenIDs   [][]string `json:"childrenIds,omitempty"` // 子节点ID

	//实例节点信息
	PreNodes    []string `json:"preNodes,omitempty"`    //上节点ID
	NextNodes   []string `json:"nextNodes,omitempty"`   //下节点ID
	LastNodes   []string `json:"lastNodes,omitempty"`   //分支节点尾节点ID
	Index       int      `json:"index,omitempty"`       // 下标
	BranchIndex int      `json:"branchIndex,omitempty"` // 分支下标
}
