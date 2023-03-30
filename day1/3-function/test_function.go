package main

import "fmt"

func foo(a int, b int) int {
	return a + b
}
func foo1(a int, b int) (int, int) {
	return a + b, a * b
}
func foo2(a int, b int) (c int, d int) {
	c = a + b
	d = a * b
	return
}
func main() {
	fmt.Println(foo(1, 2))
	fmt.Println(foo1(4, 2))
	fmt.Println(foo2(5, 6))
}
