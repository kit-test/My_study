package main

import (
	"net/http"
	"io/ioutil"
	"io"
	"log"
	"fmt"

)
func check(e error)  {
	if e != nil{
		panic(e)
	}
}
func kitHander(w http.ResponseWriter, r*http.Request)  {
	date,err := ioutil.ReadFile("./kit_html.txt")
	check(err)
	date_str := string(date)
	if r.Method == "GET"{
		io.WriteString(w,date_str)
	}
}
/*
func CodeHander(w http.ResponseWriter, r *http.Request) {
	URL:= `https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=wx5391738a2c43387a&secret=f500b4f68f74742f6f57cae2769458b4`
	resp,err := http.Get(URL)
	check(err)
	defer resp.Body.Close()
	date,err1 := ioutil.ReadAll(resp.Body)
	check(err1)

	var dat map[string]interface{}
	json.Unmarshal(date,&dat)
	fmt.Println("dat",dat)
	for k, v := range dat{
		//fmt.Printf("%5s --> %s\n",k,v)
		fmt.Println(k," --> ",v)

	}
}*/
func main() {
	http.HandleFunc("/kit", kitHander)
//	http.HandleFunc("/hello",  CodeHander)
	err := http.ListenAndServe(":8886", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
	fmt.Println("===================")
}



















/*
func ViewHander(w http.ResponseWriter, r *http.Request)  {
	URL := "https://open.weixin.qq.com/connect/qrconnect?"+"" +
		"appid=wx5391738a2c43387a&"+
		"redirect_uri=http%3a%2f%2flocalhost%3a8888%2fhello"+
		"response_type=code&"+
		"scope=snsapi_login&" +
		"state=3d6be0a4035d839573b04816624a415e#wechat_redirect"
	resp,err := http.Get(URL) //redirect_uri:跳转URL
	check(err)
	defer resp.Body.Close()
	fmt.Println("resp: ",resp)

	date, err1 := ioutil.ReadAll(resp.Body) //读取服务器发送的数据
	check(err1)
	str := string(date)
	fmt.Println("str: ",str )
	if r.Method == "GET"{
		io.WriteString(w,str)
	}
}*/