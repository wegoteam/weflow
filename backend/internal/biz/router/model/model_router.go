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
	modelGroup.POST("/list", GetModelList)
	modelGroup.POST("/page", PageModelList)
	modelGroup.POST("/save", SaveModel)
	modelGroup.POST("/publish", PublishModel)
	modelGroup.POST("/invalid", InvalidModel)
	modelGroup.POST("/version/release", ReleaseModelVersion)
	modelGroup.GET("/version/get", GetModelVersionList)
	modelGroup.GET("/group/list", GetModelGroups)
	modelGroup.POST("/group/add", AddModelGroup)
	modelGroup.POST("/group/edit", EditModelGroup)
	modelGroup.POST("/group/del", DelModelGroup)
	modelGroup.POST("/details", GetGroupModelDetails)
}

// GetGroupModelDetails 获取所有组的所有模版
// @Summary 获取所有组的所有模版
// @Tags 模板
// @Description 获取所有组的所有模版
// @Accept application/json
// @Param GroupModelQueryVO body vo.GroupModelQueryVO true "请求参数"
// @Produce application/json
// @Success 200 {object} base.Response{data=bo.ModelDetailResult} "返回结果"
// @Router /model/details [post]
func GetGroupModelDetails(ctx context.Context, reqCtx *app.RequestContext) {
	var req vo.GroupModelQueryVO
	reqCtx.Bind(&req)
	param := &entity.GroupModelQueryBO{
		ModelName: req.ModelName,
	}
	res := modelService.GetGroupModelDetails(param)
	reqCtx.JSON(hertzconsts.StatusOK, res)
}

// GetModelList 获取模板列表
// @Summary 获取模板列表
// @Tags 模板
// @Description 获取模板列表
// @Accept application/json
// @Param ModelQueryVO body vo.ModelQueryVO true "请求参数"
// @Produce application/json
// @Success 200 {object} base.Response{data=bo.ModelDetailResult} "返回结果"
// @Router /model/list [post]
func GetModelList(ctx context.Context, reqCtx *app.RequestContext) {
	var req vo.ModelQueryVO
	reqCtx.Bind(&req)
	param := &entity.ModelQueryBO{
		ModelName: req.ModelName,
		Status:    req.Status,
	}
	res := modelService.GetModelList(param)
	reqCtx.JSON(hertzconsts.StatusOK, res)
}

// PageModelList 分页获取模板列表
// @Summary 分页获取模板列表
// @Tags 模板
// @Description 分页获取模板列表
// @Accept application/json
// @Param ModelPageVO body vo.ModelPageVO true "请求参数"
// @Produce application/json
// @Success 200 {object} base.Response{data=bo.ModelDetailResult} "返回结果"
// @Router /model/page [post]
func PageModelList(ctx context.Context, reqCtx *app.RequestContext) {
	var req vo.ModelPageVO
	reqCtx.Bind(&req)
	param := &entity.ModelPageBO{
		ModelName: req.ModelName,
		Status:    req.Status,
		PageNum:   req.PageNum,
		PageSize:  req.PageSize,
	}
	res := modelService.PageModelList(param)
	reqCtx.JSON(hertzconsts.StatusOK, res)
}

// SaveModel 保存模板
// @Summary 保存模板
// @Tags 模板
// @Description 保存模板
// @Accept application/json
// @Param ModelSaveVO body vo.ModelSaveVO true "请求参数"
// @Produce application/json
// @Success 200 {object} base.Response{} "返回结果"
// @Router /model/save [post]
func SaveModel(ctx context.Context, reqCtx *app.RequestContext) {
	var req vo.ModelSaveVO
	reqCtx.Bind(&req)
	var param = &entity.ModelSaveBO{
		ModelID: req.ModelID,
		Base: entity.ModelBaseSetup{
			ModelName: req.Base.ModelName,
			Remark:    req.Base.Remark,
			GroupID:   req.Base.GroupID,
			IconURL:   req.Base.IconURL,
		},
		FlowContent: req.FlowContent,
		FormContent: req.FormContent,
		Advanced: entity.ModelAdvancedSetup{
			TitleType:    req.Advanced.TitleType,
			TitleContent: req.Advanced.TitleContent,
		},
	}
	res := modelService.SaveModel(param)
	reqCtx.JSON(hertzconsts.StatusOK, res)
}

