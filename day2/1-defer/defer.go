package main

import "fmt"

/*
*
defer会被放到最后执行  常用来关闭资源，关闭连接
*/
func main() {
	defer fmt.Println("defer方法  ")

	fmt.Println("1")
	fmt.Println("2")

}
