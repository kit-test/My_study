package My_function
import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"

	"io"
	"time"
)
type KEY struct {
	OPENID string
	ACCESS_TOKEN string
	Refresh_token string
	Expires_in string
	Scope string

}
type TOKEN struct {
	Errcode int
	Errmsg string
}
type USERinfo struct {
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
var(
	AppId string
	Secret string
)
func KitHander(w http.ResponseWriter,r *http.Request)  {
	date,err := ioutil.ReadFile("./kit_html.txt")
	check(err)
	date_str := string(date)
	if r.Method == "GET"{
		io.WriteString(w,date_str)
	}
}
func UserHander_do(w http.ResponseWriter,r *http.Request) {
	//AppId :="wx5391738a2c43387a"
	//Secret :="f500b4f68f74742f6f57cae2769458b4"
	//获取CODE数据
	code_num := r.URL.Query().Get("code")

	//拼接CODE到获取access_token的来链接中
	URL := "https://api.weixin.qq.com/sns/oauth2/access_token?appid="+AppId+"&secret="+Secret+"&code="+code_num+"&grant_type=authorization_code"
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
	/*
	//判断access_token是否过期
	//https://api.weixin.qq.com/sns/auth?access_token=ACCESS_TOKEN&openid=OPENID
	URL_token := "https://api.weixin.qq.com/sns/auth?access_token="+key.ACCESS_TOKEN +"&openid="+key.OPENID
	resp_token, err4 := http.Get(URL_token)
	check(err4)
	defer resp_token.Body.Close()
	date_token,err5 := ioutil.ReadAll(resp_token.Body)
	check(err5)
	var token TOKEN
	err6 := json.Unmarshal(date_token,&token)
	if err6 == nil{
		fmt.Println("token.Errcode： ",token.Errcode)
		fmt.Println("token.Errmsg： ",token.Errmsg)
	}

	if token.Errmsg != "ok" && key.Refresh_token != ""{
		//https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=APPID&grant_type=refresh_token&refresh_token=REFRESH_TOKEN
		URL_refresh_token := "https://api.weixin.qq.com/sns/oauth2/refresh_token?appid="+AppId+"grant_type=refresh_token&refresh_token="+key.Refresh_token
		resp_refresh_token,err10 := http.Get(URL_refresh_token)
		check(err10)
		defer resp_refresh_token.Body.Close()
		date_refresh_token,err11 := ioutil.ReadAll(resp_refresh_token.Body)
		check(err11)
		err12 := json.Unmarshal(date_refresh_token,&key)
		if err12 == nil{
			fmt.Println("key.OPENID: ",key.OPENID)
			fmt.Println("key.ACCESS_TOKEN: ",key.ACCESS_TOKEN)
		}
	}*/
	//拼接access_token和openid用于端口访问，获取用户子信息
	URL1 := "https://api.weixin.qq.com/sns/userinfo?access_token="+key.ACCESS_TOKEN+"&openid="+key.OPENID
	resp_user, err7 := http.Get(URL1)
	check(err7)
	defer resp_user.Body.Close()
	date_user, err8 := ioutil.ReadAll(resp_user.Body)
	check(err8)
	var user USERinfo
	if err9 := json.Unmarshal([]byte(date_user),&user);err9 == nil {
		fmt.Println(user)
		fmt.Println("user.Openid: ",user.Openid)
		fmt.Println("user.Openid: ",user.Openid)
	}
}
func UserHander(appid string,secret string)  {
	fmt.Println("Here is Userhander")
	AppId = appid
	Secret = secret
	fmt.Println(AppId)
	fmt.Println(Secret)

	kit := http.NewServeMux()
	kit.HandleFunc("/kit", KitHander)
	USER := http.NewServeMux()
	USER.HandleFunc("/hello",UserHander_do)
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
		return
	}()
	time.Sleep(200*time.Millisecond)
	if err != nil{
		fmt.Println("ListenAndServe:8887: ",err)
		return
	}
	time.Sleep(10*time.Second)
}
