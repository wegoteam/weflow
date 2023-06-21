package Model

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/wegoteam/weflow/internal/base"
	modelService "github.com/wegoteam/weflow/internal/service/model"
)

// Register
// @Description: 注册模板路由
// @param: h
func Register(h *server.Hertz) {
	modelGroup := h.Group("/model")
	{
		modelGroup.GET("/list", GetModelList)
		modelGroup.GET("/page", PageModels)
		modelGroup.GET("/group/list", PageModelGroups)
	}
}

// GetModelList 获取模板列表
// @Summary 获取模板列表
// @Tags 模板
// @Description 获取模板列表
// @Accept application/json
// @Produce application/json
// @Success 200 {object} base.Response{}
// @Router /model/list [get]
func GetModelList(ctx context.Context, rc *app.RequestContext) {
	modelList := modelService.GetModelList()
	res := &base.Response{
		Data: modelList,
	}
	rc.JSON(consts.StatusOK, res)
}

// PageModels 分页获取模板列表
// @Summary 分页获取模板列表
// @Tags 模板
// @Description 分页获取模板列表
// @Accept application/json
// @Produce application/json
// @Success 200 {object} base.Response{}
// @Router /model/page [get]
func PageModels(ctx context.Context, rc *app.RequestContext) {
	modelList := modelService.GetModelList()
	res := &base.Response{
		Data: modelList,
	}
	rc.JSON(consts.StatusOK, res)
}

// PageModelGroups 查询获取模板组列表
// @Summary 查询获取模板组列表
// @Tags 模板
// @Description 查询获取模板组列表
// @Accept application/json
// @Produce application/json
// @Success 200 {object} base.Response{}
// @Router /model/group/list [get]
func PageModelGroups(ctx context.Context, rc *app.RequestContext) {
	modelList := modelService.GetModelGroupList()
	res := &base.Response{
		Data: modelList,
	}
	rc.JSON(consts.StatusOK, res)
}
