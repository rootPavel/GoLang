package main

import "fmt"

func main() {
	i := 0
	var count int
	for {
		fmt.Println("Введите любое число")
		fmt.Scan(&count)
		i = i + count
		if count == 0 {
			fmt.Printf("Сумма введенных чисел = %d", i)
			break
		}
	}
}
