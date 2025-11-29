package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.OpenFile("hello.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	hello := "Hello!\nNew line\n"
	file.WriteString(hello)
}
