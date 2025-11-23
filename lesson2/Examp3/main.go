package main

import "fmt"

func main() {
	var a, b int
	a = 1
	b = 2
	incB := func() bool {
		b += 1
		return true
	}
	if (a == 1) || incB() {
		fmt.Println("All good")
	} else if b == 2 {
		fmt.Println("Bad value")
	}
	fmt.Println(a, b)
}
