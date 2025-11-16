// package main

// import "fmt"

// func Cube(x int) (res int) {
// 	res = x * x * x
// 	return
// }

// func main() {
// 	result := Cube(2)
// 	fmt.Println(result)
// }

// //############################

// package main

// import "fmt"

// func Sum(x ...int) (res int) {
// 	for _, v := range x {
// 		res += v
// 	}
// 	return
// }

// func main() {
// 	result := Sum(1, 2, 3, 4, 5, 6)
// 	fmt.Println(result)
// }

// //############################

// package main

// import "fmt"

// func Sum(x ...int) (res int) {
// 	for _, v := range x {
// 		res += v
// 	}
// 	return
// }

// func main() {
// 	result := Sum(1, 2, 3, 4, 5, 6)
// 	fmt.Println(result)
// }

// //############################

// package main

// import "fmt"

// func Say(animal string) string {
// 	if animal == "dog" {
// 		return "wuf"
// 	} else if animal == "cat" {
// 		return "Myau"
// 	} else {
// 		return "nah"
// 	}
// }

// func PrintVoice(animal string, how func(string) string) {
// 	fmt.Println(how(animal))
// }

// func main() {
// 	PrintVoice("dog", Say)
// 	PrintVoice("cat", Say)
// }

// //############################

// package main

// import "fmt"

// var name string

// func init() {
// 	name = "Sergey"
// }

// func init() {
// 	name = "Alex"
// }

// func main() {
// 	fmt.Println(name)
// }

// //############################

// package main

// import "fmt"

// func main() {
// 	fmt.Println("1st output")
// 	defer fmt.Println("2st output")
// 	defer fmt.Println("3st output")
// 	defer fmt.Println("4st output")
// }

// //############################

// package main

// func foo() {
// 	panic("Это паника функии foo()")
// }

// func main() {
// 	foo()
// }

// //############################

// package main

// import "fmt"

// func foo() {
// 	defer func() {
// 		r := recover()
// 		if r != nil {
// 			fmt.Println("Я поймал панику!")
// 		}
// 	}()
// 	panic("Это паника функии foo()")
// }

// func main() {
// 	foo()
// }

//############################
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
