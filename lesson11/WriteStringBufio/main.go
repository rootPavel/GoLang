package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	write := bufio.NewWriter(file)

	write.WriteString("Нулевая строка \n")
	write.WriteString("Первая строка \n")
	write.WriteString("Следующая строка \n")
	write.Flush()
}
