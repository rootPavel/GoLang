package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("hello.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// hello := "Hello!\n"
	// file.WriteString(hello) применяется только для строк

	hello := []byte("Hello!\nNew line")
	file.Write(hello)

}
