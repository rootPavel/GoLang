package main

import "fmt"

func main() {
	animals := [...]string{
		"Федя",
		"Тема",
		"Йо",
		"Джордж",
		"Пушок",
		"Сема",
	}
	myAnimals := animals[0:3] //[:3] [3:]
	fmt.Println(myAnimals)
	myAnimals[0] = "Жора"
	fmt.Println(animals)

	myString := "Hello, dear"
	fmt.Println(myString[:5])

	myAnimals = animals[:] // срез всех элементов исходного массива

	mySlice := []string{
		"Сергей",
		"Алиса",
		"Матвей",
		"Арсений",
	}

	fmt.Println(len(mySlice), cap(mySlice))

	mySlice = append(mySlice, "Елена")
	fmt.Println(len(mySlice), cap(mySlice))

	mySlice = append(mySlice, "Алена")
	fmt.Println(len(mySlice), cap(mySlice))

	mySlice = append(mySlice, "Павел", "Трифон", "ктото")
	fmt.Println(mySlice)
	fmt.Println(len(mySlice), cap(mySlice))

	mySlice = append(mySlice[:2], mySlice[2+1:]...)
	fmt.Println(mySlice)
	fmt.Println(len(mySlice), cap(mySlice))
}
