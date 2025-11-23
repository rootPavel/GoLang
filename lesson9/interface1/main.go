package main

import "fmt"

type Mover interface {
	move()
}

type Bus struct {
	Number int
}

func (b Bus) move() {
	fmt.Printf("Bus #%d moving\n", b.Number)
}

type Train struct {
	Number int
}

func (t Train) move() {
	fmt.Printf("Train #%d moving\n", t.Number)
}

func TransportStart(b Mover) {
	fmt.Printf("Bye Bye to Transport #\n")
}

// func TransportStart(b Bus) {
// 	fmt.Printf("Bye Bye to Transport #%d", b.Number)
// }
// func TransportStart1(b Train) {
// 	fmt.Printf("Bye Bye to Transport #%d", b.Number)
// }
func main() {
	myBus := Bus{1}
	myTrain := Train{2}
	myBus.move()
	myTrain.move()
}
