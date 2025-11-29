/*
Создайте программу, которая читает лог-файл server.log, подсчитывает количество строк, содержащих слово "error", и выводит это число.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func OpenLogFile(path string) *os.File {
	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	return file
}

func main() {
	pathToFile := "./logfile.log"
	countError := 0
	findString := "error"
	var wg sync.WaitGroup
	var mu sync.Mutex

	file := OpenLogFile(pathToFile)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		wg.Add(1)
		go func(line string) {
			defer wg.Done()
			if strings.Contains(strings.ToLower(line), findString) {
				mu.Lock()
				countError++
				mu.Unlock()
			}
		}(line)

	}
	wg.Wait()
	fmt.Printf("The %s file were found when scanning %d errors", pathToFile, countError)

}
