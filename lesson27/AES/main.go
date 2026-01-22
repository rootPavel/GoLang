package main

import (
	"fmt"

	"github.com/gtank/cryptopasta"
	// лучше использовать эту либу вместо стандартной и пользоваться AES-GCM
)

func main() {
	key := cryptopasta.NewEncryptionKey()
	data := []byte("Hello from AES!")

	cipherText, err := cryptopasta.Encrypt(data, key)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("cipherText: %x\n", cipherText)

	clearData, err := cryptopasta.Decrypt(cipherText, key)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Clear Data: %s\n", clearData)
}
