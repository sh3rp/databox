package io

import (
	"github.com/gin-gonic/gin"
	"github.com/sh3rp/databox/common"
	"github.com/sh3rp/databox/msg"
)

type NewBoxRequest struct {
	Token       *msg.Token
	Name        string
	Description string
	Password    string
}

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

func Error(code int, err string) *Response {
	return &Response{
		Code:    code,
		Message: err,
	}
}

func Respond(c *gin.Context, errorCode int, data interface{}) {
	if errorCode != 0 {
		c.JSON(200, Error(errorCode, common.ERRORS[errorCode]))
	} else {
		c.JSON(200, Success(data))
	}
}
