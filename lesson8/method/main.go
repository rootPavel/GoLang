package main

import "fmt"

type celsius float64

// Метод
func (c celsius) MFromCelsiusToKelvin() {
	c += 273.15
	fmt.Println(c)
}

type kelvin float64

// Функция
func FromCelsiusToKelvin(c celsius) kelvin {
	// 273.15
	c += 273.15
	return kelvin(c)
}

func main() {

	const degress = 20
	// Свой тип данных celsius
	var temperature celsius = degress
	temperature += 10

	fmt.Println(temperature)

	var fl64 float64 = 10.0
	temperature += celsius(fl64)

	fmt.Println(FromCelsiusToKelvin(temperature))

	temperature.MFromCelsiusToKelvin()
}
