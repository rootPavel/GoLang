package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

func main() {
	key := []byte("mySecretKey")
	data := []byte("Hello")

	h := hmac.New(sha256.New, key)
	h.Write(data)

	signature := h.Sum(nil)

	fmt.Printf("sig: %x\n", signature)
}
