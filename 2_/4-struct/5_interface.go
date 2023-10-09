package main

import "fmt"

/*
*
空接口使用
*/
type inf interface {
}

func print(date inf) {

	fmt.Printf("%v %T\n", date, date)
}
func main() {
	a := "sss"
	b := 1
	c := []int{1, 2, 3}
	print(a)
	print(b)
	print(c)

}
