package checkpass

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func CheckPass(pass, hash string) bool {

	passHash := sha256.Sum256([]byte(pass))
	passStr := hex.EncodeToString(passHash[:])

	if hash == passStr {
		fmt.Println("Пароль верный!")
		return true
	} else {
		return false
	}
}
