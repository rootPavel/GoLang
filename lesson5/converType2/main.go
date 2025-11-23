package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	fmt.Println("Введите год рождения:")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	enterString, err := strconv.Atoi(strings.TrimSpace(input))

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Ваш возраст: %d", age(enterString))
}

func age(year int) int {
	currentTime := time.Now()
	return currentTime.Year() - year
}
