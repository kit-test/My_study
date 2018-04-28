package main

import (


	"fmt"
	"encoding/json"
)

//type accessTokenArgs struct {
//	GrantType string `json:"grant_type"`
//	Appid     string `json:"appid"`
//	Secret    string `json:"secret"`
//}
//const grantType = "client_credential"
//
//func main()  {
//	//AppId :="wx5391738a2c43387a"
//	//Secret :="f500b4f68f74742f6f57cae2769458b4"
//	//My_function.UserHander(AppId,Secret)
//
//	//args := &accessTokenArgs{
//	//	GrantType: grantType,
//	//	Appid:     "123456",
//	//	Secret:    "secret_111",
//	//}
//
//
//	v := url.Values{"Appid": {"123456"}, "Secret": {"secret_111"}}
//
//	var urll string
//
//	urll = "HTTP://www.kit.com"
//
//	urll += "?" + v.Encode()
//

//	fmt.Println(urll)
//}
//type User struct {//{"user_id":1,"user_name":"tony"}
//	UserId   int    `json:"user_id" bson:"user_id"`
//	UserName string `json:"user_name" bson:"user_name"`
//}

type User struct {//{"UserId":1,"UserName":"tony"}
	UserId   int
	UserName string
}
func main()  {
	u := &User{UserId: 1, UserName: "tony"}
	j, _ := json.Marshal(u)
	fmt.Println(string(j))
}