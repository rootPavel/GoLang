package main

import (
	"bufio"
	"fmt"
	"os"
	myfunc "task3/MyFunc"
)

func main() {

	fmt.Println("Введите строку для преобразования:")
	stringValues := bufio.NewReader(os.Stdin)
	rawString, err := stringValues.ReadString('\n')

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(myfunc.CapitalizeWorlds(rawString))

}
