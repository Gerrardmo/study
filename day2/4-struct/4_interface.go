package main

import "fmt"

// 本质是一个指针
type AnimalInf interface {
	Sleep()
	GetColor() string
	GetType() string
}

// 具体的类
type Cat struct {
	Color string
}

func (cat Cat) Sleep() {
	fmt.Println("cat is sleep")
}
func (cat *Cat) GetColor() string {
	return cat.Color
}
func (cat *Cat) GetType() string {
	return "Cat"
}

type Dog struct {
	Color string
}

func (dog Dog) Sleep() {
	fmt.Println("dog is sleep")
}
func (dog *Dog) GetColor() string {
	return dog.Color
}
func (dog *Dog) GetType() string {
	return "Dog"
}

// 多态  根据传入类型决定调用什么方法
func ShowSleep(animal AnimalInf) {
	animal.Sleep()
}
func main() {
	//var animal AnimalInf //接口数据类型  父类指针
	//调用cat的sleep方法
	//animal = &Cat{"green"}
	//animal.Sleep()
	//
	//animal = &Dog{"white "}
	//animal.Sleep()
	cat := Cat{"red"}
	dog := Dog{"white"}
	ShowSleep(&cat)
	ShowSleep(&dog)
}
