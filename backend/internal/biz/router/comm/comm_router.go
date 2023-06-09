package Comm

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	commService "github.com/wegoteam/weflow/internal/biz/handler/comm"
)

// Register
// @Description: 注册公共路由
// @param: h
func Register(h *server.Hertz) {
	commGroup := h.Group("/comm")
	commGroup.GET("/snowflake", GetSnowflake)
}

// GetSnowflake 获取雪花算法唯一ID
// @Summary 获取雪花算法唯一ID
// @Tags 公共
// @ID GetSnowflake
// @Description 获取雪花算法唯一ID
// @Accept application/json
// @Produce application/json
// @Success 200 {object} base.Response{data=string}
// @Router /comm/snowflake [get]
func GetSnowflake(ctx context.Context, reqCtx *app.RequestContext) {
	res := commService.GetSnowflake()
	reqCtx.JSON(consts.StatusOK, res)
}
