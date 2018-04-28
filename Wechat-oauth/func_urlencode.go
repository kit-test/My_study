package main

import (
	"net/url"
	"fmt"
)

func main()  {
	urlStr := `http://localhost:8888/hello`
	l, err := url.ParseQuery(urlStr)
	fmt.Println("1: ",l, err)
	l2, err2 := url.ParseRequestURI(urlStr)
	fmt.Println("2: ",l2, err2)

	l3, err3 := url.Parse(urlStr)
	fmt.Println("3: ",l3, err3)
	/*fmt.Println(l3.Path)
	fmt.Println(l3.RawQuery)
	fmt.Println(l3.Query())*/
	fmt.Println("4: ",l3.Query().Encode())

	fmt.Println("5: ",l3.RequestURI())
	//fmt.Printf("Hello World! version : %s", rt.Version())
}