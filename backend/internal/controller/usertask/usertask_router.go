package Usertask

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/wegoteam/weflow/internal/base"
	"github.com/wegoteam/weflow/internal/entity/vo"
	usertaskService "github.com/wegoteam/weflow/internal/service/usertask"
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
}

// GetTodoUserTaskList 获取待办用户任务列表
// @Summary 获取待办用户任务列表（待处理）
// @Tags 用户任务
// @Param UserTaskQueryVO query vo.UserTaskQueryVO true "待处理的请求参数"
// @Description 获取待办用户任务列表（待处理）
// @Accept application/json
// @Produce application/json
// @Success 200 {object} base.Response{data=base.Page[bo.UserTaskTodoResult]{records=[]bo.UserTaskTodoResult}}
// @Router /usertask/todo [get]
func GetTodoUserTaskList(c context.Context, rc *app.RequestContext) {
	// 获取请求参数
	var req vo.UserTaskQueryVO
	rc.Bind(&req)
	param := &entity.UserTaskQueryBO{
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
	userTaskList := usertaskService.GetTodoUserTaskList(param)
	res := &base.Response{
		Data: userTaskList,
	}
	rc.JSON(consts.StatusOK, res)
}

// GetDoneUserTaskList 获取已办用户任务列表（已处理）
// @Summary 获取已办用户任务列表（已处理）
// @Tags 用户任务
// @Description 获取已办用户任务列表（已处理）
// @Param UserTaskQueryVO query vo.UserTaskQueryVO true "已处理的请求参数"
// @Accept application/json
// @Produce application/json
// @Success 200 {object} base.Response{data=base.Page[bo.UserTaskResult]{records=[]bo.UserTaskResult}}
// @Router /usertask/done [get]
func GetDoneUserTaskList(c context.Context, rc *app.RequestContext) {
	// 获取请求参数
	var req vo.UserTaskQueryVO
	rc.Bind(&req)
	param := &entity.UserTaskQueryBO{
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
	userTaskList := usertaskService.GetDoneUserTaskList(param)
	res := &base.Response{
		Data: userTaskList,
	}
	rc.JSON(consts.StatusOK, res)
}

// GetReceivedUserTaskList 获取用户任务列表（我收到的）
// @Summary 获取用户任务列表（我收到的）
// @Tags 用户任务
// @Description 获取用户任务列表（我收到的）
// @Param UserTaskQueryVO query vo.UserTaskQueryVO true "我收到的的请求参数"
// @Accept application/json
// @Produce application/json
// @Success 200 {object} base.Response{data=base.Page[bo.UserTaskResult]{records=[]bo.UserTaskResult}}
// @Router /usertask/received [get]
func GetReceivedUserTaskList(ctx context.Context, rc *app.RequestContext) {
	// 获取请求参数
	var req vo.UserTaskQueryVO
	rc.Bind(&req)
	param := &entity.UserTaskQueryBO{
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
	userTaskList := usertaskService.GetReceivedUserTaskList(param)
	res := &base.Response{
		Data: userTaskList,
	}
	rc.JSON(consts.StatusOK, res)
}
