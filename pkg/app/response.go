package app

import (
	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code bool         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (g *Gin) Response(httpCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: true,
		Msg:  "null",
		Data: data,
	})
}

func (g *Gin) ErrorResponse(httpCode int, msg string) {
	g.C.JSON(httpCode, Response{
		Code: false,
		Msg:  msg,
		Data: nil,
	})
}
