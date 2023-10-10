package main

import "fmt"

func main() {

	//声明mymap1是map类型 key是string类型   value是string类型
	var myMap1 map[string]string
	if myMap1 == nil {
		//
		fmt.Printf("mymap1是一个空map\n")
	}
	//在使用map之前先要给他分配空间
	myMap1 = make(map[string]string, 10)
	myMap1["one"] = "java"
	myMap1["two"] = "golang"
	fmt.Println(myMap1)

	//第二种声明方式
	myMap2 := make(map[string]string)
	myMap2["A"] = "C++"
	myMap2["B"] = "c#"
	fmt.Println(myMap2)
}
