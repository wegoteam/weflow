package User

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	hertzconsts "github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/wegoteam/weflow/internal/biz/entity/vo"
	userService "github.com/wegoteam/weflow/internal/biz/handler/user"
	"github.com/wegoteam/weflow/pkg/common/entity"
)

// Register
// @Description: 注册用户中心路由
// @param: h
func Register(h *server.Hertz) {
	//用户
	userGroup := h.Group("/user")
	userGroup.POST("/list", GetUserList)
	//角色
	roleGroup := h.Group("/role")
	roleGroup.POST("/list", GetRoleList)
	//组织
	orgGroup := h.Group("/org")
	orgGroup.POST("/list", GetOrgList)
}

// GetUserList 查询用户列表
// @Summary 查询用户列表
// @Tags 用户中心
// @Description 查询用户列表
// @Accept application/json
// @Param UserQueryVO body vo.UserQueryVO true "请求参数"
// @Produce application/json
// @Success 200 {object} base.Response{data=bo.UserInfoResult} "返回结果"
// @Router /user/list [post]
func GetUserList(ctx context.Context, reqCtx *app.RequestContext) {
	// 获取请求参数
	var req vo.UserQueryVO
	reqCtx.Bind(&req)
	param := &entity.UserQueryBO{
		UserName: req.UserName,
	}
	res := userService.GetUserList(param)
	reqCtx.JSON(hertzconsts.StatusOK, res)
}

// GetRoleList 查询角色列表
// @Summary 查询角色列表
// @Tags 用户中心
// @Description 查询角色列表
// @Accept application/json
// @Param RoleQueryVO body vo.RoleQueryVO true "请求参数"
// @Produce application/json
// @Success 200 {object} base.Response{data=bo.RoleInfoResult} "返回结果"
// @Router /role/list [post]
func GetRoleList(ctx context.Context, reqCtx *app.RequestContext) {
	// 获取请求参数
	var req vo.RoleQueryVO
	reqCtx.Bind(&req)
	param := &entity.RoleQueryBO{
		RoleName: req.RoleName,
	}
	res := userService.GetRoleList(param)
	reqCtx.JSON(hertzconsts.StatusOK, res)
}

// GetOrgList 查询组织列表
// @Summary 查询组织列表
// @Tags 用户中心
// @Description 查询组织列表
// @Accept application/json
// @Param OrgQueryVO body vo.OrgQueryVO true "请求参数"
// @Produce application/json
// @Success 200 {object} base.Response{data=bo.OrgInfoResult} "返回结果"
// @Router /org/list [post]
func GetOrgList(ctx context.Context, reqCtx *app.RequestContext) {
	// 获取请求参数
	var req vo.OrgQueryVO
	reqCtx.Bind(&req)
	param := &entity.OrgQueryBO{
		OrgName: req.OrgName,
	}
	res := userService.GetOrgList(param)
	reqCtx.JSON(hertzconsts.StatusOK, res)
}
