package main

import (
	"fmt"
	"strings"
)

func main() {
	string1 := "My name is Dima, Dima, Dima"
	result := strings.Replace(string1, "Dima", "Misha", 2) // -1 все заменить
	fmt.Print(result)

}
