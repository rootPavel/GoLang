// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// func main() {
// 	var a int = 42
// 	var b float64 = float64(a)
// 	var c int32 = int32(a)

// 	s1 := strconv.Itoa(42)           // int → string
// 	s2 := strconv.FormatInt(255, 16) // число в шестнадцатеричной системе

// 	i, err := strconv.Atoi("123")              // string → int
// 	f, err := strconv.ParseFloat("123.45", 64) // string → float64

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	fmt.Println(a, reflect.TypeOf(a))
// 	fmt.Println(b, reflect.TypeOf(b))
// 	fmt.Println(c, reflect.TypeOf(c))
// 	fmt.Println(s1, reflect.TypeOf(s1))
// 	fmt.Println(s2, reflect.TypeOf(s2))
// 	fmt.Println(i, reflect.TypeOf(i))
// 	fmt.Println(f, reflect.TypeOf(f))

// 	num := 88 //"88"
// 	str := strconv.Itoa(num)
// 	fmt.Printf("%v %[1]T\n", str)

// 	str2 := "no"             //-> true ??
// 	boolVar := str2 == "yes" //true, if str == "yes"
// 	fmt.Println(boolVar)
// }

// ######################################
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	fmt.Println("Введите год рождения:")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	enterString, err := strconv.Atoi(strings.TrimSpace(input))

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Ваш возраст: %d", vozrast(enterString))
}

func vozrast(year int) int {
	currentTime := time.Now()
	return currentTime.Year() - year
}
