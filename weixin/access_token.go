package weixin

import (
	"aladinfun.com/midware/tools/sdk/common"
	"aladinfun.com/midware/tools/sdk/config"
	"aladinfun.com/midware/tools/sdk/sdklog"
	"encoding/json"
	"aladinfun.com/libs/logger"
	"net/url"

	"fmt"
)

func AccessToken(buf []byte) (response common.Response, err error) {
	fmt.Println("weixinAccessToken.buf: ",string(buf))//**
	rsp := &AccessTokenRsp{}
	req := &UserInfoReq{}
	//解析前端的数据
	if err = json.Unmarshal(buf, req); err != nil {
		sdklog.LevelLogln(logger.ErrorLog, "unmarshal json failed, %s", err.Error())
		return
	}
	response = rsp   //???????????????
	fmt.Println(rsp)
	args := &accessTokenArgs{
		Code:		req.Code,
		GrantType: 	config.WeiXgrantType(),
		Appid:		config.WeiXappid(),
		Secret:   	config.WeiXsecret(),
	}
	fmt.Println("weixinAccessToken.args: ",args)//**
	v:= url.Values{}
	common.UrlQuery(args,&v)
	var urll string
	urll = config.WeiXaccessTokenURL()
	urll += "?" + v.Encode()
	fmt.Println("weixinAccessToken.urll: ",urll)//**
	rspBuf ,err := common.SendHTTPReq(common.DefaultCli,"GET",urll,nil)
	if err != nil{
		return
	}
	fmt.Println("weixinAccessToken.err: ",err)//**
	fmt.Println("weixinAccessToken.rspBuf: ",string(rspBuf))//**
	err1 := json.Unmarshal(rspBuf,rsp)
	if err1 !=nil{
		return
	}

	return
}
