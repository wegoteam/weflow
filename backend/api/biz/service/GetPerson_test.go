package service

import (
	"context"
	common "github.com/wegoteam/weflow/api/kitex_gen/common"
	"testing"
)

func TestGetPerson_Run(t *testing.T) {
	ctx := context.Background()
	s := NewGetPersonService(ctx)
	// init req and assert value

	name := &string{}
	resp, err := s.Run(name)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if resp == nil {
		t.Errorf("unexpected nil response")
	}
	// todo: edit your unit test

}
