package main

import "fmt"

func main() {
	var a string
	a = "sss"
	//pair<statictype:string,value:"sss">

	//pair<statictype:string,calue:"sss">
	var allType interface{}

	str, v := allType.(string) //类型断言  未知类型转成已知类型
	allType = a
	fmt.Println(str, v)
}
