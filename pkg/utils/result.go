package utils

import (
	"fmt"
)

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewResult(code int, msg string, data interface{}) *Result {
	return &Result{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

func (e *Result) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Msg)
}

var _ error = (*Result)(nil)
