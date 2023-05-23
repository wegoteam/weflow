package service

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/wegoteam/weflow/pkg/model"
	"testing"
)

func TestName(t *testing.T) {
	roleIds := []string{"420627966730315"}

	user := &model.UserInfo{}
	user2 := &model.UserInfo{}
	user3 := &[]model.UserInfo{}
	role := &model.RoleInfo{}

	MysqlDB.Debug().Model(&user).Where("role_id IN ?", roleIds).Association("user_id").Find(&role)
	MysqlDB.Debug().Where("role_id in (?)", roleIds).Find(&user2)

	MysqlDB.Debug().Raw("SELECT * FROM user_info u LEFT JOIN user_role_link r ON u.user_id = r.user_id WHERE r.role_id IN @roleIds",
		map[string]interface{}{"roleIds": roleIds}).Find(&user3)
	hlog.Info(user3)
}

func TestGetRoleUserInfo(t *testing.T) {
	roleIds := []string{"420627966730315"}
	userInfos := GetRoleUserInfo(roleIds)
	hlog.Info(userInfos)

}
