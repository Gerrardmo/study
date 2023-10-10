package main

import "fmt"

// 类首字母大写，其他模块可调用
type Hero struct {
	Name  string
	Level int
	Hp    float32
	Mp    float32
}

func (i *Hero) GetName() {
	fmt.Println("Name:", i.Name)
}
func (i *Hero) SetName(newName string) {
	i.Name = newName
}
func (i *Hero) Show() {
	fmt.Println("Name:", i.Name)
	fmt.Println("Level:", i.Level)
	fmt.Println("Hp:", i.Hp)
	fmt.Println("Mp:", i.Mp)
}
func main() {
	Joes := Hero{
		Name:  "Hero",
		Level: 99,
		Hp:    12.1,
		Mp:    100,
	}

	Joes.Show()
	Joes.SetName("Tawns")
	Joes.Show()

}
