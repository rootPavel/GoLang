package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	fmt.Println("Для подсчета количества символов введите строку:")
	stringValues := bufio.NewReader(os.Stdin)
	rawString, err := stringValues.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	countString := utf8.RuneCountInString(strings.TrimSpace(rawString))
	fmt.Printf("В вашей строке %d символов!", countString)
}
