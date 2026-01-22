package main

import "fmt"

func main() {
	a := []int{10, 20, 30, 40}
	b := a[:2]         //10, 20
	c := append(b, 99) //a? a = [10,20,99,40]
	a[1] = 777         //b? c?
	c = append(c, 555) //a?
	fmt.Println(a)     //10 777 99 555
	fmt.Println(b)
	fmt.Println(c)
}
