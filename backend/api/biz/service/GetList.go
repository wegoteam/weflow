package service

import (
	"context"
	flow "github.com/wegoteam/weflow/api/kitex_gen/flow"
)

type GetListService struct {
	ctx context.Context
} // NewGetListService new GetListService
func NewGetListService(ctx context.Context) *GetListService {
	return &GetListService{ctx: ctx}
}

// Run create note info
func (s *GetListService) Run() (resp []*flow.Person, err error) {
	// Finish your business logic.

	return
}
