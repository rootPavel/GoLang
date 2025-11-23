package main

import "fmt"

func main() {
	// var person map[string]string
	// person = make(map[string]string)
	person := map[string]string{"Name": "Serega", "Email": "Sererega@mail.ru"}
	person["Age"] = "40"
	person["Name"] = "Sergey"
	delete(person, "Age")
	fmt.Println(person)
	fmt.Println(person["Email"])
}
