package bo

type UserInfoResult struct {
	ID         int64  `json:"id"`         // 唯一id
	UserID     string `json:"userID"`     // 用户id
	UserName   string `json:"userName"`   // 用户名称
	Password   string `json:"password"`   // 密码
	Phone      string `json:"phone"`      // 手机号
	Email      string `json:"email"`      // 邮箱
	OrgID      string `json:"orgID"`      // 组织id
	Status     int32  `json:"status"`     // 状态【1：未启用；2：已启用；3：锁定；】
	Remark     string `json:"remark"`     // 描述
	CreateUser string `json:"createUser"` // 创建人
	UpdateUser string `json:"updateUser"` // 更新人
	CreateTime string `json:"createTime"` // 创建时间
	UpdateTime string `json:"updateTime"` // 更新时间
}

// RoleInfoResult
// @Description: 角色信息结果
type RoleInfoResult struct {
	ID         int64  `json:"id"`         // 唯一id
	RoleID     string `json:"roleID"`     // 角色id
	ParentID   string `json:"parentID"`   // 角色父id
	RoleName   string `json:"roleName"`   // 角色名称
	Status     int32  `json:"status"`     // 状态【1：未启用；2：已启用；3：锁定；】
	Remark     string `json:"remark"`     // 描述
	CreateUser string `json:"createUser"` // 创建人
	UpdateUser string `json:"updateUser"` // 更新人
	CreateTime string `json:"createTime"` // 创建时间
	UpdateTime string `json:"updateTime"` // 更新时间
}

// RoleInfoTreeResult
// @Description: 角色信息结果
type RoleInfoTreeResult struct {
	ID         int64                `json:"id"`         // 唯一id
	RoleID     string               `json:"roleID"`     // 角色id
	ParentID   string               `json:"parentID"`   // 角色父id
	RoleName   string               `json:"roleName"`   // 角色名称
	Status     int32                `json:"status"`     // 状态【1：未启用；2：已启用；3：锁定；】
	Remark     string               `json:"remark"`     // 描述
	CreateUser string               `json:"createUser"` // 创建人
	UpdateUser string               `json:"updateUser"` // 更新人
	CreateTime string               `json:"createTime"` // 创建时间
	UpdateTime string               `json:"updateTime"` // 更新时间
	Children   []RoleInfoTreeResult `json:"children"`   // 子角色
}

// RoleUserTreeResult
// @Description: 角色用户树结果
type RoleUserTreeResult struct {
	ID         int64                `json:"id"`         // 唯一id
	RoleID     string               `json:"roleID"`     // 角色id
	ParentID   string               `json:"parentID"`   // 角色父id
	RoleName   string               `json:"roleName"`   // 角色名称
	Status     int32                `json:"status"`     // 状态【1：未启用；2：已启用；3：锁定；】
	Remark     string               `json:"remark"`     // 描述
	CreateUser string               `json:"createUser"` // 创建人
	UpdateUser string               `json:"updateUser"` // 更新人
	CreateTime string               `json:"createTime"` // 创建时间
	UpdateTime string               `json:"updateTime"` // 更新时间
	Users      []UserInfoResult     `json:"users"`      // 角色用户
	Children   []RoleUserTreeResult `json:"children"`   // 子角色
}

// OrgInfoResult
// @Description: 组织信息结果
type OrgInfoResult struct {
	ID         int64  `json:"id"`         // 唯一id
	OrgID      string `json:"orgID"`      // 组id
	ParentID   string `json:"parentID"`   // 组父id
	OrgName    string `json:"orgName"`    // 组名称
	Status     int32  `json:"status"`     // 状态【1：未启用；2：已启用；3：锁定；】
	Remark     string `json:"remark"`     // 描述
	CreateUser string `json:"createUser"` // 创建人
	UpdateUser string `json:"updateUser"` // 更新人
	CreateTime string `json:"createTime"` // 创建时间
	UpdateTime string `json:"updateTime"` // 更新时间
}

// OrgInfoTreeResult
// @Description: 组织信息结果
type OrgInfoTreeResult struct {
	ID         int64               `json:"id"`         // 唯一id
	OrgID      string              `json:"orgID"`      // 组id
	ParentID   string              `json:"parentID"`   // 组父id
	OrgName    string              `json:"orgName"`    // 组名称
	Status     int32               `json:"status"`     // 状态【1：未启用；2：已启用；3：锁定；】
	Remark     string              `json:"remark"`     // 描述
	CreateUser string              `json:"createUser"` // 创建人
	UpdateUser string              `json:"updateUser"` // 更新人
	CreateTime string              `json:"createTime"` // 创建时间
	UpdateTime string              `json:"updateTime"` // 更新时间
	Children   []OrgInfoTreeResult `json:"children"`   // 子组织
}

// OrgUserTreeResult
// @Description: 组织用户树结果
type OrgUserTreeResult struct {
	ID         int64               `json:"id"`         // 唯一id
	OrgID      string              `json:"orgID"`      // 组id
	ParentID   string              `json:"parentID"`   // 组父id
	OrgName    string              `json:"orgName"`    // 组名称
	Status     int32               `json:"status"`     // 状态【1：未启用；2：已启用；3：锁定；】
	Remark     string              `json:"remark"`     // 描述
	CreateUser string              `json:"createUser"` // 创建人
	UpdateUser string              `json:"updateUser"` // 更新人
	CreateTime string              `json:"createTime"` // 创建时间
	UpdateTime string              `json:"updateTime"` // 更新时间
	Users      []UserInfoResult    `json:"users"`      // 角色用户
	Children   []OrgUserTreeResult `json:"children"`   // 子组织
}
