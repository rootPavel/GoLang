package main

import (
	"bufio"
	"fmt"
	"os"
	myfunc "task4/MyFunc"
)

func main() {

	fmt.Println("Введите формулу для проверки скобок:")
	enterFormula := bufio.NewReader(os.Stdin)
	formula, err := enterFormula.ReadString('\n')

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(myfunc.CheckParenthesis(formula))

}
