package main

import "fmt"

func main() {
	a := 1.0 / 3
	fmt.Println(a)
	fmt.Printf("%v\n", a)
	fmt.Printf("%f\n", a)
	fmt.Printf("%.2f\n", a) //.2 - после запятой 2 числа
	fmt.Printf("%08.2f\n", a)
}
