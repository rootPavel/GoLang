package main

import "fmt"

func main() {
	question := "¿Cómo estás?"
	for i, v := range question {
		fmt.Printf("%v - %c\n", i, v)
	}

	for _, v := range question {
		fmt.Printf("%c\n", v)
	}
	// 26 символов - 26 рун и 26 байт
}
