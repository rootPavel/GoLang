package main

import "fmt"

func main() {
	mySlice := []int{1, 2, -5, -6, 7, -4}
	fmt.Println(ParseInt(mySlice))
}

func ParseInt(mySlice []int) []int {
	var positiveSlice []int
	for _, v := range mySlice {
		if v >= 0 {
			positiveSlice = append(positiveSlice, v)
		}
	}
	return positiveSlice
}
