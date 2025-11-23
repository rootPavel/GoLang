package main

import "fmt"

// Структура
type Rover struct {
	lat  float64
	long float64
}

func main() {
	curosity := Rover{lat: 5.123, long: -9.888}
	fmt.Printf("[%v:%v]\n", curosity.lat, curosity.long)

	curosity.lat += 1.0
	fmt.Printf("%+v\n", curosity)

	curosity2 := Rover{
		lat:  7.123,
		long: -10.888,
	}

	fmt.Printf("[%v:%v]\n", curosity2.lat, curosity2.long)

	curosity3 := Rover{-10.888, 7.123}

	fmt.Printf("%+v\n", curosity3)
}
