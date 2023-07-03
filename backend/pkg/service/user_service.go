package service

import (
	"github.com/wegoteam/weflow/pkg/common/entity"
	"github.com/wegoteam/weflow/pkg/common/utils"
	"github.com/wegoteam/weflow/pkg/model"
	"gorm.io/gorm"
)

// GetRoleUserInfo
// @Description: 获取角色的用户信息
// @param: roleIds
// @return []entity.UserInfoResult
func GetRoleUserInfo(roleIds []string) []entity.UserInfoResult {
	userResults := make([]entity.UserInfoResult, 0)
	if roleIds == nil || len(roleIds) == 0 {
		return userResults
	}
	users := &[]model.UserInfo{}
	MysqlDB.Raw("SELECT * FROM user_info u LEFT JOIN user_role_link r ON u.user_id = r.user_id WHERE r.role_id IN ?", roleIds).Find(&users)

	if users == nil {
		return userResults
	}
	for _, user := range *users {
		userResult := &entity.UserInfoResult{
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
			CreateTime: user.CreateTime,
			UpdateTime: user.UpdateTime,
		}
		userResults = append(userResults, *userResult)
	}
	return userResults
}

// GetOrgUserInfo
// @Description: 获取组织的用户信息
// @param: roleIds
// @return []entity.UserInfoResult
func GetOrgUserInfo(orgIds []string) []entity.UserInfoResult {
	userResults := make([]entity.UserInfoResult, 0)
	if orgIds == nil || len(orgIds) == 0 {
		return userResults
	}
	users := &[]model.UserInfo{}
	MysqlDB.Model(&model.UserInfo{}).Where("org_id in (?)", orgIds).Find(&users)
	if users == nil {
		return userResults
	}
	for _, user := range *users {
		userResult := &entity.UserInfoResult{
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
			CreateTime: user.CreateTime,
			UpdateTime: user.UpdateTime,
		}
		userResults = append(userResults, *userResult)
	}
	return userResults
}

// GetUserList
// @Description: 查询用户列表
// @param: param *entity.UserQueryBO
// @return error
func GetUserList(param *entity.UserQueryBO) ([]entity.UserInfoResult, error) {
	userResult := make([]entity.UserInfoResult, 0)
	users := &[]model.UserInfo{}
	err := MysqlDB.Model(&model.UserInfo{}).Scopes(buildUserQuery(param)).Find(&users).Error
	if err != nil {
		return nil, err
	}
	if users == nil {
		return userResult, nil
	}
	for _, user := range *users {
		bo := &entity.UserInfoResult{
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
			CreateTime: user.CreateTime,
			UpdateTime: user.UpdateTime,
		}
		userResult = append(userResult, *bo)
	}
	return userResult, nil
}

// GetRoleList
// @Description: 查询角色列表
// @param: param
// @return []entity.RoleInfoResult
// @return error
func GetRoleList(param *entity.RoleQueryBO) ([]entity.RoleInfoResult, error) {
	var roleResults []entity.RoleInfoResult
	roles := []model.RoleInfo{}
	err := MysqlDB.Model(&model.RoleInfo{}).Scopes(buildRoleQuery(param)).Find(&roles).Error
	if err != nil {
		return nil, err
	}
	if roles == nil {
		return roleResults, nil
	}
	for _, role := range roles {
		roleResult := &entity.RoleInfoResult{
			ID:         role.ID,
			RoleID:     role.RoleID,
			ParentID:   role.ParentID,
			RoleName:   role.RoleName,
			Status:     role.Status,
			Remark:     role.Remark,
			CreateUser: role.CreateUser,
			UpdateUser: role.UpdateUser,
			CreateTime: role.CreateTime,
			UpdateTime: role.UpdateTime,
		}
		roleResults = append(roleResults, *roleResult)
	}
	return roleResults, nil
}

// GetOrgList
// @Description: 查询组织列表
// @param: param
// @return []entity.OrgInfoResult
// @return error
func GetOrgList(param *entity.OrgQueryBO) ([]entity.OrgInfoResult, error) {
	var orgResults []entity.OrgInfoResult
	orgs := []model.OrganizationInfo{}
	err := MysqlDB.Model(&model.OrganizationInfo{}).Scopes(buildOrgQuery(param)).Find(&orgs).Error
	if err != nil {
		return nil, err
	}
	if orgs == nil {
		return orgResults, nil
	}
	for _, org := range orgs {
		orgResult := &entity.OrgInfoResult{
			ID:         org.ID,
			OrgID:      org.OrgID,
			OrgName:    org.OrgName,
			ParentID:   org.ParentID,
			Status:     org.Status,
			Remark:     org.Remark,
			CreateUser: org.CreateUser,
			UpdateUser: org.UpdateUser,
			CreateTime: org.CreateTime,
			UpdateTime: org.UpdateTime,
		}
		orgResults = append(orgResults, *orgResult)
	}
	return orgResults, nil
}

// buildUserQuery
// @Description: 构建用户查询条件
// @param: param
// @return func(db *gorm.DB) *gorm.DB
func buildUserQuery(param *entity.UserQueryBO) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		tx := db
		if utils.IsStrNotBlank(param.UserName) {
			tx = db.Where("user_info.user_name like ?", "%"+param.UserName+"%")
		}
		return tx
	}
}

// buildRoleQuery
// @Description: 构建角色查询条件
// @param: param
// @return func(db *gorm.DB) *gorm.DB
func buildRoleQuery(param *entity.RoleQueryBO) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		tx := db
		if utils.IsStrNotBlank(param.RoleName) {
			tx = db.Where("role_info.role_name like ?", "%"+param.RoleName+"%")
		}
		return tx
	}
}

// buildOrgQuery
// @Description: 构建组织查询条件
// @param: param
// @return func(db *gorm.DB) *gorm.DB
func buildOrgQuery(param *entity.OrgQueryBO) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		tx := db
		if utils.IsStrNotBlank(param.OrgName) {
			tx = db.Where("organization_info.org_name like ?", "%"+param.OrgName+"%")
		}
		return tx
	}
}
