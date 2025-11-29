package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	hello := "Hello!\n"
	file.WriteString(hello)

	// Выше создание и запись

	data, err := os.ReadFile("test.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(data))

}
