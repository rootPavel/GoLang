package main

import "fmt"

func main() {
	var name, surname string
	var age int
	name = "Pavel"
	age = 37

	// fmt.Println("My name is", name, "San")
	fmt.Printf("My name is %s San. I am %v years old.", name, age)
	my_v := fmt.Sprintf(" Hello, i am %s", name)
	fmt.Println(my_v)

	fmt.Print("Enter ur name: ")
	fmt.Scan(&name)
	fmt.Printf("Hello, %s\n", name)

	/*
		Важно:
		fmt.Scan() - считывает ввод до первого пробела
		Необходимо изучить - bufio.NewReader() / os
	*/

	fmt.Print("Enter ur name: ")
	fmt.Scan(&name)
	fmt.Print("Enter ur surname: ")
	fmt.Scan(&surname)
	fmt.Print("Enter ur age: ")
	fmt.Scan(&age)
	fmt.Printf("Hello, %s %s, who is %d years \n", name, surname, age)
}
