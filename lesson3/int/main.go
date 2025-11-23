package main

import "fmt"

func main() {
	var a int8 = 127
	a += 1

	fmt.Printf("%v\n", a)

	var r16 uint8 = 0x00

	fmt.Printf("%x\n", r16)

}
