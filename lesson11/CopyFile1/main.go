package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	file, err := os.Open("hello.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	io.Copy(os.Stdout, file)
}
