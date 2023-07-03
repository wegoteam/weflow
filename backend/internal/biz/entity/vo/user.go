package vo

// UserQueryVO
// @Description: 用户查询VO
type UserQueryVO struct {
	UserName string `json:"userName" swaggertype:"string" example:""` //用户名称
}

// RoleQueryVO
// @Description: 角色查询VO
type RoleQueryVO struct {
	RoleName string `json:"roleName" swaggertype:"string" example:""` //角色名称
}

// OrgQueryVO
// @Description: 组织查询VO
type OrgQueryVO struct {
	OrgName string `json:"orgName" swaggertype:"string" example:""` //组织名称
}
