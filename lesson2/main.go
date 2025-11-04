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

// ##################################

// package main

// import "fmt"

// func main() {
// 	var age int
// 	fmt.Print("Введите возраст: ")
// 	fmt.Scan(&age)
// 	if (age <= 7) || (age >= 70) {
// 		fmt.Println("Проезд бесплатный")
// 	} else if (age > 65) && (age < 70) {
// 		fmt.Println("Необходимо оплатить за проезд 35 рублей")
// 	} else {
// 		fmt.Println("Необходимо оплатить за проезд 70 рублей")
// 	}
// }

// ##################################
// package main

// import "fmt"

// func main() {
// 	var a, b int
// 	a = 1
// 	b = 2
// 	incB := func() bool {
// 		b += 1
// 		return true
// 	}
// 	if (a == 1) || incB() {
// 		fmt.Println("All good")
// 	} else if b == 2 {
// 		fmt.Println("Bad value")
// 	}
// 	fmt.Println(a, b)
// }

// ##################################
// package main

// import "fmt"

// func main() {
// 	var a, b int
// 	a = 1
// 	b = 2

// 	if c := 4; c > b && c > a {
// 		fmt.Println(c)
// 	}
// 	// fmt.Println(c) - не сработает, так как переменная существует только в рамках if
// }

// ####################################
// package main

// import "fmt"

// func main() {
// 	var a, b int
// 	a = 10
// 	b = 3
// 	switch {
// 	case a == 1:
// 		fmt.Println("1")
// 	case b == 3:
// 		fmt.Println("2")
// 		fallthrough
// 	case a == 3:
// 		fmt.Println("3 or 4")
// 	default:
// 		fmt.Println("Default case")
// 	}
// }

// #########################

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	v := 0
// 	i := 0
// 	for i < 10 {
// 		fmt.Println(i)
// 		v++
// 		fmt.Println(v)
// 		i++
// 	}
// }

// #############################

// package main

// import "fmt"

// func main() {
// 	v := 0
// 	for i := 0; i < 10; i++ {
// 		// brake -закончить цикл
// 		// continue - перейти к следующей итерации
// 		v++
// 		fmt.Println(v)
// 		if v >= 5 {
// 			break
// 		}
// 	}
// }

// ############################

// package main

// import "fmt"

// func main() {
// 	for i := 0; i < 10; i++ {
// 		// brake -закончить цикл
// 		// continue - перейти к следующей итерации
// 		if i%2 != 0 || i == 0 {
// 			continue
// 		}
// 		fmt.Println(i)
// 	}
// }

// ##############################

package main

import "fmt"

func main() {
	i := 0
	var count int
	for {
		fmt.Println("Введите любое число")
		fmt.Scan(&count)
		i = i + count
		if count == 0 {
			fmt.Printf("Сумма введенных чисел = %d", i)
			break
		}
	}
}
