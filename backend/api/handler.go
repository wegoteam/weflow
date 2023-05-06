package main

import (
	"context"
	"github.com/wegoteam/weflow/api/biz/service"
	common "github.com/wegoteam/weflow/api/kitex_gen/common"
	flow "github.com/wegoteam/weflow/api/kitex_gen/flow"
)

// FlowServiceImpl implements the last service interface defined in the IDL.
type FlowServiceImpl struct{}

// GetList implements the FlowServiceImpl interface.
func (s *FlowServiceImpl) GetList(ctx context.Context) (resp []*flow.Person, err error) {
	resp, err = service.NewGetListService(ctx).Run()

	return resp, err
}

// GetMap implements the FlowServiceImpl interface.
func (s *FlowServiceImpl) GetMap(ctx context.Context, key string) (resp map[string]int32, err error) {
	resp, err = service.NewGetMapService(ctx).Run(key)

	return resp, err
}

// GetPerson implements the FlowServiceImpl interface.
func (s *FlowServiceImpl) GetPerson(ctx context.Context, name string) (resp *common.Response, err error) {
	resp, err = service.NewGetPersonService(ctx).Run(name)

	return resp, err
}
