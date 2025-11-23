package main

import "fmt"

var name string

func init() {
	name = "Sergey"
}

func init() {
	name = "Alex"
}

func main() {
	fmt.Println(name)
}
