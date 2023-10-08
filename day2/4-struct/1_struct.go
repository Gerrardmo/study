package main

import "fmt"

// 生命一个数据类型
type myint int

type Book struct {
	tittle string
	auth   string
}

func changebook(book *Book) {
	book.auth = "ggg"
}
func changebook1(book Book) {
	book.auth = "ggg"
}
func main() {
	//var a myint = 10
	//fmt.Println("a=", a)
	//fmt.Printf("type of %T:", a)
	//
	var book1 Book
	book1.tittle = "golang"
	book1.auth = "goher"
	changebook1(book1)
	fmt.Printf("%v\n", book1)
	changebook(&book1)
	fmt.Printf("%v\n", book1)

}
