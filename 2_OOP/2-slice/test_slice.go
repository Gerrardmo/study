package main

import "fmt"

func FmtArray(myarray [4]int) {
	fmt.Println("++++++++++")
	for i, i2 := range myarray {
		fmt.Println(i, ":", i2)
	}
}

func main() {
	//固定长度的数组
	var myArray1 [10]int
	myArray2 := [10]int{1, 2, 3, 4}
	myArray4 := [4]int{1, 2, 3, 4}
	for i := 0; i < len(myArray1); i++ {
		myArray1[i] = i
		fmt.Println(myArray1[i])
	}
	//for _, v := range myArray2 {
	//	fmt.Println(v)
	//}
	for index, value := range myArray2 {
		fmt.Println(index, ":", value)
	}
	FmtArray(myArray4)
	//	FmtArray(myArray2)
	fmt.Printf("myArray1 type:%T \n", myArray1)
	fmt.Printf("myArray2 type:%T \n", myArray2)
}
