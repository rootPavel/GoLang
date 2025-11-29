/*
Напишите программу, которая принимает список имен файлов в текущей директории,
объединяет их содержимое и сохраняет результат в новый файл combined.txt (работать только с текстовыми файлами)
*/

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

func CreateOrClearFile(name string) {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println(err)
	}
	file.Close()
}

func ConvertToUTF8(s string) string {
	reader := transform.NewReader(bytes.NewReader([]byte(s)), charmap.Windows1251.NewDecoder())
	utf8Bytes, err := io.ReadAll(reader)
	if err != nil {
		fmt.Println(err)
	}
	return string(utf8Bytes)
}

func CombinedFiles(basefilePath, addfilePath string, mu *sync.Mutex) {
	basefile, err := os.OpenFile(basefilePath, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer basefile.Close()

	addfile, err := os.OpenFile(addfilePath, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer addfile.Close()

	scanner := bufio.NewScanner(addfile)
	for scanner.Scan() {
		mu.Lock()
		strUTF8 := ConvertToUTF8(scanner.Text())
		if strUTF8 != "" {
			basefile.WriteString(strUTF8 + "\n")
		}
		mu.Unlock()
	}

}

func main() {
	startScanPath := "./"
	combinedFileName := "combined.txt"
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(1)

	CreateOrClearFile(combinedFileName)

	go func() {
		defer wg.Done()
		filepath.WalkDir(startScanPath, func(path string, d fs.DirEntry, err error) error {
			if d.IsDir() && path != "./" {
				return fs.SkipDir
			}
			if !d.IsDir() {
				fileName := strings.ToLower(d.Name())
				if filepath.Ext(fileName) == ".txt" && d.Name() != combinedFileName {
					CombinedFiles(combinedFileName, path, &mu)
				}
			}
			return nil
		})
	}()

	wg.Wait()
}
