package main

import (
	"github.com/kataras/iris/v12"

	"loziaki/go_demo/controller"
)

func newApp() *iris.Application {
	app := iris.New()

	controller.Regist(app)

	// app.Get("/shutdown", func(ctx iris.Context) {
	// 	app.Logger().Debug("ready to shutdown")
	// 	if err := app.Shutdown(ctx); err != nil {
	// 		app.Logger().Debug("lalalala")
	// 	} else {
	// 		app.Logger().Debug("shutdown")
	// 	}
	// 	app.Logger().Debug("shutdown done")
	// })

	return app
}

func main() {
	app := newApp()
	app.Logger().SetLevel("debug")
	app.Run(iris.Addr(":8233"))
}
