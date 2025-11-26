package main

import (
	"fmt"
	"sync"
	"time"
)

func MyPrint(v int, wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	fmt.Println(v)
	wg.Done()
}

func main() {
	var mySlice = [5]int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup

	for _, v := range mySlice {
		wg.Add(1)
		go MyPrint(v, &wg)
		wg.Wait()
	}
}
