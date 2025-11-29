/*
Реализовать на Go модуль, осуществляющий обход директории, переданной программе в качестве аргумента.
	Во время обхода директории Все файлы записываются в файл filse.txt -> директории -> folders.txt | имена файлов
	1) Go рутины для самых уверенных
	2) аргумент можно принять os.Args[1]// main.exe c:\users\user\Desktop\
	3) Для обхода воспользуйтесь filepath.WalkDir()
	4) file.Stat() -> IsDir()?
*/

package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sync"
)

func CreateFile(name string) {
	file, err := os.OpenFile(name, os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
}

func AppendToFile(namefileopen string, textaddfile string) {
	file, err := os.OpenFile(namefileopen, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	file.WriteString(textaddfile)
}

func ClearFile(namefileopen string) {
	file, err := os.OpenFile(namefileopen, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
}

func main() {
	nameFileListFile := "files.txt"
	nameFileListDir := "folders.txt"
	scanPathDir := os.Args[1]
	var wg sync.WaitGroup
	var mu sync.Mutex

	CreateFile(nameFileListFile)
	CreateFile(nameFileListDir)
	ClearFile(nameFileListFile)
	ClearFile(nameFileListDir)
	AppendToFile(nameFileListFile, "List of found files\n")
	AppendToFile(nameFileListDir, "List of found directories\n")

	wg.Go(func() {
		filepath.WalkDir(scanPathDir, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				fmt.Println(err)
				return err
			}

			if d.IsDir() {
				lineText := fmt.Sprintf("- %s in %s:\n", d.Name(), filepath.Dir(path))
				mu.Lock()
				AppendToFile(nameFileListDir, lineText)
				mu.Unlock()
			} else {
				lineText := fmt.Sprintf("- %s in %s:\n", d.Name(), filepath.Dir(path))
				mu.Lock()
				AppendToFile(nameFileListFile, lineText)
				mu.Unlock()
			}
			return nil
		})
	})

	wg.Wait()
}
