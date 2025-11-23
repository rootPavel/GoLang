package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		// brake -закончить цикл
		// continue - перейти к следующей итерации
		if i%2 != 0 || i == 0 {
			continue
		}
		fmt.Println(i)
	}
}
