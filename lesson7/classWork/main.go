// Написать программу на ЯП Go с следующей логикой:
// 	Принять на вход (bufio.NewReader) от пользователя Имя, фамилию полностью. $ Александр Сергеев
// 	Реализовать с помощью функции! Александр Сергеев -> Александр С.
// 	Учесть: АЛЕКСАНДР СЕРГЕЕВ, александр сергеев

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func converFirstLastName(firstLastName string) string {
	listFirstLastName := strings.Fields(firstLastName)

	firstName := upOneReune(listFirstLastName[0])
	lastName := upOneReune(listFirstLastName[1])

	cropLastName := formatFirstName(lastName)

	return firstName + " " + cropLastName
}

func upOneReune(word string) string {
	runes := []rune(strings.ToLower(word))

	if len(runes) > 0 {
		runes[0] = unicode.ToUpper(runes[0])
	}

	return string(runes)
}

func formatFirstName(formatFirstName string) string {
	var cropLastName string

	if len(formatFirstName) > 0 {
		cropLastName = string(formatFirstName[0]) + "."
	}

	return cropLastName
}

func main() {
	fmt.Println("Введите Имя и Фамилию полностью:")
	firstLastName := bufio.NewReader(os.Stdin)
	rawfirstLastName, err := firstLastName.ReadString('\n')

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(converFirstLastName(rawfirstLastName))
}
