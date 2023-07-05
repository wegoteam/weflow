package user

import (
	"github.com/wegoteam/weflow/internal/base"
	"github.com/wegoteam/weflow/internal/biz/entity/bo"
	"github.com/wegoteam/weflow/internal/consts"
	weflowApi "github.com/wegoteam/weflow/pkg/api"
	"github.com/wegoteam/weflow/pkg/common/entity"
	"github.com/wegoteam/weflow/pkg/common/utils"
)

// GetUserList
// @Description: 查询用户列表
// @param: param
func GetUserList(param *entity.UserQueryBO) *base.Response {
	result, err := weflowApi.GetUserList(param)
	if err != nil {
		return base.Fail(consts.ERROR, err.Error())
	}
	var userResults = make([]bo.UserInfoResult, 0)
	for _, user := range result {
		param := bo.UserInfoResult{
			ID:         user.ID,
			UserID:     user.UserID,
			UserName:   user.UserName,
			Password:   user.Password,
			Phone:      user.Phone,
			Email:      user.Email,
			OrgID:      user.OrgID,
			Status:     user.Status,
			Remark:     user.Remark,
			CreateUser: user.CreateUser,
			UpdateUser: user.UpdateUser,
			CreateTime: utils.TimeToString(user.CreateTime),
			UpdateTime: utils.TimeToString(user.UpdateTime),
		}
		userResults = append(userResults, param)
	}
	return base.OK(userResults)
}

// GetRoleList
// @Description: 查询角色列表
// @param: param *entity.RoleQueryBO
// @return *base.Response
func GetRoleList(param *entity.RoleQueryBO) *base.Response {
	result, err := weflowApi.GetRoleList(param)
	if err != nil {
		return base.Fail(consts.ERROR, err.Error())
	}
	var roleResults = make([]bo.RoleInfoResult, 0)
	for _, role := range result {
		param := bo.RoleInfoResult{
			ID:         role.ID,
			RoleID:     role.RoleID,
			ParentID:   role.ParentID,
			RoleName:   role.RoleName,
			Status:     role.Status,
			Remark:     role.Remark,
			CreateUser: role.CreateUser,
			UpdateUser: role.UpdateUser,
			CreateTime: utils.TimeToString(role.CreateTime),
			UpdateTime: utils.TimeToString(role.UpdateTime),
		}
		roleResults = append(roleResults, param)
	}
	return base.OK(roleResults)
}

// GetRoleTree
// @Description: 查询角色树
// @param: param
// @return *base.Response
func GetRoleTree(param *entity.RoleQueryBO) *base.Response {
	result, err := weflowApi.GetRoleTree(param)
	if err != nil {
		return base.Fail(consts.ERROR, err.Error())
	}
	var roleResults = make([]bo.RoleInfoTreeResult, 0)
	for _, role := range result {
		param := bo.RoleInfoTreeResult{
			ID:         role.ID,
			RoleID:     role.RoleID,
			ParentID:   role.ParentID,
			RoleName:   role.RoleName,
			Status:     role.Status,
			Remark:     role.Remark,
			CreateUser: role.CreateUser,
			UpdateUser: role.UpdateUser,
			CreateTime: utils.TimeToString(role.CreateTime),
			UpdateTime: utils.TimeToString(role.UpdateTime),
			Children:   buildRoleChildren(role.Children),
		}
		roleResults = append(roleResults, param)
	}
	return base.OK(roleResults)
}

// buildRoleChildren
// @Description: 获取角色子节点
// @param: roleResult
// @return []bo.RoleInfoTreeResult
func buildRoleChildren(roleResults []*entity.RoleInfoTreeResult) []bo.RoleInfoTreeResult {
	var roles = make([]bo.RoleInfoTreeResult, 0)
	if roleResults == nil {
		return roles
	}
	for _, role := range roleResults {
		param := bo.RoleInfoTreeResult{
			ID:         role.ID,
			RoleID:     role.RoleID,
			ParentID:   role.ParentID,
			RoleName:   role.RoleName,
			Status:     role.Status,
			Remark:     role.Remark,
			CreateUser: role.CreateUser,
			UpdateUser: role.UpdateUser,
			CreateTime: utils.TimeToString(role.CreateTime),
			UpdateTime: utils.TimeToString(role.UpdateTime),
			Children:   buildRoleChildren(role.Children),
		}
		roles = append(roles, param)
	}
	return roles
}

