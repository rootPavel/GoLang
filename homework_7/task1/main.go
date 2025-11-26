package main

import (
	"fmt"
	"sync"
	"time"
)

func PrintGorutine(i int, wg *sync.WaitGroup) {
	fmt.Printf("Im GoRutine #%d. Start my task!\n", i)
	time.Sleep(1 * time.Second)
	fmt.Printf("Im GoRutine #%d. Stop my task!\n", i)
	wg.Done()

}

func main() {
	var wg sync.WaitGroup
	countGoRutine := 3
	wg.Add(countGoRutine)

	for i := 1; i <= countGoRutine; i++ {
		go PrintGorutine(i, &wg)
	}
	wg.Wait()
}
