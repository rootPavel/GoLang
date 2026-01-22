package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

func main() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println(err)
		return
	}
	publicKey := &privateKey.PublicKey
	data := []byte("Hello from RSA")

	cipherText, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		publicKey,
		data,
		nil,
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Data: ", string(data))
	fmt.Printf("Cipher: %x\n", cipherText)

	ClearData, err := rsa.DecryptOAEP(
		sha256.New(),
		rand.Reader,
		privateKey,
		cipherText,
		nil,
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Decrypted Data: %s\n", ClearData)
}
