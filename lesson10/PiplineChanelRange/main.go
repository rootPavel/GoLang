package main

import (
	"fmt"
	"strings"
)

func GenerateData(downstream_c0 chan string) { //1
	data := []string{"Hi", "D", "bad", "", "i am very bad", "", "go"}
	for _, v := range data {
		downstream_c0 <- v
	}
	close(downstream_c0)
}
func FilterData(upperstream_c0, downstream_c1 chan string) { //2
	for item := range upperstream_c0 {
		if !strings.Contains(item, "bad") {
			downstream_c1 <- item
		}
	}
	close(downstream_c1)
}
func PrintData(upperstream_c1 chan string) { //3
	for item := range upperstream_c1 {
		fmt.Println(item)
	}
}
func main() {
	c0 := make(chan string)
	c1 := make(chan string)
	go GenerateData(c0)
	go FilterData(c0, c1)
	PrintData(c1)
}
