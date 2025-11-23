package main

import "fmt"

func main() {
	a := 1.0 / 3
	fmt.Println(a + a + a)

	b := 0.1
	b += 0.2

	fmt.Println(b)
	fmt.Printf("%.60f\n", 0.1)
	fmt.Printf("%.60f\n", 0.2)

	fmt.Println(b == 0.3)
	fmt.Println((b - 0.3) < 0.00001)
}
