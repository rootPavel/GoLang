package main

import (
	"bufio"
	"fmt"
	"os"
	myfunc "task2/MyFunc"
)

func main() {
	fmt.Println("Для подсчета количества гласных букв введите строку:")
	stringValues := bufio.NewReader(os.Stdin)
	rawString, err := stringValues.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	countVowel := myfunc.CountVowel(rawString)

	fmt.Printf("Количество гласных букв в строке %d", countVowel)
}
