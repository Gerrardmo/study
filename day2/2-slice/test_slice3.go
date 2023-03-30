package main

import "fmt"

func main() {
	var numbers = make([]int, 3, 5)
	fmt.Printf("len=%d,cap=%d,value=%v", len(numbers), cap(numbers), numbers)
	numbers = append(numbers, 1)
	numbers = append(numbers, 2)
	numbers = append(numbers, 3)
	fmt.Printf("len=%d,cap=%d,value=%v", len(numbers), cap(numbers), numbers)
}
