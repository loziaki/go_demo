package controller

import (
	"github.com/kataras/iris/v12"

	"loziaki/go_demo/module/wxsdk"
)

func registWxApi(app *iris.Application) {
	app.Get("/wx_getAccessToken", func(ctx iris.Context) {
		res, _ := wxsdk.DefaultApp.GetLatestAccessToken()
		app.Logger().Debug(res)
	})

	app.Get("/wx_createQRCode", func(ctx iris.Context) {
		res, _ := wxsdk.DefaultApp.CreateQRCode("/index", 100).GetBytes()
		// ctx.Header("content-type:image/jpeg", "")
		ctx.Write(res)
	})

}
