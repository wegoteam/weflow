package service

import (
	"context"
	common "github.com/wegoteam/weflow/api/kitex_gen/common"
)

type GetPersonService struct {
	ctx context.Context
} // NewGetPersonService new GetPersonService
func NewGetPersonService(ctx context.Context) *GetPersonService {
	return &GetPersonService{ctx: ctx}
}

// Run create note info
func (s *GetPersonService) Run(name string) (resp *common.Response, err error) {
	// Finish your business logic.

	return
}
