package main

import "fmt"

func main() {
	//slice1 := []int{1, 2, 3}  //len=3 ,slice=[1 2 3]

	//var slice1 []int
	//slice1 = make([]int, 3) //开辟三个空间  默认值0 //len=3 ,slice=[0 0 0]

	slice1 := make([]int, 3) //len=3 ,slice=[0 0 0]

	fmt.Printf("len=%d ,slice=%v \n", len(slice1), slice1)
}
