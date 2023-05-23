package service

import (
	"github.com/wegoteam/weflow/pkg/common/entity"
	"github.com/wegoteam/weflow/pkg/model"
)

// GetRoleUserInfo
// @Description: 获取角色的用户信息
// @param roleIds
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
