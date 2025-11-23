package main

import "fmt"

func main() {
	v := 0
	for i := 0; i < 10; i++ {
		// brake -закончить цикл
		// continue - перейти к следующей итерации
		v++
		fmt.Println(v)
		if v >= 5 {
			break
		}
	}
}
