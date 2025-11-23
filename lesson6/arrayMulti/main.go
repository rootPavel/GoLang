package main

import "fmt"

func main() {
	// многомерный массив
	var board [3][4]int
	board[0][0] = 1
	board[0][1] = 2
	board2 := [2][2]int{{1, 2}, {4, 5}}
	fmt.Println(board)
	fmt.Println(board2)
}
