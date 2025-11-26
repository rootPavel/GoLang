package main

import (
	"fmt"
	"sync"
)

func PrintMe(ind int, wg *sync.WaitGroup) {
	fmt.Printf("%d    =====\n", ind)
	defer wg.Done()
}
func main() {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		fmt.Printf("=====    %d\n", i)
		wg.Add(1)
		go PrintMe(i, &wg)
	}
	wg.Wait()
}
