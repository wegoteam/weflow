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
