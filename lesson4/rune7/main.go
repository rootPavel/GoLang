package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// AsciiQuestion := "Hello?"
	// fmt.Println(len(AsciiQuestion))
	question := "¿Cómo estás?"
	// fmt.Println(len(question))

	// fmt.Println(utf8.RuneCountInString(AsciiQuestion))
	fmt.Println(utf8.RuneCountInString(question))

	letter, size := utf8.DecodeRuneInString(question)
	fmt.Printf("if letter %c %v bytes", letter, size)
}
