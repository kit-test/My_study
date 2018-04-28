package main

import (






	"net/http"
	"Wechat-oauth/Function"
)
///*
//func main()  {
//	for {
//		reader := bufio.NewReader(os.Stdin)
//		input, _ := reader.ReadBytes('\n')
//		str := string(input[0:len(input)-1])
//		if str == "run" {
//			fmt.Println("succeed!!")
//		} else {
//			fmt.Println("false")
//		}
//	}
//}
//*/
///*
//func main()  {
//	//func kitHander(w http.ResponseWriter, r*http.Request)  {
//	date,_ := ioutil.ReadFile("./kit_html.txt")
//	//check(err)
//	fmt.Println(string(date))
//
//	}
//*/
///*
//type NUM struct {
//	OPENID string
//	ACCESS_TOKEN string
//}
//func main()  {
//	URL := `{"access_token":"7__myawMwVDYIDpC8_xcwq1RytSp0lENg8zZ09I-r8wcbvQvkkFEhABeT8Rbz0CPaLBFI7hnHUCASNtBv6UbneeQ","expires_in":7200,"refresh_token":"7_BqgBc48uj37V7zkgMMbyip8_EjJeQQFWLtwkhMpSsEp3uiTKElg8jvnPgqkaHliVPQrhCB2X4-V1I8SyL_9H3w","openid":"oSVxI0w0l1LzHrgletaH8be8VfNw","scope":"snsapi_userinfo"}`
//
///*
//	var config ConfigStruct
//	if err := json.Unmarshal([]byte(jsonStr), &config); err == nil {
//		fmt.Println("================json str 转struct==")
//		fmt.Println(config)
//		fmt.Println(config.Host)
//	}
//*/
///*	var dat NUM
//	if err := json.Unmarshal([]byte(URL),&dat);err == nil{
//		fmt.Println(dat.OPENID)
//		fmt.Println(dat.ACCESS_TOKEN)
//	}
//
///*
//
//	var dat map[string]interface{}
//	json.Unmarshal([]byte(URL),&dat)
//	fmt.Println("dat: ",dat)
//	fmt.Println("\n===========json_dat[]===========")
//	for k, v := range dat{
//		fmt.Println(k," --> ",v)
//	}
//	str := dat["openid"]
//	fmt.Println("dat[openid]:",reflect.TypeOf(dat["openid"]).Kind().String())
//	//var str string
//	//fmt.Fprint(str,"%s",dat["openid"])
//	fmt.Println("+++openid: ","adc"+str+"abc")
//
//	}*/
//
//
///*
//func KitHander(w http.ResponseWriter,r *http.Request)  {
//	io.WriteString(w,"Hello")
//}
//func UserHander(w http.ResponseWriter,r *http.Request)  {
//	io.WriteString(w,"nihao ")
//}
//
//func main() {
//	kit := http.NewServeMux()
//	kit.HandleFunc("/kit", KitHander)
//	USER := http.NewServeMux()
//	USER.HandleFunc("/hello",UserHander)
//	var err error
//	go func() {
//		http.ListenAndServe(":8888",kit)
//		return
//	}()
//	time.Sleep(200*time.Millisecond)
//	if err != nil{
//		fmt.Println("ListenAndServe:8888: ",err)
//		return
//	}
//
//
//	go func() {
//		http.ListenAndServe(":8887", USER)
//		return
//	}()
//	time.Sleep(200*time.Millisecond)
//	if err != nil{
//		fmt.Println("ListenAndServe:8887: ",err)
//		return
//	}
//	select {}
//}
//*/
///*
//func ONE(w http.ResponseWriter,r *http.Request)  {
//	io.WriteString(w,"Hello")
//}
//func TWO(w http.ResponseWriter,r *http.Request)  {
//	io.WriteString(w,"nihao ")
//}
//
//func main()  {
//	AA := http.NewServeMux()
//	AA.HandleFunc("/jj",ONE)
//	BB := http.NewServeMux()
//	BB.HandleFunc("/ll",TWO)
//
//	go func() {
//		http.ListenAndServe(":8888",AA)
//		return
//	}()
//	go func() {
//		http.ListenAndServe(":8880",BB)
//		return
//	}()
//	select {}
//}*/
//
////type USREinfo struct {
////	Openid string
////	Nickname string
////	SEX int
////	Language string
////	City string
////	Province string
////	Country string
////	Headimgurl string
////}
////
////func main()  {
////	URL := `{"openid":"oSVxI0w0l1LzHrgletaH8be8VfNw","nickname":"胡杨下的土豆","sex":1,"language":"zh_CN","city":"","province":"Ottawa","country":"CA","headimgurl":"http:\/\/thirdwx.qlogo.cn\/mmopen\/vi_32\/DYAIOgq83eqdib3R1go2rfmLibog05A8QUuEQby67SBicRZ5D5LoPKSycPiby9oDyjdghKZOolz3PKGxJibiadnUfpzw\/132","privilege":[]}`
//
///*
//	//json str 转struct
//	var config ConfigStruct
//	if err := json.Unmarshal([]byte(jsonStr), &config); err == nil {
//		fmt.Println("================json str 转struct==")
//		fmt.Println(config)
//		fmt.Println(config.Host)
//	}
//*/
///*
//	var user USREinfo
//	fmt.Println("+===========")
//	if err := json.Unmarshal([]byte(URL),&user);err == nil {
//		fmt.Println("err: ", err)
//		fmt.Println(user)
//		fmt.Println(user.Openid, user.Country)
//	}
//	//遍历结构体
//	t := reflect.TypeOf(user)
//	v := reflect.ValueOf(user)
//	for k := 0; k < t.NumField(); k++{
//		fmt.Printf("%s -- %v \n", t.Field(k).Name, v.Field(k).Interface())
//	}
//*/
//	//func main(){
//	//	u := User{Id:1001, Name:"xxx"/*, addr:"xxx"*/}
//	//	t := reflect.TypeOf(u)
//	//	v := reflect.ValueOf(u)
//	//	for k := 0; k < t.NumFiled(); k++ {
//	//		fmt.Printf("%s -- %v \n", t.Filed(k).Name, v.Field(k).Interface())
//	//	}
//	//}
//
//
//	/*
//		var user1 map[string]interface{}
//		json.Unmarshal([]byte(URL),&user1)
//		for k ,v := range user1{
//			fmt.Println(k," --> ",v)
//
//			}
//	*/
//	/*
//	var dat1 map[string]interface{}
//	json.Unmarshal([]byte(date1),&dat1)
//	fmt.Println("dat: ",dat1)
//	fmt.Println("\n===========json_dat[]===========")
//	for k, v := range dat1{
//		fmt.Println(k," --> ",v)
//	}
//
//}*/
///*
//func main()  {
//	a := 4
//	b := int64(a)
//var s	string
//	fmt.Println("=======")
//	ss := fmt.Sprint(s, a + int(b))
//
//	reflect.TypeOf().Kind().String()
//	//fmt.Println("dat[openid]:",reflect.TypeOf( ).Kind().String())
//	fmt.Println("+++++++")
//	fmt.Println("s: ",s)
//	fmt.Println("s: ",ss)
//}*/
////刷新Json解压后的结构体
///*
//type AAA struct {
//	Num int
//	Pot string
//	Test string
//}
//type BBB struct {
//	Num int
//	Pot string
//}
//
//func main()  {
//	str := `{"num":1,"pot":"pot","test":"abc"}`
//	str1 := `{"num":5,"pot":"tttt","txt":"wenzi"}`
//	var aaa AAA
//	json.Unmarshal([]byte(str),&aaa)
//	fmt.Println("aaa.Num: ",aaa.Num)
//	fmt.Println("aaa.pot: ",aaa.Pot)
//
//	json.Unmarshal([]byte(str1),&aaa)
//	fmt.Println("aaa.Num: ",aaa.Num)
//	fmt.Println("aaa.pot: ",aaa.Pot)
//
//
//}
//
///*
//
//import (
//	"net/http"
//
//	"fmt"
//	"io/ioutil"
//	"encoding/json"
//
//)
//
//
//type KEY struct {
//	OPENID string
//	ACCESS_TOKEN string
//	Refresh_token string
//	Expires_in string
// 	Scope string
//
//	}
//type TOKEN struct {
//	Errcode int
//	Errmsg string
//}
//
//type USERinfo struct {
//	Openid string
//	Nickname string
//	SEX int
//	Language string
//	City string
//	Province string
//	Country string
//	Headimgurl string
//}
//func check(e error)  {
//	if e != nil{
//		panic(e)
//	}
//}
////func KitHander(w http.ResponseWriter,r *http.Request)  {
////	date,err := ioutil.ReadFile("./kit_html.txt")
////	check(err)
////	date_str := string(date)
////	if r.Method == "GET"{
////		io.WriteString(w,date_str)
////	}
////}
//func UserHander(w http.ResponseWriter,r *http.Request,AppId string,Secret string) user USERinfo {
//	//AppId :="wx5391738a2c43387a"
//	//Secret :="f500b4f68f74742f6f57cae2769458b4"
//	//获取CODE数据
//	code_num := r.URL.Query().Get("code")
//
////拼接CODE到获取access_token的来链接中
//	URL := "https://api.weixin.qq.com/sns/oauth2/access_token?appid="+AppId+"&secret="+Secret+"&code="+code_num+"&grant_type=authorization_code"
//	resp, err1 := http.Get(URL)
//	check(err1)
//	defer resp.Body.Close()
//
//	//获取微信返回数据(Json结构)
//	date, err2 := ioutil.ReadAll(resp.Body)
//	check(err2)
//	//{"access_token":"7__myawMwVDYIDpC8_xcwq1RytSp0lENg8zZ09I-r8wcbvQvkkFEhABeT8Rbz0CPaLBFI7hnHUCASNtBv6UbneeQ","expires_in":7200,"refresh_token":"7_BqgBc48uj37V7zkgMMbyip8_EjJeQQFWLtwkhMpSsEp3uiTKElg8jvnPgqkaHliVPQrhCB2X4-V1I8SyL_9H3w","openid":"oSVxI0w0l1LzHrgletaH8be8VfNw","scope":"snsapi_userinfo"}
//	//解析Json数据为struct结构，获取access_token和openid
//	var key KEY
//	if err3 := json.Unmarshal([]byte(date),&key);err3 == nil{
//		fmt.Println("key.OPENID: ",key.OPENID)
//		fmt.Println("key.ACCESS_TOKEN: ",key.ACCESS_TOKEN)
//	}
//	//判断access_token是否过期
//	//https://api.weixin.qq.com/sns/auth?access_token=ACCESS_TOKEN&openid=OPENID
//	URL_token := "https://api.weixin.qq.com/sns/auth?access_token="+key.ACCESS_TOKEN +"&openid="+key.OPENID
//	resp_token, err4 := http.Get(URL_token)
//	check(err4)
//	defer resp_token.Body.Close()
//	date_token,err5 := ioutil.ReadAll(resp_token.Body)
//	check(err5)
//	var token TOKEN
//	err6 := json.Unmarshal(date_token,&token)
//	if err6 == nil{
//		fmt.Println("token.Errcode",token.Errcode)
//		fmt.Println("token.Errmsg",token.Errmsg)
//	}
//
//	if token.Errmsg != "ok" && key.Refresh_token != ""{
//		//https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=APPID&grant_type=refresh_token&refresh_token=REFRESH_TOKEN
//		URL_refresh_token := "https://api.weixin.qq.com/sns/oauth2/refresh_token?appid="+AppId+"grant_type=refresh_token&refresh_token="+key.Refresh_token
//		resp_refresh_token,err10 := http.Get(URL_refresh_token)
//		check(err10)
//		defer resp_refresh_token.Body.Close()
//		date_refresh_token,err11 := ioutil.ReadAll(resp_refresh_token.Body)
//		check(err11)
//		err12 := json.Unmarshal(date_refresh_token,&key)
//		if err12 == nil{
//			fmt.Println("key.OPENID: ",key.OPENID)
//			fmt.Println("key.ACCESS_TOKEN: ",key.ACCESS_TOKEN)
//		}
//	}
//
//	//拼接access_token和openid用于端口访问，获取用户子信息
//	URL1 := "https://api.weixin.qq.com/sns/userinfo?access_token="+key.ACCESS_TOKEN+"&openid="+key.OPENID
//	resp_user, err7 := http.Get(URL1)
//	check(err7)
//	defer resp_user.Body.Close()
//	date_user, err8 := ioutil.ReadAll(resp_user.Body)
//	check(err8)
//	var user USERinfo
//	if err9 := json.Unmarshal([]byte(date_user),&user);err9 == nil {
//		fmt.Println(user)
//		fmt.Println("user.Openid: ",user.Openid)
//		fmt.Println("user.Openid: ",user.Openid)
//	}
//	return user
//}
////func main() {
////	kit := http.NewServeMux()
////	kit.HandleFunc("/kit", KitHander)
////	USER := http.NewServeMux()
////	USER.HandleFunc("/hello",UserHander)
////	var err error
////	go func() {
////		http.ListenAndServe(":8888",kit)
////		return
////	}()
////	time.Sleep(200*time.Millisecond)
////	if err != nil{
////		fmt.Println("ListenAndServe:8888: ",err)
////		return
////	}
////
////	go func() {
////		http.ListenAndServe(":8887", USER)
////		return
////	}()
////	time.Sleep(200*time.Millisecond)
////	if err != nil{
////		fmt.Println("ListenAndServe:8887: ",err)
////		return
////	}
////	select {}
////}
//
//*/
/*
func AAA()  {
	fmt.Println("Hello!")
}
func BBB()  {
	time.Sleep(3*time.Second)
	fmt.Println("nihao")
}
func main()  {
	var i chan int
	var j chan int
	go func() {
		AAA()
		i<- 1
	}()
	go func() {
		BBB()
		j<- 2
	}()
	<-i
	<-j
}

*/
/*
type I interface {
	Get() int
	Put( int)
}
type S struct {
	i int
	}
func (p *S) Get() int {
	return p.i
	}
func (p *S) Put(v int) {
	p.i = v
	}
*/
//接口使用
/*
type My_test interface {
	print()
	Add()int
	Jian()
}
type TEST struct {
	a int
	b int
}

func (num *TEST)print()  {
	fmt.Println("a: ",num.a,"b: ",num.b)
}
func (num *TEST)Add()int   {
	return num.a+num.b
}
func main()  {
	test := TEST {1,2}
	test.print()
	a := test.Add()
	fmt.Println("a: ",a)
	fmt.Println(test.b)
}
*/


func main()  {

	print(http.HandleFunc("/ll",Function.)
	http.HandleFunc("/ni",Function.Nihaohander)
	http.ListenAndServe(":8000",nil)
	}





