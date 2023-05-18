package main

import "fmt"

func main() {
	cityMap := map[string]string{
		"China": "北京",
		"japen": "东京",
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

}
