// package main

// import "fmt"

// func main() {
// 	var planets [8]string
// 	planets[0] = "Меркурий"
// 	planets[2] = "Земля"

// 	fmt.Println(planets)
// 	fmt.Println(len(planets))
// 	fmt.Println(planets[1] == "")
// 	// i := 8
// 	// fmt.Println(planets[i])

// 	animals := [3]string{"Федя", "Тема", "Йо"}
// 	fmt.Println(animals)

// 	animals2 := [...]string{
// 		"Федя",
// 		"Тема",
// 		"Йо",
// 		"Соб",
// 	}
// 	fmt.Println(animals2, len(animals2))

// 	for i := 0; i < len(animals); i++ {
// 		fmt.Println(i)
// 	}

// 	for ib, v := range animals2 {
// 		fmt.Println(ib, v)
// 	}

// 	animals3 := animals
// 	animals3[2] = "Kevin"
// 	fmt.Println(animals3)
// }

// // ######################################

// package main

// import "fmt"

// func main() {
// 	// многомерный массив
// 	var board [3][4]int
// 	board[0][0] = 1
// 	board[0][1] = 2
// 	board2 := [2][2]int{{1, 2}, {4, 5}}
// 	fmt.Println(board)
// 	fmt.Println(board2)
// }

// // ######################################

// package main

// import "fmt"

// func main() {
// 	animals := [...]string{
// 		"Федя",
// 		"Тема",
// 		"Йо",
// 		"Джордж",
// 		"Соб",
// 		"Сема",
// 	}
// 	myAnimals := animals[0:3] //[:3] [3:]
// 	fmt.Println(myAnimals)
// 	myAnimals[0] = "Жора"
// 	fmt.Println(animals)

// 	myString := "Hello, dear"
// 	fmt.Println(myString[:5])

// 	myAnimals = animals[:] // срез всех элементов исходного массива

// 	mySlice := []string{"Сергей", "Алиса", "Матвей", "Арсений"}
// 	fmt.Println(len(mySlice), cap(mySlice))

// 	mySlice = append(mySlice, "Елена")
// 	fmt.Println(len(mySlice), cap(mySlice))

// 	mySlice = append(mySlice, "Алена")
// 	fmt.Println(len(mySlice), cap(mySlice))

// 	mySlice = append(mySlice, "Павел", "Трифон", "ктото")
// 	fmt.Println(mySlice)
// 	fmt.Println(len(mySlice), cap(mySlice))

// 	mySlice = append(mySlice[:2], mySlice[2+1:]...)
// 	fmt.Println(mySlice)
// 	fmt.Println(len(mySlice), cap(mySlice))
// }

// // ######################################

// package main

// import "fmt"

// func main() {
// 	mySlice := []int{1, 2, -5, -6, 7, -4}
// 	fmt.Println(ParseInt(mySlice))
// }

// func ParseInt(mySlice []int) []int {
// 	var positiveSlice []int
// 	for _, v := range mySlice {
// 		if v >= 0 {
// 			positiveSlice = append(positiveSlice, v)
// 		}
// 	}
// 	return positiveSlice
// }

// // ######################################

// package main

// import "fmt"

// func main() {
// 	// var person map[string]string
// 	// person = make(map[string]string)
// 	person := map[string]string{"Name": "Serega", "Email": "Sererega@mail.ru"}
// 	person["Age"] = "40"
// 	person["Name"] = "Sergey"
// 	delete(person, "Age")
// 	fmt.Println(person)
// 	fmt.Println(person["Email"])
// }

// ######################################

// На вход программе (в функции main) передается список из ip адрсов []string{"192.168.1.1, 192.168.1.2, 192.168.1.1, 192.168.1.3, 192.168.1.4"}
// 		Нужно вывести на экран:
// 			1) Сколько всего было обращений
// 			2) Сколько уникальных обращений
// 			3) Кто самый популярный ip адрес

package main

import (
	"fmt"
)

func main() {
	listIp := []string{
		"192.168.1.1",
		"192.168.1.2",
		"192.168.1.1",
		"192.168.1.3",
		"192.168.1.4",
		"192.168.1.3",
		"192.168.1.4",
		"192.168.1.1",
	}

	uniqueRequest := CountIpRequest(listIp)
	popularIp := PopularIp(uniqueRequest)

	fmt.Printf("Всего было обращений: %d\n", len(listIp))
	fmt.Printf("Уникальных обращений: %d\n", len(uniqueRequest))
	fmt.Printf("Cамый популярный ip адрес: %s, обращений: %d \n", popularIp, uniqueRequest[popularIp])
}

func CountIpRequest(listIp []string) map[string]int {
	uniqueRequest := make(map[string]int)
	for _, ip := range listIp {
		uniqueRequest[ip]++
	}
	return uniqueRequest
}

func PopularIp(mapIp map[string]int) string {
	count := 0
	var popularIp string
	for ip, request := range mapIp {
		if request > count {
			count = request
			popularIp = ip
		}
	}
	return popularIp
}
