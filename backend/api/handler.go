package main

import (
	"context"
	"github.com/wego2023/weflow/api/biz/service"
	api "github.com/wego2023/weflow/api/kitex_gen/api"
)

// ApiServiceImpl implements the last service interface defined in the IDL.
type ApiServiceImpl struct{}

// HelloMethod implements the ApiServiceImpl interface.
func (s *ApiServiceImpl) HelloMethod(ctx context.Context, request *api.ApiReq) (resp *api.ApiResp, err error) {
	resp, err = service.NewHelloMethodService(ctx).Run(request)

	return resp, err
}
