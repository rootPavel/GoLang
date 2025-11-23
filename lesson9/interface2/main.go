package main

import "fmt"

type Mover interface{}

// JSONfile := map[string]interface{}
func main() {
	var a Mover = 1
	fmt.Println(a)
	a = "asd"
	fmt.Println(a)
	a = 1.5
	fmt.Println(a)
}
