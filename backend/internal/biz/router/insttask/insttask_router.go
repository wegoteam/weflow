package Insttask

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	hertzconsts "github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/wegoteam/weflow/internal/base"
	"github.com/wegoteam/weflow/internal/biz/entity/bo"
	"github.com/wegoteam/weflow/internal/biz/entity/vo"
	insttaskService "github.com/wegoteam/weflow/internal/biz/handler/insttask"
	"github.com/wegoteam/weflow/internal/consts"
	"github.com/wegoteam/weflow/pkg/common/entity"
)

// Register
// @Description: 注册实例任务路由
// @param: h
func Register(h *server.Hertz) {
	insttaskGroup := h.Group("/insttask")
	insttaskGroup.POST("/initiated", GetInitiateInstTaskList)
	insttaskGroup.POST("/start", StartInstTask)
	insttaskGroup.POST("/stop", StopInstTask)
	insttaskGroup.POST("/suspend", SuspendInstTask)
	insttaskGroup.POST("/resume", ResumeInstTask)
	insttaskGroup.POST("/del", DeleteInstTask)
	insttaskGroup.GET("/model/detail", GetInsttaskModelDetail)
	insttaskGroup.GET("/detail", GetInsttaskAllDetail)
}

// GetInitiateInstTaskList 获取发起的实例任务列表（已发起）
// @Summary 获取发起的实例任务列表（已发起）
// @Tags 实例任务
// @Param InstTaskQueryVO body vo.InstTaskQueryVO true "已发起的请求参数"
// @Description 获取发起的实例任务列表（已发起）
// @Accept application/json
// @Produce application/json
// @Success 200 {object} base.Response{data=bo.InstTaskResult} "返回结果"
// @Router /insttask/initiated [post]
func GetInitiateInstTaskList(ctx context.Context, reqCtx *app.RequestContext) {
	var req vo.InstTaskQueryVO
	reqCtx.Bind(&req)
	if req.PageNum == 0 {
		req.PageNum = consts.DefaultPageNum
	}
	if req.PageSize == 0 {
		req.PageSize = consts.DefaultPageSize
	}
	param := &entity.InstTaskQueryBO{
		UserID:          consts.UserID,
		PageSize:        req.PageSize,
		PageNum:         req.PageNum,
		TaskName:        req.TaskName,
		InstStatus:      req.InstStatus,
		ModelID:         req.ModelID,
		CreateTimeStart: req.CreateTimeStart,
		CreateTimeEnd:   req.CreateTimeEnd,
		FinishTimeStart: req.FinishTimeStart,
		FinishTimeEnd:   req.FinishTimeEnd,
	}
	res := insttaskService.GetInitiateInstTaskList(param)
	reqCtx.JSON(hertzconsts.StatusOK, res)
}

// StartInstTask 发起实例任务
// @Summary 发起实例任务
// @Tags 实例任务
// @Description 发起实例任务
// @Param InstTaskStartVO body vo.InstTaskStartVO true "发起实例任务的请求参数"
// @Accept application/json
// @Produce application/json
// @Success 200 {object} base.Response{} "返回结果"
// @Router /insttask/start [post]
func StartInstTask(ctx context.Context, reqCtx *app.RequestContext) {
	var req vo.InstTaskStartVO
	bindErr := reqCtx.BindAndValidate(&req)
	if bindErr != nil {
		res := base.Fail(consts.ERROR, bindErr.Error())
		reqCtx.JSON(hertzconsts.StatusOK, res)
		return
	}
	param := &bo.InstTaskStartBO{
		UserID:   consts.UserID,
		UserName: consts.UserName,
		ModelID:  req.ModelID,
		Params:   req.Params,
	}
	res := insttaskService.Start(param)
	reqCtx.JSON(hertzconsts.StatusOK, res)
}

// StopInstTask 终止实例任务
// @Summary 终止实例任务
// @Tags 实例任务
// @Description 终止实例任务
// @Param InstTaskStopVO body vo.InstTaskStopVO true "终止实例任务的请求参数"
// @Accept application/json
// @Produce application/json
// @Success 200 {object} base.Response{} "返回结果"
// @Router /insttask/stop [post]
func StopInstTask(ctx context.Context, reqCtx *app.RequestContext) {
	var req vo.InstTaskStopVO
	bindErr := reqCtx.BindAndValidate(&req)
	if bindErr != nil {
		res := base.Fail(consts.ERROR, bindErr.Error())
		reqCtx.JSON(hertzconsts.StatusOK, res)
		return
	}
	param := &bo.InstTaskStopBO{
		OpUserID:    consts.UserID,
		OpUserName:  consts.UserName,
		InstTaskID:  req.InstTaskID,
		OpinionDesc: req.OpinionDesc,
	}
	res := insttaskService.Stop(param)
	reqCtx.JSON(hertzconsts.StatusOK, res)
}

