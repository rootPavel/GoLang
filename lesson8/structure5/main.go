package main

import (
	"encoding/json"
	"fmt"
)

type Pet struct {
	Name string `json:"FullName"`
	Age  int    `json:"FullAge"`
}

func main() {

	Tema := Pet{Name: "Tema", Age: 12}
	fmt.Printf("%+v\n", Tema)

	// маршал не может увидеть неэкспортируемый атрибут
	bytes, err := json.Marshal(Tema)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bytes)
	fmt.Println(string(bytes))
}
