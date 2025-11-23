package main

import "fmt"

func Sum(x ...int) (res int) {
	for _, v := range x {
		res += v
	}
	return
}

func main() {
	result := Sum(1, 2, 3, 4, 5, 6)
	fmt.Println(result)
}