// GetRoleUserList
// @Description: 查询角色用户列表
// @param: param
// @return *base.Response
func GetRoleUserList(param *entity.RoleQueryBO) *base.Response {
	result, err := weflowApi.GetAllRoleUserList()
	if err != nil {
		return base.Fail(consts.ERROR, err.Error())
	}
	var roleUserResults = make([]bo.RoleUserTreeResult, 0)
	for _, role := range result {
		param := bo.RoleUserTreeResult{
			ID:         role.ID,
			RoleID:     role.RoleID,
			ParentID:   role.ParentID,
			RoleName:   role.RoleName,
			Status:     role.Status,
			Remark:     role.Remark,
			CreateUser: role.CreateUser,
			UpdateUser: role.UpdateUser,
			CreateTime: utils.TimeToString(role.CreateTime),
			UpdateTime: utils.TimeToString(role.UpdateTime),
			Users:      buildUserList(role.Users),
			Children:   buildRoleUserChildren(role.Children),
		}
		roleUserResults = append(roleUserResults, param)
	}
	return base.OK(roleUserResults)
}

// GetRoleUserTree
// @Description: 查询角色用户树
// @param: param
// @return *base.Response
func GetRoleUserTree(param *entity.RoleQueryBO) *base.Response {
	result, err := weflowApi.GetAllRoleUserTree()
	if err != nil {
		return base.Fail(consts.ERROR, err.Error())
	}
	var roleUserResults = make([]bo.RoleUserTreeResult, 0)
	for _, role := range result {
		param := bo.RoleUserTreeResult{
			ID:         role.ID,
			RoleID:     role.RoleID,
			ParentID:   role.ParentID,
			RoleName:   role.RoleName,
			Status:     role.Status,
			Remark:     role.Remark,
			CreateUser: role.CreateUser,
			UpdateUser: role.UpdateUser,
			CreateTime: utils.TimeToString(role.CreateTime),
			UpdateTime: utils.TimeToString(role.UpdateTime),
			Users:      buildUserList(role.Users),
			Children:   buildRoleUserChildren(role.Children),
		}
		roleUserResults = append(roleUserResults, param)
	}
	return base.OK(roleUserResults)
}

// buildRoleUserChildren
// @Description: 获取角色用户子节点
// @param: roleResults
// @return []bo.RoleUserTreeResult
func buildRoleUserChildren(roleResults []*entity.RoleUserTreeResult) []bo.RoleUserTreeResult {
	var roleUserList = make([]bo.RoleUserTreeResult, 0)
	if roleResults == nil {
		return roleUserList
	}
	for _, role := range roleResults {
		param := bo.RoleUserTreeResult{
			ID:         role.ID,
			RoleID:     role.RoleID,
			ParentID:   role.ParentID,
			RoleName:   role.RoleName,
			Status:     role.Status,
			Remark:     role.Remark,
			CreateUser: role.CreateUser,
			UpdateUser: role.UpdateUser,
			CreateTime: utils.TimeToString(role.CreateTime),
			UpdateTime: utils.TimeToString(role.UpdateTime),
			Users:      buildUserList(role.Users),
			Children:   buildRoleUserChildren(role.Children),
		}
		roleUserList = append(roleUserList, param)
	}
	return roleUserList
}

// buildUserList
// @Description: 获取角色用户列表
// @param: users
// @return []bo.UserInfoResult
func buildUserList(users []entity.UserInfoResult) []bo.UserInfoResult {
	var userList = make([]bo.UserInfoResult, 0)
	if users == nil {
		return userList
	}
	for _, user := range users {
		param := bo.UserInfoResult{
			ID:         user.ID,
			UserID:     user.UserID,
			UserName:   user.UserName,
			Password:   user.Password,
			Phone:      user.Phone,
			Email:      user.Email,
			OrgID:      user.OrgID,
			Status:     user.Status,
			Remark:     user.Remark,
			CreateUser: user.CreateUser,
			UpdateUser: user.UpdateUser,
			CreateTime: utils.TimeToString(user.CreateTime),
			UpdateTime: utils.TimeToString(user.UpdateTime),
		}
		userList = append(userList, param)
	}
	return userList
}

// GetOrgList
// @Description: 查询组织列表
// @param: param *entity.OrgQueryBO
// @return *base.Response
func GetOrgList(param *entity.OrgQueryBO) *base.Response {
	result, err := weflowApi.GetOrgList(param)
	if err != nil {
		return base.Fail(consts.ERROR, err.Error())
	}
	var orgResults = make([]bo.OrgInfoResult, 0)
	for _, org := range result {
		param := bo.OrgInfoResult{
			ID:         org.ID,
			OrgID:      org.OrgID,
			ParentID:   org.ParentID,
			OrgName:    org.OrgName,
			Status:     org.Status,
			Remark:     org.Remark,
			CreateUser: org.CreateUser,
			UpdateUser: org.UpdateUser,
			CreateTime: utils.TimeToString(org.CreateTime),
			UpdateTime: utils.TimeToString(org.UpdateTime),
		}
		orgResults = append(orgResults, param)
	}
	return base.OK(orgResults)
}

