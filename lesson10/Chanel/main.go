/*
### Канал ###
c := make(chan int)
c <- 1
v := <- c
*/

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

	c := make(chan int)
	for i := 0; i < 5; i++ {
		go PrintInGoroutine(i, c)
	}

	for i := 0; i < 5; i++ {
		goId := <-c // не кринж
		fmt.Printf("Routine #%d finished\n", goId)
	}
}
