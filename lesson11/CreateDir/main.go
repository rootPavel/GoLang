package main

import (
	"fmt"
	"os"
)

func main() {
	// err := os.Mkdir("test", 0644)
	err := os.MkdirAll("./test/folder1/folder2", 0644)
	if err != nil {
		fmt.Println(err)
	}

}
