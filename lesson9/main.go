// //########## указатели

// package main

// import "fmt"

// func myFunc(val *int) {
// 	*val++
// }

// func main() {
// 	var p *int
// 	var val int = 6
// 	p = &val
// 	fmt.Println(&val)
// 	fmt.Println(p)
// 	fmt.Println(&p)
// 	fmt.Println(*p)
// 	*p = 9
// 	fmt.Println(val)

// 	myFunc(p)
// }

// // ############### указатели

// package main

// import "fmt"

// type rubles int

// func (r *rubles) Incr() {
// 	*r++
// }

// func main() {
// 	var r rubles = 100

// 	r.Incr()
// 	r.Incr()
// 	r.Incr()
// 	r.Incr()

// 	fmt.Println(r)
// }

/*
	Задача 2
	Мы пишем TODO List.
	1) Есть глобальная структура TODOList{[]Tasks}
	2) Структура Tasks - {ID, Title, Done(bool)}
	3) Реализовать для TODOList методы:
		3.1) Добавить задачу(ID должен считаться сам) На вход - title задачи
		3.2) Выполнить задачу. (+ проверку на то, что такая задача есть). На вход - ID задачи
		3.3) Отобразить все задачи с их ID
		3.4) Отобразить все невыполненные задачи
*/

// package main

// func main() {

// }

/*
В ООП с структурами и методами
1) Пишем модуль для библиотеки
	1) Есть глобальная структура Library {[]Book}
	2) Есть структура Book {Author, Title, Availble}
	3) В main нужно заранее положить в Library несколько заранее созданных книг
	4) Для Library реализовать методы:
		4.1) Добавить книгу. rcvr - *Library, входные данные - структура типа книга
		4.2) Взять книгу (Borrow): rcvr - *Library, входные данные - Title книги (строка) // реализуем поиск по названию, проверку на Avail.
		4.3) Вернуть книгу rcvr - *Library, входные данные - Title книги (строка), // реализуем поиск по названию, проверку на Avail. Если такой книги не было - написать что ошиблись
*/

// package main

// import "fmt"

// type Library struct {
// 	Book []Book
// }

// type Book struct {
// 	Author    string
// 	Title     string
// 	Available bool
// }

// func (l *Library) AddBook(book Book) {
// 	l.Book = append(l.Book, book)
// 	fmt.Printf("Книга %s добавлена в библиотеку\n", book.Title)
// }

// func (rcvr *Library) TakeBook(Title string) {
// 	for i, v := range rcvr.Book {
// 		if v.Title == Title {
// 			if v.Available != true {
// 				fmt.Printf("Книга %s уже взята!\n", Title)
// 			} else {
// 				fmt.Printf("Вы взяли книгу %s!\n", Title)
// 				rcvr.Book[i].Available = false
// 			}
// 			return
// 		}
// 	}
// 	fmt.Printf("Такой книги %s не найдено!\n", Title)
// }

// func (l *Library) ReturnBook(title string) {
// 	for i := range l.Book {
// 		if l.Book[i].Title == title {
// 			fmt.Printf("Найдена книга %s\n", l.Book[i].Title)
// 			l.Book[i].Available = true
// 			return
// 		}
// 	}
// 	fmt.Println("Ошиблись. Такой книги не найдено.")
// }

// func (p Library) MyPrint() {
// 	fmt.Println("===================")
// 	fmt.Println("Список книг")
// 	fmt.Println("===================")
// 	for _, v := range p.Book {
// 		fmt.Println(v)
// 	}
// 	fmt.Println("")
// }

// func main() {
// 	library := Library{
// 		Book: []Book{
// 			{Author: "Лев Толстой", Title: "Война и мир", Available: true},
// 			{Author: "Фёдор Достоевский", Title: "Преступление и наказание", Available: false},
// 			{Author: "Антон Чехов", Title: "Рассказы", Available: true},
// 		},
// 	}

// 	library.MyPrint()

// 	// Берем книгу Война и Мир
// 	library.TakeBook("Война и мир")
// 	library.MyPrint()

// 	// Добавляем книгу
// 	book := Book{Author: "Сергей", Title: "Bug in code", Available: true}
// 	library.AddBook(book)
// 	library.MyPrint()

// 	// Берем новую книгу Новую книгу
// 	library.TakeBook("Bug in code")
// 	library.MyPrint()

// 	// Возвращаем книгу Война и Мир
// 	library.ReturnBook("Война и мир")
// 	library.MyPrint()

// 	// Берем уже взятую книгу
// 	library.TakeBook("Bug in code")
// }

// ############## Интерфейсы

// package main

// import "fmt"

// type Mover interface {
// 	move()
// }

// type Bus struct {
// 	Number int
// }

// func (b Bus) move() {
// 	fmt.Printf("Bus #%d moving\n", b.Number)
// }

// type Train struct {
// 	Number int
// }

// func (t Train) move() {
// 	fmt.Printf("Train #%d moving\n", t.Number)
// }

// func TransportStart(b Mover) {
// 	fmt.Printf("Bye Bye to Transport #\n")
// }

// // func TransportStart(b Bus) {
// // 	fmt.Printf("Bye Bye to Transport #%d", b.Number)
// // }
// // func TransportStart1(b Train) {
// // 	fmt.Printf("Bye Bye to Transport #%d", b.Number)
// // }
// func main() {
// 	myBus := Bus{1}
// 	myTrain := Train{2}
// 	myBus.move()
// 	myTrain.move()
// }

package main

import "fmt"

type Mover interface{}

// JSONfile := map[string]interface{}
func main() {
	var a Mover = 1
	fmt.Println(a)
	a = "asd"
	fmt.Println(a)
	a = 1.5
	fmt.Println(a)
}
