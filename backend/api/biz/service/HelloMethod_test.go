package service

import (
	"context"
	api "github.com/wego2023/weflow/api/kitex_gen/api"
	"testing"
)

func TestHelloMethod_Run(t *testing.T) {
	ctx := context.Background()
	s := NewHelloMethodService(ctx)
	// init req and assert value

	request := &api.ApiReq{}
	resp, err := s.Run(request)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if resp == nil {
		t.Errorf("unexpected nil response")
	}
	// todo: edit your unit test

}
