package main

import (
	"fmt"
	printingfunctions "lesson1/PrintingFunctions"
	"reflect"
)

// printingfunctions "lesson1/PrintingFunctions"

var v_cube = printingfunctions.Cube(3)

func main() {
	// print("Hello! Pavel")
	// fmt.Println("Hello from fmt")
	// fmt.Println("Hello it is Println")
	/*
		printingfunctions.PrintMyAge()
		printingfunctions.PrintMyName()
		printingfunctions.PrintMyCity()
		printingfunctions.PrintMyHome()
	*/
	var name, surname string
	name, surname = "Pavel", "K"
	var age = 29
	heigh := 1.75
	heigh = 1
	const city = "Moscow"

	fmt.Println("Name - "+name, reflect.TypeOf(name))
	fmt.Println("Surname -", surname, reflect.TypeOf(surname))
	fmt.Println("Age -", age, reflect.TypeOf(age))
	fmt.Println("Heigh -", heigh, reflect.TypeOf(heigh))
	fmt.Println("City -", city, reflect.TypeOf(city))

	var v_square = printingfunctions.Square(5)

	fmt.Println("Scuare =", v_square)
	fmt.Println("Cube =", v_cube)
}
