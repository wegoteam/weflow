package entity

import "time"

type UserInfoResult struct {
	ID         int64     `json:"id"`         // 唯一id
	UserID     string    `json:"userId"`     // 用户id
	UserName   string    `json:"userName"`   // 用户名称
	Password   string    `json:"password"`   // 密码
	Phone      string    `json:"phone"`      // 手机号
	Email      string    `json:"email"`      // 邮箱
	OrgID      string    `json:"orgId"`      // 组织id
	Status     int32     `json:"status"`     // 状态【1：未启用；2：已启用；3：锁定；】
	Remark     string    `json:"remark"`     // 描述
	CreateUser string    `json:"createUser"` // 创建人
	UpdateUser string    `json:"updateUser"` // 更新人
	CreateTime time.Time `json:"createTime"` // 创建时间
	UpdateTime time.Time `json:"updateTime"` // 更新时间
}
