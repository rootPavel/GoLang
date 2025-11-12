package main

import (
	"bufio"
	"fmt"
	"os"
	myfunc "task4/MyFunc"
)

// Напишите программу, которая запрашивает у пользователя ввод строки-формулы, а выводит сообщение о правильности написания круглых скобок, например:
// Пример 1
// строка на вход: (1+1)*(2+2)
// вывод: Скобки расставлены верно, 2 открывающиеся, 2 закрывающиеся

// Пример 2
// Строка на вход: ((1+1) + (2+2) ))
// вывод: Скобки расставлены неправильно, 3 открывающиеся, 4 закрывающиеся

func main() {

	fmt.Println("Введите формулу для проверки скобок:")
	enterFormula := bufio.NewReader(os.Stdin)
	formula, err := enterFormula.ReadString('\n')

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(myfunc.CheckParenthesis(formula))

}
