package Usertask

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	hertzconsts "github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/wegoteam/weflow/internal/biz/entity/vo"
	usertaskService "github.com/wegoteam/weflow/internal/biz/handler/usertask"
	"github.com/wegoteam/weflow/internal/consts"
	"github.com/wegoteam/weflow/pkg/common/entity"
)

// Register
// @Description: 注册用户任务路由
// @param: h
func Register(h *server.Hertz) {
	usertaskGroup := h.Group("/usertask")
	usertaskGroup.GET("/todo", GetTodoUserTaskList)
	usertaskGroup.GET("/done", GetDoneUserTaskList)
	usertaskGroup.GET("/received", GetReceivedUserTaskList)
	usertaskGroup.POST("/agree", AgreeUserTask)
	usertaskGroup.POST("/disagree", DisagreeUserTask)
	usertaskGroup.POST("/save", SaveUserTask)
}

// GetTodoUserTaskList 获取待办用户任务列表
// @Summary 获取待办用户任务列表（待处理）
// @Tags 用户任务
// @Param UserTaskQueryVO query vo.UserTaskQueryVO true "待处理的请求参数"
// @Description 获取待办用户任务列表（待处理）
// @Accept application/json
// @Produce application/json
// @Success 200 {object} base.Response{data=bo.UserTaskTodoResult} "返回结果"
// @Router /usertask/todo [get]
func GetTodoUserTaskList(ctx context.Context, reqCtx *app.RequestContext) {
	// 获取请求参数
	var req vo.UserTaskQueryVO
	reqCtx.Bind(&req)
	param := &entity.UserTaskQueryBO{
		UserID:          consts.UserID,
		PageSize:        req.PageSize,
		PageNum:         req.PageNum,
		TaskName:        req.TaskName,
		InstStatus:      req.InstStatus,
		ModelId:         req.ModelId,
		CreateUserId:    req.CreateUserId,
		CreateTimeStart: req.CreateTimeStart,
		CreateTimeEnd:   req.CreateTimeEnd,
		FinishTimeStart: req.FinishTimeStart,
		FinishTimeEnd:   req.FinishTimeEnd,
	}
	res := usertaskService.GetTodoUserTaskList(param)
	reqCtx.JSON(hertzconsts.StatusOK, res)
}

// GetDoneUserTaskList 获取已办用户任务列表（已处理）
// @Summary 获取已办用户任务列表（已处理）
// @Tags 用户任务
// @Description 获取已办用户任务列表（已处理）
// @Param UserTaskQueryVO query vo.UserTaskQueryVO true "已处理的请求参数"
// @Accept application/json
// @Produce application/json
// @Success 200 {object} base.Response{data=bo.UserTaskResult} "返回结果"
// @Router /usertask/done [get]
func GetDoneUserTaskList(ctx context.Context, reqCtx *app.RequestContext) {
	// 获取请求参数
	var req vo.UserTaskQueryVO
	reqCtx.Bind(&req)
	param := &entity.UserTaskQueryBO{
		UserID:          consts.UserID,
		PageSize:        req.PageSize,
		PageNum:         req.PageNum,
		TaskName:        req.TaskName,
		InstStatus:      req.InstStatus,
		ModelId:         req.ModelId,
		CreateUserId:    req.CreateUserId,
		CreateTimeStart: req.CreateTimeStart,
		CreateTimeEnd:   req.CreateTimeEnd,
		FinishTimeStart: req.FinishTimeStart,
		FinishTimeEnd:   req.FinishTimeEnd,
	}
	res := usertaskService.GetDoneUserTaskList(param)
	reqCtx.JSON(hertzconsts.StatusOK, res)
}

// GetReceivedUserTaskList 获取用户任务列表（我收到的）
// @Summary 获取用户任务列表（我收到的）
// @Tags 用户任务
// @Description 获取用户任务列表（我收到的）
// @Param UserTaskQueryVO query vo.UserTaskQueryVO true "我收到的的请求参数"
// @Accept application/json
// @Produce application/json
// @Success 200 {object} base.Response{data=bo.UserTaskResult} "返回结果"
// @Router /usertask/received [get]
func GetReceivedUserTaskList(ctx context.Context, reqCtx *app.RequestContext) {
	// 获取请求参数
	var req vo.UserTaskQueryVO
	reqCtx.Bind(&req)
	param := &entity.UserTaskQueryBO{
		UserID:          consts.UserID,
		PageSize:        req.PageSize,
		PageNum:         req.PageNum,
		TaskName:        req.TaskName,
		InstStatus:      req.InstStatus,
		ModelId:         req.ModelId,
		CreateUserId:    req.CreateUserId,
		CreateTimeStart: req.CreateTimeStart,
		CreateTimeEnd:   req.CreateTimeEnd,
		FinishTimeStart: req.FinishTimeStart,
		FinishTimeEnd:   req.FinishTimeEnd,
	}
	res := usertaskService.GetReceivedUserTaskList(param)
	reqCtx.JSON(hertzconsts.StatusOK, res)
}

// AgreeUserTask 同意用户任务
// @Summary 同意用户任务
// @Tags 用户任务
// @Description 同意用户任务
// @Accept application/json
// @Produce application/json
// @Success 200 {object} base.Response{} "返回结果"
// @Router /usertask/agree [post]
func AgreeUserTask(ctx context.Context, reqCtx *app.RequestContext) {

}

// DisagreeUserTask 不同意用户任务
// @Summary 不同意用户任务
// @Tags 用户任务
// @Description 不同意用户任务
// @Accept application/json
// @Produce application/json
// @Success 200 {object} base.Response{} "返回结果"
// @Router /usertask/disagree [post]
func DisagreeUserTask(ctx context.Context, reqCtx *app.RequestContext) {

}

// SaveUserTask 保存用户任务
// @Summary 保存用户任务
// @Tags 用户任务
// @Description 保存用户任务
// @Accept application/json
// @Produce application/json
// @Success 200 {object} base.Response{} "返回结果"
// @Router /usertask/save [post]
func SaveUserTask(ctx context.Context, reqCtx *app.RequestContext) {

}
