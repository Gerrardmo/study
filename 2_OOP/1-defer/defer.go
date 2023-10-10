package main

import "fmt"

/*
*
defer会被放到最后执行  常用来关闭资源，关闭连接
先return 后defer
defer  先进后出，最先defer的被压入栈底，最后执行
*/
func main() {

	defer fmt.Println("defer方法 1 ")
	defer fmt.Println("defer方法 2 ")
	defer fmt.Println("defer方法 3 ")
	fmt.Println("1")
	fmt.Println("2")

}
