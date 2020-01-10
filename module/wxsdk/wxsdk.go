package wxsdk

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/tidwall/gjson"
)

var DefaultApp = &defaultApp

var defaultApp WxApp

func init() {
	DefaultApp = NewDefaultWxApp()
}

func NewDefaultWxApp() *WxApp {
	appid := "bububu"
	appSecret := "lalala"
	return NewWxApp(appid, appSecret)
}

func NewWxApp(appid, appSecret string) *WxApp {
	app := new(WxApp)
	app.appid = appid
	app.appSecret = appSecret
	app.WxLoginExpires = 7200                        //s
	var networkTimeout time.Duration = 5 * 100000000 //nanosecond
	http.DefaultClient.Timeout = networkTimeout
	return app
}

type WxApp struct {
	appid          string
	appSecret      string
	act            *AccessToken
	WxLoginExpires time.Duration
}

func (app *WxApp) GetLatestAccessToken() (string, error) {
	if app.act == nil || app.act.ifExpire() {
		wxRespJSON, err := app.GetAccessToken().GetJson()
		if err != nil {
			return "", err
		}

		accessToken := new(AccessToken)
		accessToken.Value = wxRespJSON.Get("access_token").String()
		expiredIn := wxRespJSON.Get("expires_in").Int()
		expired := time.Now().Unix() + expiredIn
		accessToken.Expired = time.Unix(expired, 0)

		app.act = accessToken
	}
	return app.act.Value, nil
}

type AccessToken struct {
	Expired time.Time
	Value   string
}

func (token *AccessToken) ifExpire() bool {
	if token == nil {
		return true
	}
	return time.Now().After(token.Expired)
}

type WxApiResponse struct {
	raw *http.Response
	err error
}

func NewWxResp(resp *http.Response) *WxApiResponse {
	return &WxApiResponse{raw: resp}
}

func (wxresp *WxApiResponse) GetBytes() ([]byte, error) {
	if wxresp.err != nil {
		return nil, wxresp.err
	}

	body, err := ioutil.ReadAll(wxresp.raw.Body)
	if err != nil {
		return nil, err
	}

	if gjson.ValidBytes(body) {
		data := gjson.ParseBytes(body)
		if code := data.Get("errcode"); code.Exists() && code.Int() != 0 {
			errmsg := fmt.Sprintf("[WX-ERROR][%d]%s", code.Int(), data.Get("errmsg").String())
			return nil, errors.New(errmsg)
		}
	}

	return body, nil
}

func (resp *WxApiResponse) GetJson() (*gjson.Result, error) {
	body, err := resp.GetBytes()
	if err != nil {
		return nil, err
	}
	data := gjson.ParseBytes(body)
	return &data, nil
}

func (resp *WxApiResponse) GetMap() (map[string]interface{}, error) {
	jsonData, err := resp.GetJson()
	if err != nil {
		return nil, err
	}

	m, ok := jsonData.Value().(map[string]interface{})
	if !ok {
		return nil, err
	}
	return m, nil
}
