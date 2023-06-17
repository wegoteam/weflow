package Model

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/wegoteam/weflow/base"
	"github.com/wegoteam/weflow/internal/biz/handler/model"
)

func Register(h *server.Hertz) {
	modelGroup := h.Group("/model")
	{
		modelGroup.GET("/list", GetModelList)
		modelGroup.GET("/page", PageModels)
	}
}

// GetModelList 获取模板列表
// @Summary 获取模板列表
// @Description 获取模板列表
// @Accept application/json
// @Produce application/json
// @Router /model/list [get]
func GetModelList(ctx context.Context, req *app.RequestContext) {
	reqCtx := &base.ReqContext{
		Ctx: ctx,
		Req: req,
	}
	model.GetModelList(reqCtx)
}

// PageModels 分页获取模板列表
// @Summary 分页获取模板列表
// @Description 分页获取模板列表
// @Accept application/json
// @Produce application/json
// @Router /model/page [get]
func PageModels(ctx context.Context, req *app.RequestContext) {
	reqCtx := &base.ReqContext{
		Ctx: ctx,
		Req: req,
	}
	model.GetModelList(reqCtx)
}
