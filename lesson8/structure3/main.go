package main

import (
	"encoding/json"
	"fmt"
)

type Pet struct {
	Name string
	Age  int
}

func main() {
	// mySlice := []Pet{
	// 	Pet{Name: "Tema", Age: 12},
	// 	Pet{Name: "Yo", Age: 14},
	// }
	// mySlice := []Pet{
	// 	{Name: "Tema", Age: 12},
	// 	{Name: "Yo", Age: 14},
	// }
	// fmt.Println(mySlice)
	// fmt.Println(mySlice[0].Name)

	Tema := Pet{Name: "Tema", Age: 12}
	fmt.Printf("%+v\n", Tema)

	// json

	bytes, err := json.Marshal(Tema)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bytes)
	fmt.Println(string(bytes))
}
