package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Name() string
}

type Circle struct {
	Type   string
	Radius float64
}

func (c Circle) Area() float64 {
	s := math.Pi * c.Radius * c.Radius
	return s
}

func (n Circle) Name() string {
	return n.Type
}

type Rectangle struct {
	Type    string
	LengthA float64
	LengthB float64
}

func (r Rectangle) Area() float64 {
	s := r.LengthA * r.LengthB
	return s
}

func (n Rectangle) Name() string {
	return n.Type
}

func PrintArea(figure []Shape) {
	for _, v := range figure {
		fmt.Printf("Площадь фигуры %s равна %.2f\n", v.Name(), v.Area())
	}
}

func main() {
	figure := []Shape{
		Circle{Type: "Круг", Radius: 12.34},
		Rectangle{Type: "Прямоугольник", LengthA: 23.56, LengthB: 11},
	}

	PrintArea(figure)
}
