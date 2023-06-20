package base

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Response
// @Description: 响应体
type Response struct {
	Code int         `json:"code"` // 0:成功，其他：失败
	Data interface{} `json:"data"` // 数据
	Msg  string      `json:"msg"`  // 错误信息
}

// ReqContext
// @Description: 基础请求上下文
type ReqContext struct {
	Ctx context.Context
	Req *app.RequestContext
}

// Error
// @Description: 响应错误
// @receiver: base
// @param: code
// @param: err
func (base *ReqContext) Error(code int, err error) {
	var res = &Response{
		Code: code,
		Data: nil,
		Msg:  err.Error(),
	}
	base.Req.JSON(consts.StatusOK, res)
}

// ErrorData
// @Description: 响应错误
// @receiver base
// @param: code
// @param: data
// @param: err
func (base *ReqContext) ErrorData(code int, data interface{}, err string) {
	var res = &Response{
		Code: code,
		Data: data,
		Msg:  err,
	}
	base.Req.JSON(consts.StatusOK, res)
}

// OK
// @Description: 响应成功
// @receiver: base
// @param: data
func (base *ReqContext) OK(data interface{}) {
	var res = &Response{
		Code: 0,
		Data: data,
		Msg:  "",
	}
	base.Req.JSON(consts.StatusOK, res)
}

// OkMsg
// @Description: 响应成功
// @receiver: base
// @param: data
func (base *ReqContext) OkMsg(data interface{}, err string) {
	var res = &Response{
		Code: 0,
		Data: data,
		Msg:  err,
	}
	base.Req.JSON(consts.StatusOK, res)
}
