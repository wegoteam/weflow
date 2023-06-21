package Model

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	hertzconsts "github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/wegoteam/weflow/internal/biz/entity/vo"
	modelService "github.com/wegoteam/weflow/internal/biz/handler/model"
	"github.com/wegoteam/weflow/internal/consts"
	"github.com/wegoteam/weflow/pkg/common/entity"
	"time"
)

// Register
// @Description: 注册模板路由
// @param: h
func Register(h *server.Hertz) {
	modelGroup := h.Group("/model")
	modelGroup.GET("/list", GetModelList)
	modelGroup.GET("/page", PageModels)
	modelGroup.GET("/group/list", GetModelGroups)
	modelGroup.POST("/group/add", AddModelGroup)
	modelGroup.POST("/group/edit", EditModelGroup)
	modelGroup.POST("/group/del", DelModelGroup)
}

// GetModelList 获取模板列表
// @Summary 获取模板列表
// @Tags 模板
// @Description 获取模板列表
// @Accept application/json
// @Produce application/json
// @Success 200 {object} bo.ModelDetailResult
// @Router /model/list [get]
func GetModelList(ctx context.Context, rc *app.RequestContext) {
	res := modelService.GetModelList()
	rc.JSON(hertzconsts.StatusOK, res)
}

// PageModels 分页获取模板列表
// @Summary 分页获取模板列表
// @Tags 模板
// @Description 分页获取模板列表
// @Accept application/json
// @Produce application/json
// @Success 200 {object} bo.ModelDetailResult
// @Router /model/page [get]
func PageModels(ctx context.Context, rc *app.RequestContext) {
	res := modelService.GetModelList()
	rc.JSON(hertzconsts.StatusOK, res)
}

// GetModelGroups 查询获取模板组列表
// @Summary 查询获取模板组列表
// @Tags 模板
// @Description 查询获取模板组列表
// @Accept application/json
// @Produce application/json
// @Success 200 {object} bo.ModelGroupResult
// @Router /model/group/list [get]
func GetModelGroups(ctx context.Context, rc *app.RequestContext) {
	res := modelService.GetModelGroupList()
	rc.JSON(hertzconsts.StatusOK, res)
}

// AddModelGroup 添加模板组
// @Summary 添加模板组
// @Tags 模板
// @Param ModelGroupAddVO body vo.ModelGroupAddVO true "请求参数"
// @Description 添加模板组
// @Accept application/json
// @Produce application/json
// @Success 200 {object} base.Response{}
// @Router /model/group/add [post]
func AddModelGroup(ctx context.Context, rc *app.RequestContext) {
	var req vo.ModelGroupAddVO
	rc.Bind(&req)
	now := time.Now()
	param := &entity.ModelGroupAddBO{
		GroupName:  req.GroupName,
		Remark:     req.Remark,
		CreateUser: consts.UserID,
		UpdateUser: consts.UserID,
		CreateTime: now,
		UpdateTime: now,
	}
	res := modelService.AddModelGroup(param)
	rc.JSON(hertzconsts.StatusOK, res)
}

// EditModelGroup 编辑模板组
// @Summary 编辑模板组
// @Tags 模板
// @Param ModelGroupEditVO body vo.ModelGroupEditVO true "请求参数"
// @Description 编辑模板组
// @Accept application/json
// @Produce application/json
// @Success 200 {object} base.Response{}
// @Router /model/group/edit [post]
func EditModelGroup(ctx context.Context, rc *app.RequestContext) {
	var req vo.ModelGroupEditVO
	rc.Bind(&req)
	now := time.Now()
	param := &entity.ModelGroupEditBO{
		GroupID:    req.GroupID,
		GroupName:  req.GroupName,
		Remark:     req.Remark,
		UpdateUser: consts.UserID,
		UpdateTime: now,
	}
	res := modelService.EditModelGroup(param)
	rc.JSON(hertzconsts.StatusOK, res)
}

// DelModelGroup 删除模板组
// @Summary 删除模板组
// @Tags 模板
// @Param ModelGroupDelVO body vo.ModelGroupDelVO true "请求参数"
// @Description 删除模板组
// @Accept application/json
// @Produce application/json
// @Success 200 {object} base.Response{}
// @Router /model/group/del [post]
func DelModelGroup(ctx context.Context, rc *app.RequestContext) {
	var req vo.ModelGroupDelVO
	rc.Bind(&req)
	param := &entity.ModelGroupDelBO{
		GroupID: req.GroupID,
	}
	res := modelService.DelModelGroup(param)
	rc.JSON(hertzconsts.StatusOK, res)
}
