package wxsdk

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
)

func (app *WxApp) GetAccessToken() *WxApiResponse {
	queryValue := url.Values{}
	queryValue.Set("grant_type", "client_credential")
	queryValue.Set("appid", app.appid)
	queryValue.Set("secret", app.appSecret)

	wxAPIURL, _ := url.Parse("http://dev.hylinkgz.cn/192/test2.php")
	wxAPIURL.RawQuery = queryValue.Encode()
	resp, err := http.Get(wxAPIURL.String())

	result := new(WxApiResponse)
	if err != nil {
		result.err = err
	} else {
		result.raw = resp
	}

	return result
}

func (app *WxApp) CreateQRCode(path string, width int) *WxApiResponse {
	switch {
	case width < 280:
		width = 280
	case width > 1280:
		width = 1280
	}

	act, err := app.GetLatestAccessToken()
	if err != nil {
		return &WxApiResponse{err: err}
	}

	queryValue := url.Values{}
	queryValue.Set("access_token", act)

	wxAPIURL, _ := url.Parse("http://dev.hylinkgz.cn/192/test.php")
	wxAPIURL.RawQuery = queryValue.Encode()
	requestURL := wxAPIURL.String()

	data := map[string]string{"path": path, "width": string(width)}
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return &WxApiResponse{err: err}
	}

	result := new(WxApiResponse)
	resp, err := http.Post(requestURL, "application/json", bytes.NewReader(dataBytes))
	if err != nil {
		result.err = err
	} else {
		result.raw = resp
	}
	return result
}

func (app *WxApp) Code2Session(code string) *WxApiResponse {
	queryValue := url.Values{}
	queryValue.Set("grant_type", "authorization_code")
	queryValue.Set("appid", app.appid)
	queryValue.Set("secret", app.appSecret)
	queryValue.Set("js_code", code)

	wxAPIURL, _ := url.Parse("http://dev.hylinkgz.cn/192/test3.php")
	wxAPIURL.RawQuery = queryValue.Encode()
	resp, err := http.Get(wxAPIURL.String())
	if err != nil {
		return &WxApiResponse{err: err}
	}

	result := new(WxApiResponse)
	result.raw = resp
	return result
}
