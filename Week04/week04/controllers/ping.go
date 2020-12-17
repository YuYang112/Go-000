package controllers

import (
	"github.com/kataras/iris/v12"
)

//测试服务运行状态
func TestPing(ctx iris.Context) {
	rep := SuccessRep("dgyt")
	ctx.JSON(rep)
}
