package main

import "fmt"

func foo() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("Я поймал панику!")
		}
	}()
	panic("Это паника функии foo()")
}

func main() {
	foo()
}
