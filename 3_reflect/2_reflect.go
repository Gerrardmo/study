package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 100

	fmt.Println(reflect.ValueOf(a), reflect.TypeOf(a))
}
