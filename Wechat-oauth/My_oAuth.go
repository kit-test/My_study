package main

import (
	"net/http"
	"io"
	"time"
	"fmt"
	"io/ioutil"
	"encoding/json"
)
type KEY struct {
	OPENID string
	ACCESS_TOKEN string
}
type TOKEN struct {
	Errcode int
	Errmsg string
}

type USREinfo struct {
		Openid string
		Nickname string
		SEX int
		Language string
		City string
		Province string
		Country string
		Headimgurl string
	}
func check(e error)  {
	if e != nil{
		panic(e)
	}
}
func KitHander(w http.ResponseWriter,r *http.Request)  {
	date,err := ioutil.ReadFile("./kit_html.txt")
	check(err)
	date_str := string(date)
	if r.Method == "GET"{
		io.WriteString(w,date_str)
	}
}
func UserHander(w http.ResponseWriter,r *http.Request)  {

	//获取CODE数据
	code_num := r.URL.Query().Get("code")
	fmt.Println("code: ",code_num)//***
	//拼接CODE到获取access_token的来链接中//appid,secret等关键字不区分大小写但拼写必须正确
	URL := "https://api.weixin.qq.com/sns/oauth2/access_token?Appid=wx5391738a2c43387a&secret=f500b4f68f74742f6f57cae2769458b4&grant_type=authorization_code"+"&code="+code_num
		//	https://api.weixin.qq.com/sns/oauth2/access_token?appid=wx5391738a2c43387a&secret=f500b4f68f74742f6f57cae2769458b4&grant_type=authorization_code&code=071tQSQM1XLV551ccvQM1nzSQM1tQSQt&grant_type=authorization_code
	resp, err1 := http.Get(URL)
	check(err1)
	defer resp.Body.Close()

	//获取微信返回数据(Json结构)
	date, err2 := ioutil.ReadAll(resp.Body)
	check(err2)
	//{"access_token":"7__myawMwVDYIDpC8_xcwq1RytSp0lENg8zZ09I-r8wcbvQvkkFEhABeT8Rbz0CPaLBFI7hnHUCASNtBv6UbneeQ","expires_in":7200,"refresh_token":"7_BqgBc48uj37V7zkgMMbyip8_EjJeQQFWLtwkhMpSsEp3uiTKElg8jvnPgqkaHliVPQrhCB2X4-V1I8SyL_9H3w","openid":"oSVxI0w0l1LzHrgletaH8be8VfNw","scope":"snsapi_userinfo"}
	//解析Json数据为struct结构，获取access_token和openid
	var key KEY
	if err3 := json.Unmarshal([]byte(date),&key);err3 == nil{
		fmt.Println("key.OPENID: ",key.OPENID)
		fmt.Println("key.ACCESS_TOKEN: ",key.ACCESS_TOKEN)
	}

	fmt.Println("accesstoken: ",key.ACCESS_TOKEN)//**
	fmt.Println("OPENID: ",key.OPENID)//**

	//判断access_token是否过期
	//https://api.weixin.qq.com/sns/auth?access_token=ACCESS_TOKEN&openid=OPENID
	URL_token := "https://api.weixin.qq.com/sns/auth?access_token="+key.ACCESS_TOKEN +"&openid="+key.OPENID
				//https://api.weixin.qq.com/sns/auth?access_token=8_WAsaa6hIFwtqHrgOUyXFXI90reiS08Ma2PbutCUTujj8y3r26m3kRuIExYnp5pEGbd3S_ut8kj-CvlMDGw6keg&openid=oSVxI0w0l1LzHrgletaH8be8VfNw
	resp_token, err4 := http.Get(URL_token)
	check(err4)
	defer resp_token.Body.Close()
	date_token,err5 := ioutil.ReadAll(resp_token.Body)
	check(err5)
	var token TOKEN
	err6 := json.Unmarshal(date_token,&token)
	if err6 == nil{
		fmt.Println("token.Errcode",token.Errcode)
		fmt.Println("token.Errmsg",token.Errmsg)
	}

	//拼接access_token和openid用于端口访问，获取用户子信息
	URL1 := "https://api.weixin.qq.com/sns/userinfo?access_token="+key.ACCESS_TOKEN+"&openid="+key.OPENID
		   //https://api.weixin.qq.com/sns/userinfo?access_token=8_WAsaa6hIFwtqHrgOUyXFXI90reiS08Ma2PbutCUTujj8y3r26m3kRuIExYnp5pEGbd3S_ut8kj-CvlMDGw6keg&openid=oSVxI0w0l1LzHrgletaH8be8VfNw
	resp_user, err7 := http.Get(URL1)
	check(err7)
	defer resp_user.Body.Close()
	date_user, err8 := ioutil.ReadAll(resp_user.Body)
	check(err8)
	var user USREinfo
	if err9 := json.Unmarshal([]byte(date_user),&user);err9 == nil {
		fmt.Println(user)
		fmt.Println("user.Openid: ",user.Openid)
		fmt.Println("user.Openid: ",user.Openid)
	}

}

func main() {
	var i chan int
//	var j chan int
	kit := http.NewServeMux()
	kit.HandleFunc("/kit", KitHander)
	USER := http.NewServeMux()
	USER.HandleFunc("/hello",UserHander)
	var err error
	go func() {
		http.ListenAndServe(":8888",kit)
		return
	}()
	time.Sleep(200*time.Millisecond)
	if err != nil{
		fmt.Println("ListenAndServe:8888: ",err)
		return
	}

	go func() {
		http.ListenAndServe(":8887", USER)
		i <- 1
		return
	}()
	time.Sleep(200*time.Millisecond)
	if err != nil{
		fmt.Println("ListenAndServe:8887: ",err)

		return
	}
	<-i

}