/*
Задача 1: Создание структуры "Книга"
Описание:
Создайте структуру Book, которая содержит поля: Title (строка), Author (строка), Year (целое число).
Добавьте метод GetInfo(), который возвращает строку с информацией о книге в формате: "Title" by Author, Year.
Что нужно сделать:
Объявить структуру Book.
Реализовать метод GetInfo() для этой структуры.
Создать экземпляр книги и вывести информацию с помощью метода.
*/

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
