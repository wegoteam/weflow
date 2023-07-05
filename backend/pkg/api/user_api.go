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

// GetRoleTree
// @Description: 查询角色树
// @param: param
// @return []*entity.RoleInfoTreeResult
// @return error
func GetRoleTree(param *entity.RoleQueryBO) ([]*entity.RoleInfoTreeResult, error) {
	return service.GetRoleTree(param)
}

// GetOrgList
// @Description: 查询组织列表
// @param: param
// @return []entity.OrgInfoResult
// @return error
func GetOrgList(param *entity.OrgQueryBO) ([]entity.OrgInfoResult, error) {
	return service.GetOrgList(param)
}

// GetOrgTree
// @Description: 查询组织树
// @param: param *entity.OrgQueryBO
// @return []entity.OrgInfoResult
// @return error
func GetOrgTree(param *entity.OrgQueryBO) ([]*entity.OrgInfoTreeResult, error) {
	return service.GetOrgTree(param)
}

// GetAllRoleUserTree
// @Description: 查询所有角色用户树
// @return []*entity.RoleUserTreeResult
// @return error
func GetAllRoleUserTree() ([]*entity.RoleUserTreeResult, error) {
	return service.GetAllRoleUserTree()
}

// GetAllRoleUserList
// @Description: 查询所有角色用户列表
// @return []*entity.RoleUserTreeResult
// @return error
func GetAllRoleUserList() ([]entity.RoleUserTreeResult, error) {
	return service.GetAllRoleUserList()
}

// GetAllOrgUserTree
// @Description: 查询所有组织用户树
// @return []*entity.OrgUserTreeResult
// @return error
func GetAllOrgUserTree() ([]*entity.OrgUserTreeResult, error) {
	return service.GetAllOrgUserTree()
}

// GetAllOrgUserList
// @Description: 查询所有组织用户列表
// @return []*entity.OrgUserTreeResult
// @return error
func GetAllOrgUserList() ([]entity.OrgUserTreeResult, error) {
	return service.GetAllOrgUserList()
}
