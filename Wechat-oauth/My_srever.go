package main

import (
	"net/http"
	"fmt"

	"log"
	"io/ioutil"
	"encoding/json"

)

type NUM struct {
	OPENID string
	ACCESS_TOKEN string
}

func check(e error)  {
	if e != nil{
		panic(e)
	}
}
//http%3a%2f%2fwww.kit.com%3a8880%2fhello
//https://open.weixin.qq.com/connect/qrconnect?appid=wx5391738a2c43387a&redirect_uri=http%3a%2f%2fwww.kit.com%3a8887%2fhello&response_type=code&scope=snsapi_login,snsapi_userinfo&state=3d6be0a4035d839573b04816624a415e#wechat_redirect
func CodeHander(w http.ResponseWriter, r *http.Request) {

	code_num := r.URL.Query().Get("code")
	URL := "https://api.weixin.qq.com/sns/oauth2/access_token?appid=wx5391738a2c43387a&secret=f500b4f68f74742f6f57cae2769458b4&code="+code_num+"&grant_type=authorization_code"
	resp, err := http.Get(URL)
	check(err)
	defer resp.Body.Close()
	date, err2 := ioutil.ReadAll(resp.Body)
	check(err2)
	fmt.Println("resp: ",string(date))
//{"access_token":"7__myawMwVDYIDpC8_xcwq1RytSp0lENg8zZ09I-r8wcbvQvkkFEhABeT8Rbz0CPaLBFI7hnHUCASNtBv6UbneeQ","expires_in":7200,"refresh_token":"7_BqgBc48uj37V7zkgMMbyip8_EjJeQQFWLtwkhMpSsEp3uiTKElg8jvnPgqkaHliVPQrhCB2X4-V1I8SyL_9H3w","openid":"oSVxI0w0l1LzHrgletaH8be8VfNw","scope":"snsapi_userinfo"}

	var dat NUM
	if err := json.Unmarshal([]byte(date),&dat);err == nil{
		fmt.Println(dat.OPENID)
		fmt.Println(dat.ACCESS_TOKEN)
	}

	URL1 := "https://api.weixin.qq.com/sns/userinfo?access_token="+dat.ACCESS_TOKEN+"&openid="+dat.OPENID
	resp1, err3 := http.Get(URL1)
	check(err3)
	defer resp1.Body.Close()
	date1, err4 := ioutil.ReadAll(resp1.Body)//获取网页数据
	check(err4)
	fmt.Println("resp: ",string(date1))
	var dat1 map[string]interface{}
	json.Unmarshal([]byte(date1),&dat1)//json转map类型
	fmt.Println("dat: ",dat1)
	fmt.Println("\n===========json_dat[]===========")
	for k, v := range dat1{
		fmt.Println(k," --> ",v)
	}

	////遍历结构体
	//t := reflect.TypeOf(user)
	//v := reflect.ValueOf(user)
	//for k := 0; k < t.NumField(); k++{
	//	fmt.Printf("%s -- %v \n", t.Field(k).Name, v.Field(k).Interface())
	//}

}
//http://www.kit.com:8880/hello    进行urlEncode
//http%3a%2f%2fwww.kit.com%3a8880%2fhello
func main() {
	http.HandleFunc("/hello", CodeHander)
	err := http.ListenAndServe(":8887", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
















/*URL:= `https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=wx5391738a2c43387a&secret=f500b4f68f74742f6f57cae2769458b4`
	resp,err := http.Get(URL)
	check(err)
	defer resp.Body.Close()
	date,err1 := ioutil.ReadAll(resp.Body)
	check(err1)
	fmt.Println("date: ",string(date))

	var dat map[string]string
	json.Unmarshal(date,&dat)
	fmt.Println("dat",dat)
	for k, v := range dat{
		//fmt.Printf("%5s --> %s\n",k,v)
		fmt.Println(k," --> ",v)
	}
	fmt.Println("access_token:",dat["access_token"])

	fmt.Println("===============================")
	//URL2 := "https://api.weixin.qq.com/sns/userinfo?access_token="+dat["access_token"]+"&openid="+str
	URL2 := `https://api.weixin.qq.com/sns/userinfo?access_token=ACCESS_TOKEN&openid=OPENID`
	resp1,err1 := http.Get(URL2)
	check(err1)
	defer resp.Body.Close()
	date1,err2 := ioutil.ReadAll(resp1.Body)
	check(err2)
	fmt.Println("date1: ",string(date1))

	var dat1 map[string]interface{}
	json.Unmarshal(date1,&dat1)
	fmt.Println("dat",dat1)
	for k, v := range dat1{
		//fmt.Printf("%5s --> %s\n",k,v)
		fmt.Println(k," --> ",v)
	}
	*/
