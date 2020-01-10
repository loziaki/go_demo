package controller

import (
	"github.com/kataras/iris/v12"
)

var defaultApp *iris.Application

func Regist(app *iris.Application) {
	app.Get("/ping", func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"message": "pong",
		})
	})

	registUserApi(app)
	//测试wx api
	registWxApi(app)

}

func newResp(code int, msg string, data interface{}) *iris.Map {
	return &iris.Map{
		"status": code,
		"msg":    msg,
		"res":    data,
	}
}
