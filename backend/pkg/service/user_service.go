package service

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/pkg/errors"
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
		hlog.Errorf("查询用户列表失败: %s", err.Error())
		return nil, errors.New("查询用户列表失败")
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
	var roleResults = make([]entity.RoleInfoResult, 0)
	roles := []model.RoleInfo{}
	err := MysqlDB.Model(&model.RoleInfo{}).Scopes(buildRoleQuery(param)).Order("role_info.create_time desc").Find(&roles).Error
	if err != nil {
		hlog.Errorf("查询角色列表失败: %v", err)
		return nil, errors.New("查询角色列表失败")
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

// GetRoleTree
// @Description: 查询角色树
// @param: param
// @return []entity.RoleInfoTreeResult
// @return error
func GetRoleTree(param *entity.RoleQueryBO) ([]*entity.RoleInfoTreeResult, error) {
	var roleResults []*entity.RoleInfoTreeResult
	roles := []model.RoleInfo{}
	err := MysqlDB.Model(&model.RoleInfo{}).Scopes(buildRoleQuery(param)).Order("role_info.create_time desc").Find(&roles).Error
	if err != nil {
		hlog.Errorf("查询角色列表失败: %v", err)
		return nil, errors.New("查询角色列表失败")
	}
	if roles == nil {
		return roleResults, nil
	}
	return builAllRoleTree(roles), nil
}

// builAllRoleTree
// @Description: 构建所有角色树
// @param: data
// @return []entity.RoleInfoTreeResult
func builAllRoleTree(data []model.RoleInfo) []*entity.RoleInfoTreeResult {
	var roles = make([]*entity.RoleInfoTreeResult, 0)
	var rootRoles = make([]*entity.RoleInfoTreeResult, 0)
	var roleMap = make(map[string]*entity.RoleInfoTreeResult, len(data))
	for _, role := range data {
		roleResult := &entity.RoleInfoTreeResult{
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
		roleMap[role.RoleID] = roleResult
		if role.ParentID == "" || role.ParentID == "0" {
			rootRoles = append(rootRoles, roleResult)
		}
		roles = append(roles, roleResult)
	}
	for _, role := range roles {
		existRole, ok := roleMap[role.ParentID]
		if !ok {
			continue
		}
		existRole.Children = append(existRole.Children, role)
	}
	return rootRoles
}

// GetAllRoleUserList
// @Description: 查询所有角色用户列表
// @return []*entity.RoleUserTreeResult
// @return error
func GetAllRoleUserList() ([]entity.RoleUserTreeResult, error) {
	var roles = make([]entity.RoleUserTreeResult, 0)
	data := []model.RoleInfo{}
	err := MysqlDB.Model(&model.RoleInfo{}).Order("role_info.create_time desc").Find(&data).Error
	if err != nil {
		hlog.Errorf("查询角色列表失败: %v", err)
		return nil, errors.New("查询角色列表失败")
	}
	if data == nil {
		return roles, nil
	}
	roleUserMap, err := getAllRoleUserMap()
	if err != nil {
		return roles, err
	}

	for _, role := range data {
		roleUserList, ok := roleUserMap[role.RoleID]
		if !ok {
			roleUserList = make([]entity.UserInfoResult, 0)
		}
		roleResult := &entity.RoleUserTreeResult{
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
			Users:      roleUserList,
		}
		roles = append(roles, *roleResult)
	}

	return roles, nil
}

// GetAllRoleUserTree
// @Description: 查询角色用户列表
// @param: param
// @return []*entity.RoleInfoTreeResult
// @return error
func GetAllRoleUserTree() ([]*entity.RoleUserTreeResult, error) {
	var rootRoles = make([]*entity.RoleUserTreeResult, 0)
	data := []model.RoleInfo{}
	err := MysqlDB.Model(&model.RoleInfo{}).Order("role_info.create_time desc").Find(&data).Error
	if err != nil {
		hlog.Errorf("查询角色列表失败: %v", err)
		return nil, errors.New("查询角色列表失败")
	}
	if data == nil {
		return rootRoles, nil
	}
	roleUserMap, err := getAllRoleUserMap()
	if err != nil {
		return rootRoles, err
	}
	var roles = make([]*entity.RoleUserTreeResult, 0)
	var roleMap = make(map[string]*entity.RoleUserTreeResult, len(data))
	for _, role := range data {
		roleUserList, ok := roleUserMap[role.RoleID]
		if !ok {
			roleUserList = make([]entity.UserInfoResult, 0)
		}
		roleResult := &entity.RoleUserTreeResult{
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
			Users:      roleUserList,
		}
		roleMap[role.RoleID] = roleResult
		if role.ParentID == "" || role.ParentID == "0" {
			rootRoles = append(rootRoles, roleResult)
		}
		roles = append(roles, roleResult)
	}
	for _, role := range roles {
		existRole, ok := roleMap[role.ParentID]
		if !ok {
			continue
		}
		existRole.Children = append(existRole.Children, role)
	}
	return rootRoles, nil
}

// getAllRoleUserMap
// @Description: 查询所有角色用户map
// @return map[string][]entity.UserInfoResult
// @return error
func getAllRoleUserMap() (map[string][]entity.UserInfoResult, error) {
	var roleUserMap = make(map[string][]entity.UserInfoResult)
	userRoles := []model.UserRoleLink{}
	userRolesErr := MysqlDB.Model(&model.UserRoleLink{}).Order("create_time desc").Find(&userRoles).Error
	if userRolesErr != nil {
		hlog.Errorf("查询用户角色列表失败: %v", userRolesErr)
		return nil, errors.New("查询角色列表失败")
	}
	users := []model.UserInfo{}
	userErr := MysqlDB.Model(&model.UserInfo{}).Order("create_time desc").Find(&users).Error
	if userErr != nil {
		hlog.Errorf("查询用户列表失败: %v", userErr)
		return nil, errors.New("查询角色列表失败")
	}
	if userRoles == nil || users == nil {
		return roleUserMap, nil
	}
	var userMap = make(map[string]entity.UserInfoResult, len(users))
	for _, user := range users {
		userMap[user.UserID] = entity.UserInfoResult{
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
	}
	for _, userRole := range userRoles {
		user, ok := userMap[userRole.UserID]
		if !ok {
			continue
		}
		userList, exist := roleUserMap[userRole.RoleID]
		if !exist {
			userList = make([]entity.UserInfoResult, 0)
		}
		userList = append(userList, user)
		roleUserMap[userRole.RoleID] = userList
	}
	return roleUserMap, nil
}

// GetOrgList
// @Description: 查询组织列表
// @param: param
// @return []entity.OrgInfoResult
// @return error
func GetOrgList(param *entity.OrgQueryBO) ([]entity.OrgInfoResult, error) {
	var orgResults = make([]entity.OrgInfoResult, 0)
	orgs := []model.OrganizationInfo{}
	err := MysqlDB.Model(&model.OrganizationInfo{}).Scopes(buildOrgQuery(param)).Order("organization_info.create_time desc").Find(&orgs).Error
	if err != nil {
		hlog.Errorf("查询组织列表失败: %v", err)
		return nil, errors.New("查询组织列表失败")
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

// GetOrgTree
// @Description: 查询组织树
// @param: param *entity.OrgQueryBO
// @return []entity.OrgInfoTreeResult
// @return error
func GetOrgTree(param *entity.OrgQueryBO) ([]*entity.OrgInfoTreeResult, error) {
	var orgResults = make([]*entity.OrgInfoTreeResult, 0)
	orgs := []model.OrganizationInfo{}
	err := MysqlDB.Model(&model.OrganizationInfo{}).Scopes(buildOrgQuery(param)).Order("organization_info.create_time desc").Find(&orgs).Error
	if err != nil {
		hlog.Errorf("查询组织列表失败: %v", err)
		return nil, errors.New("查询组织列表失败")
	}
	if orgs == nil {
		return orgResults, nil
	}
	return builAllOrgTree(orgs), nil
}

// builAllOrgTree
// @Description: 构建所有组织树
// @param: data
// @return []*entity.OrgInfoTreeResult
func builAllOrgTree(data []model.OrganizationInfo) []*entity.OrgInfoTreeResult {
	var orgs = make([]*entity.OrgInfoTreeResult, 0)
	var rootOrgs = make([]*entity.OrgInfoTreeResult, 0)
	var orgMap = make(map[string]*entity.OrgInfoTreeResult, len(data))
	for _, org := range data {
		orgResult := &entity.OrgInfoTreeResult{
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
		orgMap[org.OrgID] = orgResult
		if org.ParentID == "" || org.ParentID == "0" {
			rootOrgs = append(rootOrgs, orgResult)
		}
		orgs = append(orgs, orgResult)
	}
	for _, org := range orgs {
		existOrg, ok := orgMap[org.ParentID]
		if !ok {
			continue
		}
		existOrg.Children = append(existOrg.Children, org)
	}
	return rootOrgs
}

// GetAllOrgUserTree
// @Description: 查询所有组织用户树
// @return []*entity.OrgUserTreeResult
// @return error
func GetAllOrgUserTree() ([]*entity.OrgUserTreeResult, error) {
	var rootOrgs = make([]*entity.OrgUserTreeResult, 0)
	orgList := []model.OrganizationInfo{}
	err := MysqlDB.Model(&model.OrganizationInfo{}).Order("create_time desc").Find(&orgList).Error
	if err != nil {
		hlog.Errorf("查询组织列表失败: %v", err)
		return rootOrgs, errors.New("查询组织列表失败")
	}
	if orgList == nil {
		return rootOrgs, nil
	}
	orgUserMap, orgUserErr := getAllOrgUserMap()
	if orgUserErr != nil {
		return rootOrgs, errors.New("查询组织列表失败")
	}
	var orgs = make([]*entity.OrgUserTreeResult, 0)
	var orgMap = make(map[string]*entity.OrgUserTreeResult, len(orgList))
	for _, org := range orgList {
		orgUserList, ok := orgUserMap[org.OrgID]
		if !ok {
			orgUserList = make([]entity.UserInfoResult, 0)
		}
		orgResult := &entity.OrgUserTreeResult{
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
			Users:      orgUserList,
		}
		orgMap[org.OrgID] = orgResult
		if org.ParentID == "" || org.ParentID == "0" {
			rootOrgs = append(rootOrgs, orgResult)
		}
		orgs = append(orgs, orgResult)
	}
	for _, org := range orgs {
		existOrg, ok := orgMap[org.ParentID]
		if !ok {
			continue
		}
		existOrg.Children = append(existOrg.Children, org)
	}
	return rootOrgs, nil
}

// getAllOrgUserMap
// @Description: 查询所有组织用户map
// @return map[string][]entity.UserInfoResult
// @return error
func getAllOrgUserMap() (map[string][]entity.UserInfoResult, error) {
	var orgUserMap = make(map[string][]entity.UserInfoResult)
	var users = []model.UserInfo{}
	err := MysqlDB.Model(&model.UserInfo{}).Find(&users).Error
	if err != nil {
		hlog.Errorf("查询用户列表失败: %v", err)
		return orgUserMap, errors.New("查询用户列表失败")
	}
	if users == nil {
		return orgUserMap, nil
	}
	for _, user := range users {
		userResult := entity.UserInfoResult{
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
		userList, exist := orgUserMap[user.OrgID]
		if !exist {
			userList = make([]entity.UserInfoResult, 0)
		}
		userList = append(userList, userResult)
		orgUserMap[user.OrgID] = userList
	}
	return orgUserMap, nil
}

// GetAllOrgUserList
// @Description: 查询所有组织用户列表
// @return []*entity.OrgUserTreeResult
// @return error
func GetAllOrgUserList() ([]entity.OrgUserTreeResult, error) {
	var orgs = make([]entity.OrgUserTreeResult, 0)
	orgList := []model.OrganizationInfo{}
	err := MysqlDB.Model(&model.OrganizationInfo{}).Order("create_time desc").Find(&orgList).Error
	if err != nil {
		hlog.Errorf("查询组织列表失败: %v", err)
		return orgs, errors.New("查询组织列表失败")
	}
	if orgList == nil {
		return orgs, nil
	}
	orgUserMap, orgUserErr := getAllOrgUserMap()
	if orgUserErr != nil {
		return orgs, errors.New("查询组织列表失败")
	}
	for _, org := range orgList {
		orgUserList, ok := orgUserMap[org.OrgID]
		if !ok {
			orgUserList = make([]entity.UserInfoResult, 0)
		}
		orgResult := &entity.OrgUserTreeResult{
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
			Users:      orgUserList,
		}
		orgs = append(orgs, *orgResult)
	}
	return orgs, nil
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
