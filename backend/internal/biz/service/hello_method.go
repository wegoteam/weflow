package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	hello "hello/hertz_gen/hello"
)

type HelloMethodService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHelloMethodService(Context context.Context, RequestContext *app.RequestContext) *HelloMethodService {
	return &HelloMethodService{RequestContext: RequestContext, Context: Context}
}

func (h *HelloMethodService) Run(req *hello.HelloReq) (resp *hello.HelloResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
