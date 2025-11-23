package main

import (
	"fmt"
)

func crypto(letter rune) string {
	shifered := letter + 3
	if shifered > 'z' {
		shifered = shifered - 26
	}
	return (fmt.Sprintf("%c", shifered))
}
func main() {
	var letter rune = 'h'
	fmt.Println(crypto(letter))
}
