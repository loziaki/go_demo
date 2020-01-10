package controller

import (
	"github.com/kataras/iris/v12"
	"loziaki/go_demo/model"
	"loziaki/go_demo/module/wxsdk"
)

func registUserApi(app *iris.Application) {
	app.Get("/lalala", lalala)

	app.Get("/login", loginViaWx)
}

func lalala(ctx iris.Context) {
	ctx.JSON(iris.Map{
		"message": "here is abc",
	})
}

func loginViaWx(ctx iris.Context) {
	code := ctx.GetHeader("wx_code")
	respGjson, err := wxsdk.DefaultApp.Code2Session(code).GetJson()
	if err != nil {
		ctx.JSON(newResp(0, err.Error(), nil))
		return
	}

	openID := respGjson.Get("openid").String()
	sessionKey := respGjson.Get("session_key").String()
	if openID == "" {
		ctx.JSON(newResp(0, "openid is empty", nil))
		return
	}

	wxUserInfo := &model.WxUser{
		Openid:     openID,
		Sessionkey: sessionKey,
	}

	if model.DB.NewRecord(wxUserInfo) {
		model.DB.Create(wxUserInfo)
	} else {
		model.DB.Model(&wxUserInfo).Update("sessionkey", sessionKey)
	}

}
