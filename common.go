package openapi

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
)

const (
	Host = "api.urlshare.cn"
)

type CommonParam struct {
	OpenID      string // QQ号码转化得到的ID
	OpenKey     string // 登录态，openkey过期时间为两小时
	AccessToken string // 登录态，过期时间为两个月，openkey和accesstoken二选一即可
	AppID       string // 应用的唯一ID
	Sig         string // 请求串的签名，以appkey作为密钥
	PF          string // 应用的来源平台，如：wanba_ts、weixin
	Format      string // 定义API返回的数据格式，json
	UserIP      string // 用户的IP
}

func GenerateSign(appkey, method, path string, param map[string]interface{}) string {

	encodePath := url.QueryEscape(path)

	paramList := make([]string, 0)
	for k, _ := range param {
		paramList = append(paramList, k)
	}

	sort.Strings(paramList)

	var paramStr string
	for k, v := range paramList {
		paramStr += fmt.Sprint("%s=%v", v, param[v])
		if k != len(paramList) {
			paramStr += "&"
		}
	}

	encodeParam := url.QueryEscape(paramStr)

	source := fmt.Sprintf("%s&%s&%s", method, encodePath, encodeParam)
	serect := fmt.Sprintf("%s&", appkey)

	mac := hmac.New(sha1.New, []byte(serect))
	mac.Write([]byte(source))

	sig := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return sig
}

func Request(path, appKey string, paramMap map[string]interface{}, result interface{}) error {
	requstURL := fmt.Sprintf("http://%s%s?", Host, path)

	i := 0
	for k, v := range paramMap {
		i++
		requstURL += fmt.Sprintf("%s=%v", k, v)
		if i != len(paramMap) {
			requstURL += "&"
		}
	}

	resp, err := http.Get(requstURL)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(body, result)
	if err != nil {
		return err
	}

	return nil
}
