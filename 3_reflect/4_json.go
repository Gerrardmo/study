package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Tittle string  `json:"tittle"`
	Year   int     `json:"year"`
	Price  float32 `json:"price"`
}

func main() {
	movie1 := Movie{
		Tittle: "龙族",
		Year:   2008,
		Price:  100,
	}
	marshal, err := json.Marshal(movie1)
	if err != nil {
		fmt.Println("err" +
			"!")
	}
	fmt.Printf("jsonstr %s\n", marshal)
}
