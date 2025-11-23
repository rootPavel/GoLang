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
