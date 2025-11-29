package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("test.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fmt.Fprint(file, "Hello, World\n")
}
