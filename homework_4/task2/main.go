/*
Функция высшего порядка: передача функции как аргумента
Создайте функцию applyOperation(a, b int, op func(int, int) int) int, которая применяет переданную функцию op к числам a и b.

Создайте несколько функций-операций: сложение, вычитание, умножение.
В основной программе вызовите applyOperation с разными операциями и выведите результаты.
*/

package main

import "fmt"

func applyOperation(a, b int, op func(int, int) int) int {
	return op(a, b)
}

func sumInteger(a, b int) int {
	return a + b
}

func minusInteger(a, b int) int {
	return a - b
}

func multipyInteger(a, b int) int {
	return a * b
}

func main() {
	a := 13
	b := 5
	fmt.Printf("Складываем %d и %d, получаем %d\n", a, b, applyOperation(a, b, sumInteger))
	fmt.Printf("Отнимает %d от %d, получаем %d\n", a, b, applyOperation(a, b, minusInteger))
	fmt.Printf("Умножаем %d на %d, получаем %d\n", a, b, applyOperation(a, b, multipyInteger))

}
