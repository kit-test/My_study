package weixin

import (
	"aladinfun.com/midware/tools/sdk/common"
)

type UserInfoRsp struct { //用户信息存储
	common.ErrBase
	common.UserInfoBase
	Province  		string   `json:"province"`
	City      		string   `json:"city"`
	Country   		string   `json:"country"`
	Privilege 		[]string `json:"privilege"`
	Unionid   		string   `json:"unionid"`
	Language        string   `json:"language"`
}
type UserInfoReq struct { //前端回送关键字解析
	Code 			string 	`json:"js_code" weixin:"js_code"`
	AccessToken 	string 	`json:"access_token"weixin:"accrss_token"`
	RefreshToken 	string 	`json:"refresh_token"weixin:"refresh_token"`
	Openid 			string	`json:"openid"weixin:"openid"`
}
type wxUserInfo struct {//用户信息
	Errcode  	  	int      `json:"errcode"`
	Errmsg     		string   `json:"errmsg"`
	OpenID     		string   `json:"open_id"`
	Nickname   		string   `json:"nickname"`
	Sex        		int      `json:"sex"`
	Language 		string	 `json:"language"`
	Province   		string   `json:"province"`
	City       		string   `json:"city"`
	Country    		string   `json:"country"`
	Headimgurl 		string   `json:"headimgurl"`
}
type accessTokenArgs struct {
	Code			string 	`json:"code"`
	GrantType 		string 	`json:"grant_type"`
	Appid     		string 	`json:"appid"`
	Secret   	 	string 	`json:"secret"`

}
type AccessTokenRsp struct {
	common.ErrBase
	AccessToken 	string 	`json:"access_token"`
	ExpiresIn		string 	`json:"expires_in"`
	RefreshToken 	string 	`json:"refresh_token"`
	Openid 			string	`json:"openid"`
	Scope			string	`json:"Scope"`
	Unionid			string 	`json:"Unionid"`
}
type commonArgs struct { //查验/获取信息共用
	AccessToken 	string 	`json:"access_token"`
	Openid 			string 	`json:"openid"`
}
