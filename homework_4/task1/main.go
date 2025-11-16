/*
Реализуйте функцию divide(a, b float64) (float64, error), которая делит a на b.
Если b равно нулю, возвращайте ошибку.
В основной программе вызовите эту функцию и обработайте возможную ошибку.
*/

package main

import (
	"fmt"
	"strconv"
)

func enterFloat(text string) (float64, error) {
	fmt.Println(text)
	var a string
	fmt.Scan(&a)

	return strconv.ParseFloat(a, 64)
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Деление на ноль")
	} else {
		return a / b, nil
	}
}

func main() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("Вы ввели не число! Попробуйте еще раз")
		}
	}()
	a, err := enterFloat("Введите число a: ")
	if err != nil {
		panic(err)
	}
	b, err := enterFloat("Введите число b: ")
	if err != nil {
		panic(err)
	}
	result, err := divide(a, b)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Результат:", result)
	}
}
