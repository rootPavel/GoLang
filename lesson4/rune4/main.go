package main

import "fmt"

func main() {
	message := "Hello!!"

	fmt.Println(message[5])
	fmt.Printf("%c\n", message[5])

	for i := 0; i < 7; i++ {
		fmt.Printf("%c\n", message[i])
	}

}
