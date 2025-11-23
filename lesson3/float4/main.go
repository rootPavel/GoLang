package main

import "fmt"

func main() {
	var course float32 = 84.8
	var count float32

	fmt.Print("Введите сумму которую необходимо перевести в тугрики: ")
	fmt.Scan(&count)

	fmt.Printf("Необходимо выдать: %.2f\n", transfer(count, course))

}

func transfer(money, course_2 float32) float32 {
	return money / course_2
}
