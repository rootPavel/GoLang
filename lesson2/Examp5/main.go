package main

import "fmt"

func main() {
	var a, b int
	a = 10
	b = 3
	switch {
	case a == 1:
		fmt.Println("1")
	case b == 3:
		fmt.Println("2")
		fallthrough
	case a == 3:
		fmt.Println("3 or 4")
	default:
		fmt.Println("Default case")
	}
}
