package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("Введите пароль для хеширования: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Println(err)
		return
	}
	pass := []byte(strings.TrimSpace(input))

	hash := sha256.Sum256(pass)
	fmt.Printf("hash: %x\n", hash)
}
