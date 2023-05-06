package service

import (
	"context"
)

type GetMapService struct {
	ctx context.Context
} // NewGetMapService new GetMapService
func NewGetMapService(ctx context.Context) *GetMapService {
	return &GetMapService{ctx: ctx}
}

// Run create note info
func (s *GetMapService) Run(key string) (resp map[string]int32, err error) {
	// Finish your business logic.

	return
}
