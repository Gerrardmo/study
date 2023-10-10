package main

import (
	"fmt"
)

/*
*
 */
func main() {
	fmt.Println("hello golang!~")
	//四种变量声明方式
	var a int
	var b int = 1
	c := 2
	var d = 4

	fmt.Println(a, b, c, d)
	e := "你好a1 "
	fmt.Println(len(e))
	for _, i2 := range e {
		fmt.Printf("unicode %c %d \n", i2, i2)
	}

}
