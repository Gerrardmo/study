package main

import "fmt"

type Human struct {
	Name string
	sex  string
}

func (hero Human) Eat() {
	fmt.Println("hero eat")
}

func (hero Human) Move() {
	fmt.Println("hero Walk")
}
func (hero Human) Talk() {
	fmt.Println("hero Talk")
}

type Surperman struct {
	Human
	Level int
}

// 重写父类的walk方法
func (superman Surperman) Move() {
	fmt.Println("hero Fly")
}

func main() {
	h := Human{Name: "zhangsan", sex: "man"}

	h.Eat()

	s := Surperman{
		Human: Human{
			Name: "lisi",
			sex:  "man",
		},
		Level: 1,
	}

	s1 := Surperman{
		Human: h,
		Level: 0,
	}

	s.Move() //fly  调用的是重写的move
	s.Talk()
	s1.Move()
}
