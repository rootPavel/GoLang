package main

import "fmt"

func main() {
	var planets [8]string
	planets[0] = "Меркурий"
	planets[2] = "Земля"

	fmt.Println(planets)
	fmt.Println(len(planets))
	fmt.Println(planets[1] == "")
	// i := 8
	// fmt.Println(planets[i])

	animals := [3]string{"Федя", "Тема", "Йо"}
	fmt.Println(animals)

	animals2 := [...]string{
		"Федя",
		"Тема",
		"Йо",
		"Джордж",
	}
	fmt.Println(animals2, len(animals2))

	for i := 0; i < len(animals); i++ {
		fmt.Println(i)
	}

	for ib, v := range animals2 {
		fmt.Println(ib, v)
	}

	animals3 := animals
	animals3[2] = "Kevin"
	fmt.Println(animals3)
}
