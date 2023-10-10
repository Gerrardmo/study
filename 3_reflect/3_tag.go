package main

import "reflect"

type Person struct {
	Name string `info:"name" doc:"名字"`
	Sex  string `info:"sex" doc:"性别"`
}

func findTag(str interface{}) {
	//通过反射
	t := reflect.TypeOf(str).Elem()
	for i := 0; i < t.NumField(); i++ {

		println(t.Field(i).Tag.Get("doc"))
	}
}

func main() {
	var a Person
	findTag(&a)
}
