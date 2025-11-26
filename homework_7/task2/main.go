package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func GenerateSend(push chan int) {
	numbers := []int{}
	for i := 0; i < 10; i++ {
		numbers = append(numbers, rand.Intn(9)+1)
	}

	for _, v := range numbers {
		push <- v
	}
	close(push)
}

func ReceivePrint(push chan int, wg *sync.WaitGroup) {
	for v := range push {
		fmt.Println(v)
	}
	wg.Done()

}

func main() {
	push := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go GenerateSend(push)
	go ReceivePrint(push, &wg)
	wg.Wait()
}
