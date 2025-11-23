package main

import "fmt"

func myFunc(val *int) {
	*val++
}

func main() {
	var p *int
	var val int = 6
	p = &val
	fmt.Println(&val)
	fmt.Println(p)
	fmt.Println(&p)
	fmt.Println(*p)
	*p = 9
	fmt.Println(val)

	myFunc(p)
}
