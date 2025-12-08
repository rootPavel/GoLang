package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

// fmt.Errorf()
// errors.New()

type TimeError struct {
	Time time.Time
	Err  error
}

func (te *TimeError) Error() string {
	return fmt.Sprintf("[%v]: %v", te.Time.Format("2006/01/02 15:04:05"), te.Err)
}

func NewTimeError(err error) error {
	return &TimeError{
		Time: time.Now(),
		Err:  err,
	}
}

func ReadFile(fileName string) (string, error) {
	data, err := os.ReadFile(fileName) //err = os.FileNotExist = "Error file not Exist"
	if err != nil {
		return "", NewTimeError(err) // TimeError {Time: 11:26, err: os.FileNotExist}
	}
	return string(data), nil
}

func (te *TimeError) Unwrap() error {
	return te.Err
}

func main() {
	_, err := ReadFile("config.txt")
	if err != nil {
		var myErr *TimeError
		if errors.As(err, &myErr) {
			fmt.Println(myErr.Time)
			fmt.Println(myErr.Err)
		}
	}
}
