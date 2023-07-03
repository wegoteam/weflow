package api

import (
	"github.com/wegoteam/weflow/pkg/common/entity"
	"github.com/wegoteam/weflow/pkg/service"
)

// GetUserList
// @Description: 查询用户列表
// @param: param
// @return error
func GetUserList(param *entity.UserQueryBO) ([]entity.UserInfoResult, error) {
	return service.GetUserList(param)
}

// GetRoleList
// @Description: 查询角色列表
// @param: param
// @return []entity.RoleInfoResult
// @return error
func GetRoleList(param *entity.RoleQueryBO) ([]entity.RoleInfoResult, error) {
	return service.GetRoleList(param)
}

// GetOrgList
// @Description: 查询组织列表
// @param: param
// @return []entity.OrgInfoResult
// @return error
func GetOrgList(param *entity.OrgQueryBO) ([]entity.OrgInfoResult, error) {
	return service.GetOrgList(param)
}