// GetOrgTree
// @Description: 查询组织树
// @param: param *entity.OrgQueryBO
// @return *base.Response
func GetOrgTree(param *entity.OrgQueryBO) *base.Response {
	result, err := weflowApi.GetOrgTree(param)
	if err != nil {
		return base.Fail(consts.ERROR, err.Error())
	}
	var orgResults = make([]bo.OrgInfoTreeResult, 0)
	for _, org := range result {
		param := bo.OrgInfoTreeResult{
			ID:         org.ID,
			OrgID:      org.OrgID,
			ParentID:   org.ParentID,
			OrgName:    org.OrgName,
			Status:     org.Status,
			Remark:     org.Remark,
			CreateUser: org.CreateUser,
			UpdateUser: org.UpdateUser,
			CreateTime: utils.TimeToString(org.CreateTime),
			UpdateTime: utils.TimeToString(org.UpdateTime),
			Children:   buildOrgChildren(org.Children),
		}
		orgResults = append(orgResults, param)
	}
	return base.OK(orgResults)
}

// buildOrgChildren
// @Description: 获取组织子节点
// @param: orgResults
// @return []bo.OrgInfoTreeResult
func buildOrgChildren(orgResults []*entity.OrgInfoTreeResult) []bo.OrgInfoTreeResult {
	var orgs = make([]bo.OrgInfoTreeResult, 0)
	if orgResults == nil {
		return orgs
	}
	for _, org := range orgResults {
		param := bo.OrgInfoTreeResult{
			ID:         org.ID,
			OrgID:      org.OrgID,
			ParentID:   org.ParentID,
			OrgName:    org.OrgName,
			Status:     org.Status,
			Remark:     org.Remark,
			CreateUser: org.CreateUser,
			UpdateUser: org.UpdateUser,
			CreateTime: utils.TimeToString(org.CreateTime),
			UpdateTime: utils.TimeToString(org.UpdateTime),
			Children:   buildOrgChildren(org.Children),
		}
		orgs = append(orgs, param)
	}
	return orgs
}

// GetOrgUserTree
// @Description: 查询组织用户树
// @param: param
// @return *base.Response
func GetOrgUserTree(param *entity.OrgQueryBO) *base.Response {
	result, err := weflowApi.GetAllOrgUserTree()
	if err != nil {
		return base.Fail(consts.ERROR, err.Error())
	}
	var orgResults = make([]bo.OrgUserTreeResult, 0)
	for _, org := range result {
		param := bo.OrgUserTreeResult{
			ID:         org.ID,
			OrgID:      org.OrgID,
			ParentID:   org.ParentID,
			OrgName:    org.OrgName,
			Status:     org.Status,
			Remark:     org.Remark,
			CreateUser: org.CreateUser,
			UpdateUser: org.UpdateUser,
			CreateTime: utils.TimeToString(org.CreateTime),
			UpdateTime: utils.TimeToString(org.UpdateTime),
			Users:      buildUserList(org.Users),
			Children:   buildOrgUserChildren(org.Children),
		}
		orgResults = append(orgResults, param)
	}
	return base.OK(orgResults)
}

// buildOrgUserChildren
// @Description: 获取组织用户子节点
// @param: orgResults
// @return []bo.OrgUserTreeResult
func buildOrgUserChildren(orgResults []*entity.OrgUserTreeResult) []bo.OrgUserTreeResult {
	var orgs = make([]bo.OrgUserTreeResult, 0)
	if orgResults == nil {
		return orgs
	}
	for _, org := range orgResults {
		param := bo.OrgUserTreeResult{
			ID:         org.ID,
			OrgID:      org.OrgID,
			ParentID:   org.ParentID,
			OrgName:    org.OrgName,
			Status:     org.Status,
			Remark:     org.Remark,
			CreateUser: org.CreateUser,
			UpdateUser: org.UpdateUser,
			CreateTime: utils.TimeToString(org.CreateTime),
			UpdateTime: utils.TimeToString(org.UpdateTime),
			Users:      buildUserList(org.Users),
			Children:   buildOrgUserChildren(org.Children),
		}
		orgs = append(orgs, param)
	}
	return orgs
}

// GetOrgUserList
// @Description: 查询组织用户列表
// @param: param
// @return *base.Response
func GetOrgUserList(param *entity.OrgQueryBO) *base.Response {
	result, err := weflowApi.GetAllOrgUserList()
	if err != nil {
		return base.Fail(consts.ERROR, err.Error())
	}
	var orgResults = make([]bo.OrgUserTreeResult, 0)
	for _, org := range result {
		param := bo.OrgUserTreeResult{
			ID:         org.ID,
			OrgID:      org.OrgID,
			ParentID:   org.ParentID,
			OrgName:    org.OrgName,
			Status:     org.Status,
			Remark:     org.Remark,
			CreateUser: org.CreateUser,
			UpdateUser: org.UpdateUser,
			CreateTime: utils.TimeToString(org.CreateTime),
			UpdateTime: utils.TimeToString(org.UpdateTime),
			Users:      buildUserList(org.Users),
		}
		orgResults = append(orgResults, param)
	}
	return base.OK(orgResults)
}
