package openapi

type UserInfo struct {
	Ret             int32  `json:"ret"`
	Msg             string `json:"msg"`
	IsLost          int32  `json:"is_lost"`
	NickName        string `json:"nickname"`
	Gender          string `json:"gender"`
	Country         string `json:"country"`
	Province        string `json:"province"`
	City            string `json:"city"`
	FigureURl       string `json:"firgureurl"`
	IsYellowVIP     int32  `json:"is_yellow_vip"`
	IsYellowYearVIP int32  `json:"is_yellow_year_vip"`
	YellowVIPLevel  int32  `json:"yellow_vip_level"`
	IsYellowHighVIP int32  `json:"is_yellow_high_vip"`
}

// https://wiki.qzone.qq.com/openapi/v3/user/get_info
func GetInfo(appkey string, param *CommonParam) (error, *UserInfo) {
	path := "/v3/user/get_info"
	paramMap := make(map[string]interface{})

	paramMap["openid"] = param.OpenID
	if param.OpenKey != "" {
		paramMap["openkey"] = param.OpenKey
	}
	if param.AccessToken != "" {
		paramMap["accesstoken"] = param.AccessToken
	}
	paramMap["appid"] = param.AppID
	paramMap["pf"] = param.PF
	paramMap["format"] = param.Format
	paramMap["userip"] = param.UserIP

	sig := GenerateSign(appkey, "GET", path, paramMap)
	paramMap["sig"] = sig

	var ret UserInfo
	err := Request(path, appkey, paramMap, &ret)
	return err, &ret
}

type LoginStatus struct {
	Ret int32  `json:"ret"`
	Msg string `json:"msg"`
}

// https://wiki.qzone.qq.com/openapi/v3/user/is_login
func IsLogin(appkey string, param *CommonParam) (error, *LoginStatus) {
	path := "/v3/user/is_login"
	paramMap := make(map[string]interface{})

	paramMap["openid"] = param.OpenID
	if param.OpenKey != "" {
		paramMap["openkey"] = param.OpenKey
	}
	if param.AccessToken != "" {
		paramMap["accesstoken"] = param.AccessToken
	}
	paramMap["appid"] = param.AppID
	paramMap["pf"] = param.PF
	paramMap["format"] = param.Format
	paramMap["userip"] = param.UserIP

	sig := GenerateSign(appkey, "GET", path, paramMap)
	paramMap["sig"] = sig

	var ret LoginStatus
	err := Request(path, appkey, paramMap, &ret)
	return err, &ret
}

type AppFriends struct {
	Ret    int32  `json:"ret"`
	Msg    string `json:"msg"`
	IsLost int32  `json:"is_lost"`
	Items  []struct {
		OpenID string `json:"openid"`
	} `json:"items"`
}

// https://wiki.qzone.qq.com/openapi/v3/relation/get_app_friends
func GetAppFriend(appkey string, param *CommonParam) (error, *AppFriends) {
	path := "/v3/relation/get_app_friends"
	paramMap := make(map[string]interface{})

	paramMap["openid"] = param.OpenID
	if param.OpenKey != "" {
		paramMap["openkey"] = param.OpenKey
	}
	if param.AccessToken != "" {
		paramMap["accesstoken"] = param.AccessToken
	}
	paramMap["appid"] = param.AppID
	paramMap["pf"] = param.PF
	paramMap["format"] = param.Format
	paramMap["userip"] = param.UserIP

	sig := GenerateSign(appkey, "GET", path, paramMap)
	paramMap["sig"] = sig

	var ret AppFriends
	err := Request(path, appkey, paramMap, &ret)
	return err, &ret
}

type PlayzoneUserInfo struct {
	Code            int32  `json:"code"`
	Subcode         int32  `json:"subcode"`
	Message         string `json:"message"`
	IsVIP           int32  `json:"is_vip"`
	VIPLevel        int32  `json:"vip_level"`
	IsWhiteListUser int32  `json:"isWhiteListUser"`
	ExpiredTime     int32  `json:"expiredtime"`
}

// https://wiki.qzone.qq.com/openapi/v3/user/get_playzone_userinfo
func GetPlayzoneUserInfo(appkey string, param *CommonParam, zoneID int32) (error, *PlayzoneUserInfo) {
	path := "/v3/user/get_playzone_userinfo"
	paramMap := make(map[string]interface{})

	paramMap["openid"] = param.OpenID
	if param.OpenKey != "" {
		paramMap["openkey"] = param.OpenKey
	}
	if param.AccessToken != "" {
		paramMap["accesstoken"] = param.AccessToken
	}
	paramMap["appid"] = param.AppID
	paramMap["pf"] = param.PF
	paramMap["format"] = param.Format
	paramMap["userip"] = param.UserIP
	paramMap["zoneid"] = zoneID

	sig := GenerateSign(appkey, "GET", path, paramMap)
	paramMap["sig"] = sig

	var ret PlayzoneUserInfo
	err := Request(path, appkey, paramMap, &ret)
	return err, &ret
}

type BuyPlayzoneItemRet struct {
	Code            int32  `json:"code"`
	Subcode         int32  `json:"subcode"`
	Message         string `json:"message"`
	Billno          string `json:"billno"`
	IsWhiteListUser int32  `json:"isWhiteListUser"`
	Cost            int32  `json:"cost"`
}

// https://wiki.qzone.qq.com/openapi/v3/user/buy_playzone_item
func BuyPlayzoneItem(appkey string, param *CommonParam, zoneID int32, billno string, itemID string, count int32) (error, *BuyPlayzoneItemRet) {
	path := "/v3/user/buy_playzone_item"
	paramMap := make(map[string]interface{})

	paramMap["openid"] = param.OpenID
	if param.OpenKey != "" {
		paramMap["openkey"] = param.OpenKey
	}
	if param.AccessToken != "" {
		paramMap["accesstoken"] = param.AccessToken
	}
	paramMap["appid"] = param.AppID
	paramMap["pf"] = param.PF
	paramMap["format"] = param.Format
	paramMap["userip"] = param.UserIP
	paramMap["zoneid"] = zoneID
	paramMap["billno"] = billno
	paramMap["itemid"] = itemID
	paramMap["count"] = count

	sig := GenerateSign(appkey, "GET", path, paramMap)
	paramMap["sig"] = sig

	var ret BuyPlayzoneItemRet
	err := Request(path, appkey, paramMap, &ret)
	return err, &ret
}

type GamebarMsg struct {
	Ret     int32  `json:"ret"`
	Message string `json:"message"`
}

// https://wiki.qzone.qq.com/openapi/v3/user/send_gamebar_msg
func SendGamebarMsg(appkey string, param *CommonParam, frd string, msgType int32, content string) (error, *GamebarMsg) {
	path := "/v3/user/send_gamebar_msg"
	paramMap := make(map[string]interface{})

	paramMap["openid"] = param.OpenID
	if param.OpenKey != "" {
		paramMap["openkey"] = param.OpenKey
	}
	if param.AccessToken != "" {
		paramMap["accesstoken"] = param.AccessToken
	}
	paramMap["appid"] = param.AppID
	paramMap["pf"] = param.PF
	paramMap["format"] = param.Format
	paramMap["userip"] = param.UserIP
	paramMap["frd"] = frd
	paramMap["msgtype"] = msgType
	paramMap["cotent"] = content

	sig := GenerateSign(appkey, "GET", path, paramMap)
	paramMap["sig"] = sig

	var ret GamebarMsg
	err := Request(path, appkey, paramMap, &ret)
	return err, &ret
}