// SuspendInstTask 挂起实例任务
// @Summary 挂起实例任务
// @Tags 实例任务
// @Description 挂起实例任务
// @Accept application/json
// @Param InstTaskSuspendVO body vo.InstTaskSuspendVO true "请求参数"
// @Produce application/json
// @Success 200 {object} base.Response{} "返回结果"
// @Router /insttask/suspend [post]
func SuspendInstTask(ctx context.Context, reqCtx *app.RequestContext) {
	var req vo.InstTaskSuspendVO
	bindErr := reqCtx.BindAndValidate(&req)
	if bindErr != nil {
		res := base.Fail(consts.ERROR, bindErr.Error())
		reqCtx.JSON(hertzconsts.StatusOK, res)
		return
	}
	param := &bo.InstTaskSuspendBO{
		OpUserID:    consts.UserID,
		OpUserName:  consts.UserName,
		InstTaskID:  req.InstTaskID,
		OpinionDesc: req.OpinionDesc,
	}
	res := insttaskService.Suspend(param)
	reqCtx.JSON(hertzconsts.StatusOK, res)
}

// ResumeInstTask 恢复实例任务
// @Summary 恢复实例任务
// @Tags 实例任务
// @Description 恢复实例任务
// @Param InstTaskSesumeVO body vo.InstTaskSesumeVO true "请求参数"
// @Accept application/json
// @Produce application/json
// @Success 200 {object} base.Response{} "返回结果"
// @Router /insttask/resume [post]
func ResumeInstTask(ctx context.Context, reqCtx *app.RequestContext) {
	var req vo.InstTaskSesumeVO
	bindErr := reqCtx.BindAndValidate(&req)
	if bindErr != nil {
		res := base.Fail(consts.ERROR, bindErr.Error())
		reqCtx.JSON(hertzconsts.StatusOK, res)
		return
	}
	param := &bo.InstTaskSesumeBO{
		OpUserID:    consts.UserID,
		OpUserName:  consts.UserName,
		InstTaskID:  req.InstTaskID,
		OpinionDesc: req.OpinionDesc,
	}
	res := insttaskService.Sesume(param)
	reqCtx.JSON(hertzconsts.StatusOK, res)
}

// DeleteInstTask 删除实例任务
// @Summary 删除实例任务
// @Tags 实例任务
// @Description 删除实例任务
// @Param InstTaskDeleteVO body vo.InstTaskDeleteVO true "请求参数"
// @Accept application/json
// @Produce application/json
// @Success 200 {object} base.Response{} "返回结果"
// @Router /insttask/del [post]
func DeleteInstTask(ctx context.Context, reqCtx *app.RequestContext) {
	var req vo.InstTaskDeleteVO
	bindErr := reqCtx.BindAndValidate(&req)
	if bindErr != nil {
		res := base.Fail(consts.ERROR, bindErr.Error())
		reqCtx.JSON(hertzconsts.StatusOK, res)
		return
	}
	param := &bo.InstTaskDeleteBO{
		OpUserID:    consts.UserID,
		OpUserName:  consts.UserName,
		InstTaskID:  req.InstTaskID,
		OpinionDesc: req.OpinionDesc,
	}
	res := insttaskService.Delete(param)
	reqCtx.JSON(hertzconsts.StatusOK, res)
}

// GetInsttaskModelDetail 查询实例任务模板详情(模板信息，版本，流程定义，流程定义详情)
// @Summary 查询实例任务模板详情(模板信息，版本，流程定义，流程定义详情)
// @Tags 实例任务
// @Description 查询实例任务模板详情(模板信息，版本，流程定义，流程定义详情)
// @Accept application/json
// @Param InsttaskModelDetailQueryVO query vo.InsttaskModelDetailQueryVO true "请求参数"
// @Produce application/json
// @Success 200 {object} base.Response{data=bo.ModelDetailResult} "返回结果"
// @Router /insttask/model/detail [get]
func GetInsttaskModelDetail(ctx context.Context, reqCtx *app.RequestContext) {
	var req vo.InsttaskModelDetailQueryVO
	reqCtx.Bind(&req)
	res := insttaskService.GetInsttaskModelDetail(req.InstTaskID)
	reqCtx.JSON(hertzconsts.StatusOK, res)
}

// GetInsttaskAllDetail 查询实例任务详情（实例任务，节点任务，用户任务信息）
// @Summary 查询实例任务详情（实例任务，节点任务，用户任务信息）
// @Tags 实例任务
// @Description 查询实例任务详情（实例任务，节点任务，用户任务信息）
// @Accept application/json
// @Param InsttaskDetailQueryVO query vo.InsttaskDetailQueryVO true "请求参数"
// @Produce application/json
// @Success 200 {object} base.Response{data=bo.ModelDetailResult} "返回结果"
// @Router /insttask/detail [get]
func GetInsttaskAllDetail(ctx context.Context, reqCtx *app.RequestContext) {
	var req vo.InsttaskDetailQueryVO
	reqCtx.Bind(&req)
	res := insttaskService.GetInsttaskAllDetail(req.InstTaskID, req.UserTaskID)
	reqCtx.JSON(hertzconsts.StatusOK, res)
}