// PublishModel 发布模板
// @Summary 发布模板
// @Tags 模板
// @Description 发布模板
// @Accept application/json
// @Param ModelSaveVO body vo.ModelSaveVO true "请求参数"
// @Produce application/json
// @Success 200 {object} base.Response{} "返回结果"
// @Router /model/publish [post]
func PublishModel(ctx context.Context, reqCtx *app.RequestContext) {
	var req vo.ModelSaveVO
	reqCtx.Bind(&req)
	var param = &entity.ModelSaveBO{
		ModelID: req.ModelID,
		Base: entity.ModelBaseSetup{
			ModelName: req.Base.ModelName,
			Remark:    req.Base.Remark,
			GroupID:   req.Base.GroupID,
			IconURL:   req.Base.IconURL,
		},
		FlowContent: req.FlowContent,
		FormContent: req.FormContent,
		Advanced: entity.ModelAdvancedSetup{
			TitleType:    req.Advanced.TitleType,
			TitleContent: req.Advanced.TitleContent,
		},
	}
	res := modelService.PublishModel(param)
	reqCtx.JSON(hertzconsts.StatusOK, res)
}

// InvalidModel 停用模板
// @Summary 停用模板
// @Tags 模板
// @Description 停用模板
// @Accept application/json
// @Param ModelInvalidVO body vo.ModelInvalidVO true "请求参数"
// @Produce application/json
// @Success 200 {object} base.Response{} "返回结果"
// @Router /model/invalid [post]
func InvalidModel(ctx context.Context, reqCtx *app.RequestContext) {
	var req vo.ModelInvalidVO
	reqCtx.Bind(&req)
	res := modelService.InvalidModel(req.ModelID)
	reqCtx.JSON(hertzconsts.StatusOK, res)
}

// ReleaseModelVersion 上线模板版本
// @Summary 上线模板版本
// @Tags 模板
// @Description 上线模板版本(模板版本列表启用版本)
// @Accept application/json
// @Param ReleaseModelVersionVO body vo.ReleaseModelVersionVO true "请求参数"
// @Produce application/json
// @Success 200 {object} base.Response{} "返回结果"
// @Router /model/version/release [post]
func ReleaseModelVersion(ctx context.Context, reqCtx *app.RequestContext) {
	var req vo.ReleaseModelVersionVO
	reqCtx.Bind(&req)
	res := modelService.ReleaseModelVersion(req.VersionID)
	reqCtx.JSON(hertzconsts.StatusOK, res)
}

// GetModelVersionList 根据模板查询获取模板版本列表
// @Summary 根据模板查询获取模板版本列表
// @Tags 模板
// @Description 根据模板查询获取模板版本列表
// @Param ModelVersionQueryVO query vo.ModelVersionQueryVO true "请求参数"
// @Accept application/json
// @Produce application/json
// @Success 200 {object} base.Response{data=bo.ModelVersionResult} "返回结果"
// @Router /model/version/get [get]
func GetModelVersionList(ctx context.Context, reqCtx *app.RequestContext) {
	var req vo.ModelVersionQueryVO
	reqCtx.Bind(&req)
	res := modelService.GetModelVersionList(req.ModelID)
	reqCtx.JSON(hertzconsts.StatusOK, res)
}

// GetModelGroups 查询获取模板组列表
// @Summary 查询获取模板组列表
// @Tags 模板
// @Description 查询获取模板组列表
// @Accept application/json
// @Produce application/json
// @Success 200 {object} base.Response{data=bo.ModelGroupResult} "返回结果"
// @Router /model/group/list [get]
func GetModelGroups(ctx context.Context, reqCtx *app.RequestContext) {
	res := modelService.GetModelGroupList()
	reqCtx.JSON(hertzconsts.StatusOK, res)
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
func AddModelGroup(ctx context.Context, reqCtx *app.RequestContext) {
	var req vo.ModelGroupAddVO
	reqCtx.Bind(&req)
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
	reqCtx.JSON(hertzconsts.StatusOK, res)
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
func EditModelGroup(ctx context.Context, reqCtx *app.RequestContext) {
	var req vo.ModelGroupEditVO
	reqCtx.Bind(&req)
	now := time.Now()
	param := &entity.ModelGroupEditBO{
		GroupID:    req.GroupID,
		GroupName:  req.GroupName,
		Remark:     req.Remark,
		UpdateUser: consts.UserID,
		UpdateTime: now,
	}
	res := modelService.EditModelGroup(param)
	reqCtx.JSON(hertzconsts.StatusOK, res)
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
func DelModelGroup(ctx context.Context, reqCtx *app.RequestContext) {
	var req vo.ModelGroupDelVO
	reqCtx.Bind(&req)
	param := &entity.ModelGroupDelBO{
		GroupID: req.GroupID,
	}
	res := modelService.DelModelGroup(param)
	reqCtx.JSON(hertzconsts.StatusOK, res)
}
