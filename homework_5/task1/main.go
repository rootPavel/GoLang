package main

import "fmt"

type Book struct {
	Tile   string
	Author string
	Year   int
}

func (p Book) GetInfo() {
	fmt.Printf("\"%s\" by %s, %d", p.Tile, p.Author, p.Year)
}

func main() {
	goBook := Book{
		Tile:   "Golang для профи",
		Author: "Михалис Цукалос",
		Year:   2025,
	}

	goBook.GetInfo()
}
