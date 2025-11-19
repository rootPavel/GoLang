// package main

// import "fmt"

// type celsius float64

// // Метод
// func (c celsius) MFromCelsiusToKelvin() {
// 	c += 273.15
// 	fmt.Println(c)
// }

// type kelvin float64

// // Функция
// func FromCelsiusToKelvin(c celsius) kelvin {
// 	// 273.15
// 	c += 273.15
// 	return kelvin(c)
// }

// func main() {

// 	const degress = 20
// 	// Свой тип данных celsius
// 	var temperature celsius = degress
// 	temperature += 10

// 	fmt.Println(temperature)

// 	var fl64 float64 = 10.0
// 	temperature += celsius(fl64)

// 	fmt.Println(FromCelsiusToKelvin(temperature))

// 	temperature.MFromCelsiusToKelvin()
// }

// // #############################################

// package main

// import "fmt"

// type cat string
// type dog string

// func (c cat) Speak() {
// 	fmt.Printf("Кошка %s говорит 'Мяу'\n", c)
// }

// func (d dog) Speak() {
// 	fmt.Printf("Собака %s говорит 'Гав'\n", d)
// }

// func main() {
// 	var Musya cat = "Муся"
// 	var Sharik dog = "Шарик"

// 	Musya.Speak()
// 	Sharik.Speak()

// }

// // #############################################

// package main

// import "fmt"

// // Структура
// type Rover struct {
// 	lat  float64
// 	long float64
// }

// func main() {
// 	// Структура
// 	var curosity struct {
// 		lat  float64
// 		long float64
// 	}
// 	curosity.lat = 5.123
// 	curosity.long = -9.888

// 	fmt.Printf("[%v:%v]\n", curosity.lat, curosity.long)

// 	// Экземпляр структуры Rover
// 	var curosity2 Rover
// 	curosity2.lat = 7.123
// 	curosity2.long = -10.888

// 	fmt.Printf("[%v:%v]\n", curosity2.lat, curosity2.long)
// }

// // #############################################

// package main

// import "fmt"

// // Структура
// type Rover struct {
// 	lat  float64
// 	long float64
// }

// func main() {
// 	curosity := Rover{lat: 5.123, long: -9.888}
// 	fmt.Printf("[%v:%v]\n", curosity.lat, curosity.long)

// 	curosity.lat += 1.0
// 	fmt.Printf("%+v\n", curosity)

// 	curosity2 := Rover{
// 		lat:  7.123,
// 		long: -10.888,
// 	}

// 	fmt.Printf("[%v:%v]\n", curosity2.lat, curosity2.long)

// 	curosity3 := Rover{-10.888, 7.123}

// 	fmt.Printf("%+v\n", curosity3)
// }

// // #############################################

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type Pet struct {
// 	Name string
// 	Age  int
// }

// func main() {
// 	// mySlice := []Pet{
// 	// 	Pet{Name: "Tema", Age: 12},
// 	// 	Pet{Name: "Yo", Age: 14},
// 	// }
// 	// mySlice := []Pet{
// 	// 	{Name: "Tema", Age: 12},
// 	// 	{Name: "Yo", Age: 14},
// 	// }
// 	// fmt.Println(mySlice)
// 	// fmt.Println(mySlice[0].Name)

// 	Tema := Pet{Name: "Tema", Age: 12}
// 	fmt.Printf("%+v\n", Tema)

// 	// json

// 	bytes, err := json.Marshal(Tema)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(bytes)
// 	fmt.Println(string(bytes))
// }

// // #############################################

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type Pet struct {
// 	// неэкспортируемые атрибуты
// 	name string
// 	age  int
// }

// func main() {

// 	Tema := Pet{name: "Tema", age: 12}
// 	fmt.Printf("%+v\n", Tema)

// 	// маршал не может увидеть неэкспортируемый атрибут
// 	bytes, err := json.Marshal(Tema)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(bytes)
// 	fmt.Println(string(bytes))
// }

// // #############################################

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type Pet struct {
// 	Name string `json:"FullName"`
// 	Age  int    `json:"FullAge"`
// }

// func main() {

// 	Tema := Pet{Name: "Tema", Age: 12}
// 	fmt.Printf("%+v\n", Tema)

// 	// маршал не может увидеть неэкспортируемый атрибут
// 	bytes, err := json.Marshal(Tema)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(bytes)
// 	fmt.Println(string(bytes))
// }

// // #############################################

// package main

// import "fmt"

// type Wheels struct {
// 	Brand  string
// 	Radius int
// }

// func (w Wheels) Run() {
// 	fmt.Println("Wheels are going")
// }

// type Vehicle struct {
// 	Brand string // мерседес, ауди, бмв
// 	Color string
// 	// Wheels Wheels
// 	Wheels
// }

// func main() {

// 	myCar := Vehicle{
// 		Brand: "Lada",
// 		Color: "Red",
// 		Wheels: Wheels{
// 			Brand:  "Icon",
// 			Radius: 18,
// 		},
// 	}

// 	fmt.Println(myCar)
// 	myCar.Wheels.Run()
// 	myCar.Run()
// }

// #############################################
/*
Реализовать 2 структуры Кошка, Собака
	У кошки и собаки: Имя, голос
	У каждого есть метод Speak
	Вывод: Кошка Муся(%s) говорит Мяу(%s)
*/

package main

import "fmt"

// Структура
type Animals struct {
	Type  string
	Name  string
	Voice string
}

// Метод
func (v Animals) PrintVoice() {
	fmt.Printf("%s %s говорит %s\n", v.Type, v.Name, v.Voice)
}

func main() {
	cat := Animals{Type: "Кошка", Name: "Муся", Voice: "Мяу"}
	dog := Animals{Type: "Собака", Name: "Шарик", Voice: "Гав"}
	// Вызов метода
	cat.PrintVoice()
	dog.PrintVoice()
}
