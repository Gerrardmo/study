package main

import "fmt"

func printSlice(mySlice []int) {
	//引用传递，传过来的是一个指针，改变了原来的值
	for k, v := range mySlice {
		fmt.Println(k, ":", v)
	}
	mySlice[0] = 100 //原切片的值会被改变

}

func main() {
	//动态数组  切片
	mySlice := []int{1, 2, 3, 4}
	printSlice(mySlice)
	fmt.Println(mySlice)
}
