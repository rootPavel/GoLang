package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

func GenerateSignature(data, key []byte) []byte {
	h := hmac.New(sha256.New, key)
	h.Write(data)
	return h.Sum(nil)
}

func VerifyHMAC(key, data, signature []byte) bool {
	h := hmac.New(sha256.New, key)
	h.Write(data)
	mySig := h.Sum(nil)
	return hmac.Equal(mySig, signature)
}

func main() {
	//Bob
	key := []byte("mySecretKey")
	data := []byte("Hello")
	signature := GenerateSignature(data, key)
	fmt.Println("date: ", string(data))
	fmt.Printf("Sig: %x\n", signature)

	//Alice
	isValid := VerifyHMAC(key, data, signature)
	fmt.Println("Is sig valid: ", isValid)

	//Errors
	keyBad := GenerateSignature(data, []byte("MyNewSecretKey"))
	BadSig := VerifyHMAC(key, data, keyBad)
	fmt.Println("BadSig Is sig valid: ", BadSig)

	dataBad := []byte("Bye!")
	BadData := VerifyHMAC(key, dataBad, signature)
	fmt.Println("date: ", string(dataBad))
	fmt.Println("BadData Is sig valid: ", BadData)
}
