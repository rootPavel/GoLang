package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

func str() {
	fmt.Println(strconv.IntSize)
}

func main() {
	myTime := time.Now()
	fmt.Println(myTime)
	fmt.Println(myTime.Format("02 Jan 2006, 15:04"))
	fmt.Println(myTime.Format("02.01.06 3:04"))
	fmt.Println(myTime.Format("2006-01-02 15-04-05"))
	fmt.Println(myTime.Format("Jan, 02."))

	str()
	fmt.Println((math.MaxFloat64))
	fmt.Println((math.MaxInt8))

}

// 06 Nov 2025, 21:22
// 06.11.25 9:22
// 2025-11-06 21-23-50
// Nov, 06.
