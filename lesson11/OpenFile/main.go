package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("hello.txt") //открывает файл только на чтение
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	hello := "Hello!\nNew line\n"
	file.WriteString(hello)
}
