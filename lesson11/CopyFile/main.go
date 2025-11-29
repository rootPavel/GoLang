package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	src, err := os.Open("hello.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer src.Close()

	dst, err := os.Create("hello_copy.txt")
	if err != nil {
		fmt.Println(err)
	}

	io.Copy(dst, src)

}
