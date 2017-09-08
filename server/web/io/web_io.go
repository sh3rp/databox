package io

import (
	"github.com/sh3rp/databox/common"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(data interface{}) *Response {
	return &Response{
		Code:    common.SUCCESS,
		Message: common.SUCCESS_MSG,
		Data:    data,
	}
}

func Error(code int, err error) *Response {
	return &Response{
		Code:    code,
		Message: err.Error(),
	}
}
