package router

import "github.com/cloudwego/hertz/pkg/app/server/binding"

type BindError struct {
	ErrType, FailField, Msg string
}

// Error implements error interface.
func (e *BindError) Error() string {
	if e.Msg == "" {
		return "参数绑定错误"
	}
	return e.Msg
}

type ValidateError struct {
	ErrType, FailField, Msg string
}

// Error implements error interface.
func (e *ValidateError) Error() string {
	if e.Msg == "" {
		return "参数验证错误"
	}
	return e.Msg
}

func init() {
	CustomBindErrFunc := func(failField, msg string) error {
		err := BindError{
			ErrType:   "bindErr",
			FailField: failField,
			Msg:       msg,
		}

		return &err
	}

	CustomValidateErrFunc := func(failField, msg string) error {
		err := ValidateError{
			ErrType:   "validateErr",
			FailField: failField,
			Msg:       msg,
		}

		return &err
	}

	binding.SetErrorFactory(CustomBindErrFunc, CustomValidateErrFunc)
}
