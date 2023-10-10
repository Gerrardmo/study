package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go func() {
		defer fmt.Println("gorou1 end")
		fmt.Println("gorou1 star!")
		i := 111
		time.Sleep(1 * time.Second)
		c <- i
		//time.Sleep(5 * time.Second)
	}()

	num := <-c
	fmt.Println(num)
	fmt.Println("main end")
}
