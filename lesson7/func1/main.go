package main

import "fmt"

func Cube(x int) (res int) {
	res = x * x * x
	return
}

func main() {
	result := Cube(2)
	fmt.Println(result)
}
