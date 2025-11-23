package main

import (
	"encoding/json"
	"fmt"
)

type Pet struct {
	// неэкспортируемые атрибуты
	name string
	age  int
}

func main() {

	Tema := Pet{name: "Tema", age: 12}
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
