// package main

// import "fmt"

// func main() {
// 	// 	fmt.Println("Hello \n World! \t Yes")
// 	// 	fmt.Println(`Hello \n World! \t Yes`)
// 	// 	fmt.Println(`Hello
// 	// World!		Yes`)

// 	// fmt.Printf("%v - %[1]T", `text in backs`)
// }

// ###############################

// package main

// import "fmt"

// func main() {
// 	var pi rune = 960
// 	var alpha rune = 940
// 	var omega rune = 969
// 	var bang byte = 33
// 	fmt.Printf("%c\n%c\n%c\n%c", pi, alpha, omega, bang)
// }

// // ###############################

// package main

// import "fmt"

// func main() {
// 	var smile rune = 128515
// 	fmt.Printf("%c", smile)

// }

// // ###############################

// package main

// import "fmt"

// func main() {
// 	var smile rune = '😧'
// 	var rune_A rune = 'A'
// 	var byte_A byte = 'A'

// 	fmt.Printf("%v\n", smile)
// 	fmt.Printf("%v\n", rune_A)
// 	fmt.Printf("%v\n", byte_A)

// }

// // ###############################

// package main

// import "fmt"

// func main() {
// 	message := "Hello!!"

// 	fmt.Println(message[5])
// 	fmt.Printf("%c\n", message[5])

// 	for i := 0; i < 7; i++ {
// 		fmt.Printf("%c\n", message[i])
// 	}

// }

// // ###############################

// package main

// import (
// 	"fmt"
// )

// func cypto(letter rune) string {
// 	shifered := letter + 3
// 	if shifered > 'z' {
// 		shifered = shifered - 26
// 	}
// 	return (fmt.Sprintf("%c", shifered))
// }
// func main() {
// 	var letter rune = 'h'
// 	fmt.Println(cypto(letter))
// }

// // ###############################

// package main

// import (
// 	"fmt"
// )

// func cypto(letter rune) string {
// 	shifered := letter - 3
// 	if shifered < 'a' {
// 		shifered = shifered + 26
// 	}
// 	return (fmt.Sprintf("%c", shifered))
// }
// func main() {
// 	var letter rune = 'h'
// 	fmt.Println(cypto(letter))
// }

// // ###############################

// package main

// import (
// 	"fmt"
// 	"unicode/utf8"
// )

// func main() {
// 	// AsciiQuestion := "Hello?"
// 	// fmt.Println(len(AsciiQuestion))
// 	question := "¿Cómo estás?"
// 	// fmt.Println(len(question))

// 	// fmt.Println(utf8.RuneCountInString(AsciiQuestion))
// 	fmt.Println(utf8.RuneCountInString(question))

// 	letter, size := utf8.DecodeRuneInString(question)
// 	fmt.Printf("if letter %c %v bytes", letter, size)
// }

// // ###############################

// package main

// import "fmt"

// func main() {
// 	question := "¿Cómo estás?"
// 	for i, v := range question {
// 		fmt.Printf("%v - %c\n", i, v)
// 	}

// 	for _, v := range question {
// 		fmt.Printf("%c\n", v)
// 	}
// 	// 26 символов - 26 рун и 26 байт
// }

// // ###############################

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	string1 := "My name is Dima, Dima, Dima"
// 	result := strings.Replace(string1, "Dima", "Misha", 2) // -1 все заменить
// 	fmt.Print(result)

// }

// // ###############################

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	fmt.Println("Введите строку:")
// 	reader := bufio.NewReader(os.Stdin)
// 	input, err := reader.ReadString('\n')
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(input)

// }

// ###############################
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
