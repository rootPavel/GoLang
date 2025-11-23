// Задача:
// 		Написать программу на ЯП Go cо следующим функционалом:
// 		внутри функции main создать пеменную phoneNumber := "89995431232"; "+7(999)1232321" "8-999-321-32-43"
// 		потом в main выозвите функцию isPhoneNumber(phoneNumber)
// 		isPhoneNumber - функци, примнимает на вход строку, возвращает строку - только цифры без других символов и проверяет равно ли длинна 11?
// 		подсказка: пройти по всему номеру с помощью range;
// 		unicode.IsDigit(rune) - проверка нав цифру
// 		Пусть выводится на экран: только номер без символов и bool номер/не номер телефона

package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func isPhoneNumber(phoneNumber string) (string, bool) {
	var parsNumber string
	for _, v := range phoneNumber {
		if unicode.IsDigit(v) {
			parsNumber += string(v)
		}
	}
	len_ok := len(parsNumber) == 11
	return parsNumber, len_ok
}

func main() {
	println("Введите номер телефона:")
	phoneNumber := bufio.NewReader(os.Stdin)
	input, err := phoneNumber.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	parsNumber, valid := isPhoneNumber(input)

	fmt.Println(parsNumber, valid)

}
