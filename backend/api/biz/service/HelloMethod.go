package service

import (
	"context"
	api "github.com/wego2023/weflow/api/kitex_gen/api"
)

type HelloMethodService struct {
	ctx context.Context
} // NewHelloMethodService new HelloMethodService
func NewHelloMethodService(ctx context.Context) *HelloMethodService {
	return &HelloMethodService{ctx: ctx}
}

// Run create note info
func (s *HelloMethodService) Run(request *api.ApiReq) (resp *api.ApiResp, err error) {
	// Finish your business logic.

	return
}
