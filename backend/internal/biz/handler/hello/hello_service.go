package hello

import (
	"context"
	"fmt"
	"github.com/bytedance/gopkg/util/logger"
	"github.com/pkg/errors"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/wego2023/weflow/internal/biz/service"
	"github.com/wego2023/weflow/internal/biz/utils"
	hello "github.com/wego2023/weflow/internal/hertz_gen/hello"
)

// HelloMethod .
// @router /hello [GET]
func HelloMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req hello.HelloReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewHelloMethodService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

func GlobalErrorHandler(ctx context.Context, c *app.RequestContext) {
	c.Next(ctx)

	if len(c.Errors) == 0 {
		// 没有收集到异常直接返回
		fmt.Println("retun")
		return
	}
	hertzErr := c.Errors[0]
	// 获取errors包装的err
	err := hertzErr.Unwrap()
	// 打印异常堆栈
	logger.CtxErrorf(ctx, "%+v", err)
	// 获取原始err
	err = errors.Unwrap(err)
	// todo 进行错误类型判断
	c.JSON(400, map[string]interface{}{
		"code":    400,
		"message": err.Error(),
	})
}
