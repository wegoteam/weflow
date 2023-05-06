package service

import (
	"context"
	"testing"
)

func TestGetMap_Run(t *testing.T) {
	ctx := context.Background()
	s := NewGetMapService(ctx)
	// init req and assert value

	key := &string{}
	resp, err := s.Run(key)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if resp == nil {
		t.Errorf("unexpected nil response")
	}
	// todo: edit your unit test

}
