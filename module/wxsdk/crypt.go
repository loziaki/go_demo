package wxsdk

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"

	"github.com/tidwall/gjson"
)

func (app *WxApp) DecryptSecureData(encryptedData string, sessionKey string, iv string) (map[string]interface{}, error) {
	ctypter := &WxBizDataCrypt{app.appid, sessionKey}
	bytes, err := ctypter.Decrypt(encryptedData, iv)
	if err != nil {
		return nil, err
	}

	data := gjson.ParseBytes(bytes)
	// m["watermark"].(map[string]interface{})["appid"] != app.appid
	if data.Get("watermark.appid").String() != app.appid {
		return nil, errors.New("watermark.appID is not match")
	}

	m, ok := data.Value().(map[string]interface{})
	if !ok {
		return nil, errors.New("response data can not parse")
	}
	return m, nil
}

type WxBizDataCrypt struct {
	AppID      string
	SessionKey string
}

func (wxCrypt *WxBizDataCrypt) Decrypt(encryptedData string, iv string) ([]byte, error) {
	if len(wxCrypt.SessionKey) != 24 {
		return nil, errors.New("sessionKey length is error")
	}
	aesKey, err := base64.StdEncoding.DecodeString(wxCrypt.SessionKey)
	if err != nil {
		return nil, err
	}

	if len(iv) != 24 {
		return nil, errors.New("iv length is error")
	}
	aesIV, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return nil, err
	}

	aesCipherText, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return nil, err
	}
	aesPlantText := make([]byte, len(aesCipherText))

	aesBlock, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}

	mode := cipher.NewCBCDecrypter(aesBlock, aesIV)
	mode.CryptBlocks(aesPlantText, aesCipherText)
	aesPlantText = PKCS7UnPadding(aesPlantText)

	return aesPlantText, nil
}

// PKCS7UnPadding return unpadding []Byte plantText
func PKCS7UnPadding(plantText []byte) []byte {
	length := len(plantText)
	unPadding := int(plantText[length-1])
	return plantText[:(length - unPadding)]
}
