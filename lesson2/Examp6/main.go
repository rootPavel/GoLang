package main

import (
	"fmt"
)

func main() {
	v := 0
	i := 0
	for i < 10 {
		fmt.Println(i)
		v++
		fmt.Println(v)
		i++
	}
}
