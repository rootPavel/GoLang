// package main

// import "fmt"

//func main() {
// name := "Pavel"
// age := 37

// // fmt.Println("My name is", name, "San")
// fmt.Printf("My name is %s San. I am %v years old.", name, age)
// my_v := fmt.Sprintf(" Hello, i am %s", name)
// fmt.Println(my_v)

// var name string
// fmt.Print("Enter ur name: ")
// fmt.Scan(&name)
// fmt.Printf("Hello, %s\n", name)

/*
	Важно:
	fmt.Scan() - считывает ввод до первого пробела
	Необходимо изучить - bufio.NewReader() / os
*/

// 	var name, surname string
// 	var age int
// 	fmt.Print("Enter ur name: ")
// 	fmt.Scan(&name)
// 	fmt.Print("Enter ur surname: ")
// 	fmt.Scan(&surname)
// 	fmt.Print("Enter ur age: ")
// 	fmt.Scan(&age)
// 	fmt.Printf("Hello, %s %s, who is %d years \n", name, surname, age)
// }

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
