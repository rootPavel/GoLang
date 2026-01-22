package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	password := []byte("12345678")
	hash := sha256.Sum256(password)
	fmt.Printf("hash: %x\n", hash)
}
