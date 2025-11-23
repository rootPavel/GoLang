package main

import "fmt"

type Wheels struct {
	Brand  string
	Radius int
}

func (w Wheels) Run() {
	fmt.Println("Wheels are going")
}

type Vehicle struct {
	Brand string // мерседес, ауди, бмв
	Color string
	// Wheels Wheels
	Wheels
}

func main() {

	myCar := Vehicle{
		Brand: "Lada",
		Color: "Red",
		Wheels: Wheels{
			Brand:  "Icon",
			Radius: 18,
		},
	}

	fmt.Println(myCar)
	myCar.Wheels.Run()
	myCar.Run()
}
