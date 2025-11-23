package main

import "fmt"

func main() {
	var age int
	fmt.Print("Введите возраст: ")
	fmt.Scan(&age)
	if (age <= 7) || (age >= 70) {
		fmt.Println("Проезд бесплатный")
	} else if (age > 65) && (age < 70) {
		fmt.Println("Необходимо оплатить за проезд 35 рублей")
	} else {
		fmt.Println("Необходимо оплатить за проезд 70 рублей")
	}
}
