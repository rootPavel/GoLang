package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello from other os!")
	time.Sleep(5 * time.Second)
}

/*
Компиляция под другие ос
go env - посомтреть переменные окружения
Set GOOS="linux" | go build -o main .\main.go
*/
