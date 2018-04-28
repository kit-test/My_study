package weixin

import (
	"aladinfun.com/midware/tools/sdk/common"
	"encoding/json"
	"aladinfun.com/midware/tools/sdk/sdklog"
	"aladinfun.com/libs/logger"
	"net/url"
	"aladinfun.com/midware/tools/sdk/config"
	"fmt"

)

func UserInfo(buf []byte) (response common.Response, err error) {
	fmt.Println("weixinUserInfo.buf: ",string(buf))//**
	req := &UserInfoReq{}
	rsp := &UserInfoRsp{}
	response = rsp

	if err = json.Unmarshal(buf, req); err != nil {
		sdklog.LevelLogln(logger.ErrorLog, "unmarshal json failed, %s", err.Error())
		return
	}
	//判断access_token是否有效
	args := &commonArgs{
		AccessToken:  	req.AccessToken,
		Openid :		req.Openid,
	}
	fmt.Println("weixinUserInfo.args: ",args)//**
	v := url.Values{}
	common.UrlQuery(args,&v)

	var urll string
	urll = config.WeiXrefreshTokenURL()
	urll += "?"+v.Encode()
	fmt.Println("weixinUserInfo.fresh_urll: ",urll)//**
	rspBuf ,err := common.SendHTTPReq(common.DefaultCli,"GET",urll,nil)
	//if err != nil{
	//	return
	//}
	fmt.Println("weixinUserInfo.fresh_err: ",err)//**
	fmt.Println("weixinUserInfo.fresh_rspBuf: ",string(rspBuf))//**
	wxRsp := &wxUserInfo{}
	err1 := json.Unmarshal(rspBuf,wxRsp)
	if err1!= nil{
		return
	}
	fmt.Println("weixinUserInfo.fresh_wxRsp.Errmsg: ",wxRsp.Errmsg)//**
	//if wxRsp.Errmsg != "ok"{
	//	return
	//
	//
	//}
	//取用户数据
	user_args := &commonArgs{
		AccessToken:  	req.AccessToken,
		Openid :		req.Openid,
	}
	user_v := url.Values{}
	common.UrlQuery(user_args,&user_v)

	var user_urll string
	user_urll = config.WeiXuserInFoURL()
	fmt.Println("weixinUserInfo.user_urll_a: ",user_urll)//**
	user_urll += "?"+v.Encode()
	fmt.Println("weixinUserInfo.user_urll:_l ",user_urll)//**

	user_rspBuf ,err := common.SendHTTPReq(common.DefaultCli,"GET",urll,nil)
	//if err != nil{
	//	return
	//}
	fmt.Println("weixinUserInfo.err: ",err)//**
	fmt.Println("weixinUserInfo.user_rspBuf: ",string(user_rspBuf))//**
	wxUser:= &wxUserInfo{}
	err2 := json.Unmarshal(user_rspBuf,wxUser)
	if err2!= nil{
		return
	}
	rsp.OpenID = wxUser.OpenID
	rsp.Name = wxUser.Nickname
	rsp.Gender = wxUser.Sex
	rsp.HeadURL = wxUser.Headimgurl
	rsp.Country = wxUser.Country
	rsp.Province = wxUser.Province
	rsp.City    =  wxUser.City
	rsp.Language = wxUser.Language


	return
}
