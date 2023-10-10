package main

import (
	"fmt"
	"time"
)

func main() {
	go func(a int) {
		a++
		fmt.Println("a:", 1)
		time.Sleep(1 * time.Second)
	}(1)
	i := 0
	for {
		i++
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
}
