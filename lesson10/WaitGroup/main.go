package main

import (
	"fmt"
	"sync"
	"time"
)

func PrintMe(ind int, wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	fmt.Printf("Route #%d finished\n", ind)
	defer wg.Done()
}

func main() {
	var wg sync.WaitGroup
	// wg.Add(3)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go PrintMe(i, &wg)
	}
	wg.Wait()
}
