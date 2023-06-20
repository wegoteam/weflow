package Insttask

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/wegoteam/weflow/internal/base"
	"github.com/wegoteam/weflow/internal/biz/entity/vo"
	"github.com/wegoteam/weflow/internal/biz/handler/insttask"
	"github.com/wegoteam/weflow/pkg/common/entity"
)

// Register
// @Description: 注册实例任务路由
// @param: h
func Register(h *server.Hertz) {
	insttaskGroup := h.Group("/insttask")
	insttaskGroup.GET("/initiated", GetInitiateInstTaskList)
}

// GetInitiateInstTaskList 获取发起的实例任务列表（已发起）
// @Summary 获取发起的实例任务列表（已发起）
// @Tags 实例任务
// @Param InstTaskQueryVO query vo.InstTaskQueryVO true "已发起的请求参数"
// @Description 获取发起的实例任务列表（已发起）
// @Accept application/json
// @Produce application/json
// @Success 200 {object} base.Response{data=base.Page[bo.InstTaskResult]{records=[]bo.InstTaskResult}}
// @Router /insttask/initiated [get]
func GetInitiateInstTaskList(ctx context.Context, rc *app.RequestContext) {
	var req vo.InstTaskQueryVO
	rc.Bind(&req)
	param := &entity.InstTaskQueryBO{
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
	taskList := insttask.GetInitiateInstTaskList(param)
	res := &base.Response{
		Data: taskList,
	}
	rc.JSON(consts.StatusOK, res)
}
