package main

import "fmt"

// Структура
type Rover struct {
	lat  float64
	long float64
}

func main() {
	// Структура
	var curosity struct {
		lat  float64
		long float64
	}
	curosity.lat = 5.123
	curosity.long = -9.888

	fmt.Printf("[%v:%v]\n", curosity.lat, curosity.long)

	// Экземпляр структуры Rover
	var curosity2 Rover
	curosity2.lat = 7.123
	curosity2.long = -10.888

	fmt.Printf("[%v:%v]\n", curosity2.lat, curosity2.long)
}
