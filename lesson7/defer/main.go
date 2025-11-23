package main

import "fmt"

func main() {
	fmt.Println("1st output")
	defer fmt.Println("2st output")
	defer fmt.Println("3st output")
	defer fmt.Println("4st output")
}
