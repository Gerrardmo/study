package main

import "fmt"

func main() {
	cityMap := map[string]string{
		"China": "北京",
		"japen": "东京",
		"koren": "韩国",
	}
	for key, value := range cityMap {
		fmt.Println("key=", key)
		fmt.Println("value", value)
	}
	//删除
	delete(cityMap, "japen")
	for key, value := range cityMap {
		fmt.Println("key=", key)
		fmt.Println("value", value)
	}
	//修改
	cityMap["koren"] = "首尔"
	fmt.Println("===")
	for key, value := range cityMap {

		fmt.Println("key=", key)
		fmt.Println("value", value)
	}
}
