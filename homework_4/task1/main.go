/*
Реализуйте функцию divide(a, b float64) (float64, error), которая делит a на b.
Если b равно нулю, возвращайте ошибку.
В основной программе вызовите эту функцию и обработайте возможную ошибку.
*/

package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

var (
	stdin  io.Reader = os.Stdin
	stdout io.Writer = os.Stdout
)

func enterFloat(text string) (float64, error) {
	fmt.Fprintln(stdout, text)
	var input string
	fmt.Fscan(stdin, &input)
	value, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Fprintln(stdout, "Вы ввели не число! Попробуйте еще раз")
		return 0, err
	}
	return value, err
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("вы пытаетесь делить на ноль!\nВ итоге получите ноль, оно вам надо?")
	}
	return a / b, nil
}

func main() {
	fmt.Fprintln(stdout, "Эта программа умеет делить число a на число b!")

	a, err := enterFloat("Введите число a: ")
	if err != nil {
		fmt.Fprintln(stdout, "Ошибка:", err)
	}
	b, err := enterFloat("Введите число b: ")
	if err != nil {
		fmt.Fprintln(stdout, "Ошибка:", err)
	}

	result, err := divide(a, b)
	if err != nil {
		fmt.Fprintln(stdout, "Ошибка:", err)
	} else {
		fmt.Fprintf(stdout, "Результат: %.2f", result)
	}
}
