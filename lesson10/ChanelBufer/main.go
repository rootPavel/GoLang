package main

import (
	"fmt"
	"time"
)

func PrintInGoroutine(i int, c chan int) {
	time.Sleep(1 * time.Second)
	c <- i
}

func main() {
	c := make(chan int, 2) // буферезированный канал
	c <- 2
	fmt.Println(<-c)
	c <- 3
	c <- 4
	fmt.Println(<-c)
	fmt.Println(<-c)
}
