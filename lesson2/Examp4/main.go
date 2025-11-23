package main

import "fmt"

func main() {
	var a, b int
	a = 1
	b = 2

	if c := 4; c > b && c > a {
		fmt.Println(c)
	}
	// fmt.Println(c) - не сработает, так как переменная существует только в рамках if
}
