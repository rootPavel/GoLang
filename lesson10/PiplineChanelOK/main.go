package main

import (
	"fmt"
	"strings"
)

func GenerateData(downstream_c0 chan string) {
	data := []string{"Hi", "D", "bad", "", "a am very bad", "go"}

	for _, v := range data {
		downstream_c0 <- v
	}
	close(downstream_c0) // хороший вариант закрытия канала
}

func FilterData(upperstream_c0, downstream_c1 chan string) {
	for {
		item, ok := <-upperstream_c0
		if !ok {
			close(downstream_c1)
			return
		}
		if !strings.Contains(item, "bad") {
			downstream_c1 <- item
		}
	}
}

func PrintData(upperstream_c1 chan string) {
	for {
		item, ok := <-upperstream_c1
		if !ok {
			return
		}
		fmt.Println(item)
	}
}

func main() {
	c0 := make(chan string)
	c1 := make(chan string)

	go GenerateData(c0)
	go FilterData(c0, c1)

	PrintData(c1) // не го рутина что бы функция Print успела отработать
}
