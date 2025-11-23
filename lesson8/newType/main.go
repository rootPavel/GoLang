package main

import "fmt"

type cat string
type dog string

func (c cat) Speak() {
	fmt.Printf("Кошка %s говорит 'Мяу'\n", c)
}

func (d dog) Speak() {
	fmt.Printf("Собака %s говорит 'Гав'\n", d)
}

func main() {
	var Musya cat = "Муся"
	var Sharik dog = "Шарик"

	Musya.Speak()
	Sharik.Speak()

}
