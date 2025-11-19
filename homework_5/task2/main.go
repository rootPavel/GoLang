package main

import "fmt"

type Student struct {
	Name   string
	Grades []float64
}

func (g Student) AverageGrade() float64 {
	var sum float64
	for _, v := range g.Grades {
		sum += float64(v)
	}
	averageGrade := sum / float64(len(g.Grades))
	return averageGrade
}

func main() {
	Sergey := Student{
		Name:   "Sergey",
		Grades: []float64{5, 4, 3, 5, 2, 4},
	}
	fmt.Printf("Студент %s получил %.f за семестр!\n", Sergey.Name, Sergey.Grades[:])
	fmt.Printf("Его средний балл %.1f из %d оценок!\n", Sergey.AverageGrade(), len(Sergey.Grades))
}
