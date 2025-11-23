package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Введите строку:")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(input)
}
