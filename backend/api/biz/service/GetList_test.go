package service

import (
	"context"
	flow "github.com/wegoteam/weflow/api/kitex_gen/flow"
	"testing"
)

func TestGetList_Run(t *testing.T) {
	ctx := context.Background()
	s := NewGetListService(ctx)
	// init req and assert value
	resp, err := s.Run()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if resp == nil {
		t.Errorf("unexpected nil response")
	}
	// todo: edit your unit test

}
