package main

import "fmt"

type rubles int

func (r *rubles) Incr() {
	*r++
}

func main() {
	var r rubles = 100

	r.Incr()
	r.Incr()
	r.Incr()
	r.Incr()

	fmt.Println(r)
}
