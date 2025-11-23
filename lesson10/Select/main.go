/*
Множественные каналы select
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func PrintInGoroutine(i int, c1 chan int) {
	time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond)
	c1 <- i
}

func main() {
	c := make(chan int)

	for i := 0; i < 5; i++ {
		go PrintInGoroutine(i, c)
	}

	timeout := time.After(2 * time.Second)

	for i := 0; i < 5; i++ {
		select {
		case goId := <-c:
			fmt.Printf("Routine #%d finished\n", goId)
		case <-timeout:
			fmt.Println("timeout")
			return
		}
	}
}
