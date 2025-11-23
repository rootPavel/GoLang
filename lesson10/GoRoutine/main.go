// ### GoRoutine ###
/*
	go <func>
*/

package main

import (
	"time"
)

func PrintInGoroutine(i int) {
	time.Sleep(1 * time.Second)

}

func main() {

	for i := 0; i < 5; i++ {
		go PrintInGoroutine(i)
	}

	time.Sleep(2 * time.Second) //кринж

}
