package main

import "fmt"

func main() {
	var cat int = 1
	//&符取地址
	fmt.Println(&cat)

	//地址的指针取值为原值
	ptr := &cat
	fmt.Println(*ptr)

	var a int = 10
	nopointer(a)
	fmt.Println(a)
	pointer(&a)
	fmt.Println(a)

}
func nopointer(a int) {
	a = 20
}

func pointer(a *int) {
	*a = 20
}
