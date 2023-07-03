package entity

import "time"

type UserInfoResult struct {
	ID         int64     `json:"id"`         // 唯一id
	UserID     string    `json:"userID"`     // 用户id
	UserName   string    `json:"userName"`   // 用户名称
	Password   string    `json:"password"`   // 密码
	Phone      string    `json:"phone"`      // 手机号
	Email      string    `json:"email"`      // 邮箱
	OrgID      string    `json:"orgID"`      // 组织id
	Status     int32     `json:"status"`     // 状态【1：未启用；2：已启用；3：锁定；】
	Remark     string    `json:"remark"`     // 描述
	CreateUser string    `json:"createUser"` // 创建人
	UpdateUser string    `json:"updateUser"` // 更新人
	CreateTime time.Time `json:"createTime"` // 创建时间
	UpdateTime time.Time `json:"updateTime"` // 更新时间
}

// RoleInfoResult
// @Description: 角色信息结果
type RoleInfoResult struct {
	ID         int64     // 唯一id
	RoleID     string    // 角色id
	ParentID   string    // 角色父id
	RoleName   string    // 角色名称
	Status     int32     // 状态【1：未启用；2：已启用；3：锁定；】
	Remark     string    // 描述
	CreateUser string    // 创建人
	UpdateUser string    // 更新人
	CreateTime time.Time // 创建时间
	UpdateTime time.Time // 更新时间
}

// OrgInfoResult
// @Description: 组织信息结果
type OrgInfoResult struct {
	ID         int64     // 唯一id
	OrgID      string    // 组id
	ParentID   string    // 组父id
	OrgName    string    // 组名称
	Status     int32     // 状态【1：未启用；2：已启用；3：锁定；】
	Remark     string    // 描述
	CreateUser string    // 创建人
	UpdateUser string    // 更新人
	CreateTime time.Time // 创建时间
	UpdateTime time.Time // 更新时间
}

// UserQueryBO
// @Description: 用户查询BO
type UserQueryBO struct {
	UserName string `json:"userName"` //用户名称
}

// RoleQueryBO
// @Description: 角色查询BO
type RoleQueryBO struct {
	RoleName string `json:"roleName"` //角色名称
}

// OrgQueryBO
// @Description: 组织查询BO
type OrgQueryBO struct {
	OrgName string `json:"orgName"` //组织名称
}
