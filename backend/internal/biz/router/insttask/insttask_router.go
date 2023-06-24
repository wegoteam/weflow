package Insttask

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	hertzconsts "github.com/cloudwego/hertz/pkg/protocol/consts"
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
	insttaskGroup.GET("/initiated", GetInitiateInstTaskList)
	insttaskGroup.POST("/start", StartInstTask)
	insttaskGroup.POST("/stop", StopInstTask)
	insttaskGroup.POST("/suspend", SuspendInstTask)
	insttaskGroup.POST("/resume", ResumeInstTask)
	insttaskGroup.POST("/del", DeleteInstTask)
}

// GetInitiateInstTaskList 获取发起的实例任务列表（已发起）
// @Summary 获取发起的实例任务列表（已发起）
// @Tags 实例任务
// @Param InstTaskQueryVO query vo.InstTaskQueryVO true "已发起的请求参数"
// @Description 获取发起的实例任务列表（已发起）
// @Accept application/json
// @Produce application/json
// @Success 200 {object} base.Response{data=bo.InstTaskResult} "返回结果"
// @Router /insttask/initiated [get]
func GetInitiateInstTaskList(ctx context.Context, reqCtx *app.RequestContext) {
	var req vo.InstTaskQueryVO
	reqCtx.Bind(&req)
	param := &entity.InstTaskQueryBO{
		UserID:          consts.UserID,
		PageSize:        req.PageSize,
		PageNum:         req.PageNum,
		TaskName:        req.TaskName,
		InstStatus:      req.InstStatus,
		ModelId:         req.ModelId,
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
// @Accept application/json
// @Produce application/json
// @Success 200 {object} base.Response{} "返回结果"
// @Router /insttask/start [post]
func StartInstTask(ctx context.Context, reqCtx *app.RequestContext) {

}

// StopInstTask 终止实例任务
// @Summary 终止实例任务
// @Tags 实例任务
// @Description 终止实例任务
// @Accept application/json
// @Produce application/json
// @Success 200 {object} base.Response{} "返回结果"
// @Router /insttask/stop [post]
func StopInstTask(ctx context.Context, reqCtx *app.RequestContext) {

}

// SuspendInstTask 挂起实例任务
// @Summary 挂起实例任务
// @Tags 实例任务
// @Description 挂起实例任务
// @Accept application/json
// @Produce application/json
// @Success 200 {object} base.Response{} "返回结果"
// @Router /insttask/suspend [post]
func SuspendInstTask(ctx context.Context, reqCtx *app.RequestContext) {

}

// ResumeInstTask 恢复实例任务
// @Summary 恢复实例任务
// @Tags 实例任务
// @Description 恢复实例任务
// @Accept application/json
// @Produce application/json
// @Success 200 {object} base.Response{} "返回结果"
// @Router /insttask/resume [post]
func ResumeInstTask(ctx context.Context, reqCtx *app.RequestContext) {

}

// DeleteInstTask 删除实例任务
// @Summary 删除实例任务
// @Tags 实例任务
// @Description 删除实例任务
// @Accept application/json
// @Produce application/json
// @Success 200 {object} base.Response{} "返回结果"
// @Router /insttask/del [post]
func DeleteInstTask(ctx context.Context, reqCtx *app.RequestContext) {

}
