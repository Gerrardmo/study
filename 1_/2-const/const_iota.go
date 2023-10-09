package main

import "fmt"

const (
	//可以在const（）中定义一个iota第一行的iota会默认为0  妹一行都会默认加一
	beijing = iota
	shanghai
	foshan    = 6
	guangzhou = iota
)

func main() {
	const length int = 10
	//length=20  会报错  常量不允许修改

	fmt.Println(length)
	fmt.Println("beijing,shanghai ,guangzhou ", beijing, shanghai, guangzhou)
}
