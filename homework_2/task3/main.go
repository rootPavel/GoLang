/*
Создайте функцию capitalizeWords(s string) string,
которая преобразует каждое слово в строке так,
чтобы первая буква была заглавной, а остальные — строчными.
Например: "привет мир" → "Привет Мир".
*/

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
