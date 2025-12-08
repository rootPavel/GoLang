package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

// LabelError описывает ошибку с дополнительной меткой.
type LabelError struct {
	Label string // метка должна быть в верхнем регистре
	Err   error
}

func (err *LabelError) Error() string {
	return fmt.Sprintf("[%s] %v", strings.ToUpper(err.Label), err.Err)
}

func NewLabelError(label string, err error) error {
	return &LabelError{
		Label: label,
		Err:   err,
	}
}

// добавьте методы Error() и NewLabelError(label string, err error) error
// ...
func main() {
	_, err := os.ReadFile("mytest.txt") //net.Dial() err-> NewLableErorr("net", err)
	if err != nil {
		err = NewLabelError("file", err)
		fmt.Println(err)
		// return
	}

	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		err = NewLabelError("net", err)
		fmt.Println(err)
		// return
	}
	if conn != nil {
		defer conn.Close()
	}
	// должна выводить
	// [FILE] open mytest.txt: no such file or directory
}
