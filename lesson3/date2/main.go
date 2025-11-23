package main

import (
	"fmt"
	"time"
)

func main() {
	myTime := time.Now()
	fmt.Println(myTime)
	fmt.Println(myTime.Format("02.01.2006 15:04"))
	fmt.Println(myTime.Format("2006 Jan 02 03:04:05"))

}
