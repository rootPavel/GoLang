/*
Напишите программу, которая создает массив из 10 целых чисел, заполняет его случайными значениями от 1 до 100.
Затем скопируйте этот массив в слайс и отсортируйте его по возрастанию.
Выведите исходный массив и отсортированный слайс.
*/

package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	var arrayRandomNumber [10]int
	sliceCopyRandom := make([]int, len(arrayRandomNumber))

	for i := range arrayRandomNumber {
		arrayRandomNumber[i] = rand.Intn(100) + 1
	}

	copy(sliceCopyRandom, arrayRandomNumber[:])
	sort.Ints(sliceCopyRandom)

	fmt.Println(arrayRandomNumber)
	fmt.Println(sliceCopyRandom)
}
