package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(info.IsDir())
	fmt.Println(info.Name())
	fmt.Println(info.Size())
	fmt.Println(info.Mode())
}
