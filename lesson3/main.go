// package main

// import "fmt"

// func main() {
// 	const pi float64 = 3.14
// 	const (
// 		long  int = 12
// 		width int = 12
// 	)
// 	const long2, width2 = 13, 14
// 	const (
// 		a = 1
// 		b //b = ?
// 		c //c = ?
// 		d = 2
// 		e //e = ?
// 	)

// 	fmt.Println(a, b, c, d, e)
// }

// ###################################

// package main

// import "fmt"

// func main() {
// 	a := 1.0 / 3
// 	fmt.Println(a)
// 	fmt.Printf("%v\n", a)
// 	fmt.Printf("%f\n", a)
// 	fmt.Printf("%.2f\n", a) //.2 - после запятой 2 числа
// 	fmt.Printf("%08.2f\n", a)
// }

//#####################################

// package main

// import "fmt"

// func main() {
// 	a := 1.0 / 3
// 	fmt.Println(a + a + a)

// 	b := 0.1
// 	b += 0.2

// 	fmt.Println(b)
// 	fmt.Printf("%.60f\n", 0.1)
// 	fmt.Printf("%.60f\n", 0.2)

// 	fmt.Println(b == 0.3)
// 	fmt.Println((b - 0.3) < 0.00001)
// }

// ##############################

// package main

// import "fmt"

// func main() {
// 	var course float32 = 84.8
// 	var count float32

// 	fmt.Print("Введите сумму которую необходимо перевести в тугрики: ")
// 	fmt.Scan(&count)

// 	fmt.Printf("Необходимо выдать: %.2f\n", transfer(count, course))

// }

// func transfer(money, course_2 float32) float32 {
// 	return money / course_2
// }

// #############################

// package main

// import "fmt"

// func main() {
// 	var a int8 = 127
// 	a += 1

// 	fmt.Printf("%v\n", a)

// 	var r16 uint8 = 0x00

// 	fmt.Printf("%x\n", r16)

// }

// // ################################

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	/*
// 		01 - месяц (январь)
// 		02 - число (2 января)
// 		03 - час (2 января 3 часа)
// 		04 - минута (2 января 3 часа 4 минуты)
// 		05 - секунда (2 января 3 часа 4 минуты 5 секунд)
// 		06 - год (2 января 3 часа 4 минуты 2006 года)
// 		any languages: fmt.Printf("%h:%m:%s")
// 	*/
// 	template := "2006-01-02"
// 	myDate := "2025-11-06"

// 	myTime, err := time.Parse(template, myDate)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(myTime)
// 	fmt.Printf("%T", myTime)

// }

// // ################################

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	myTime := time.Now()
// 	fmt.Println(myTime)
// 	fmt.Println(myTime.Format("02.01.2006 15:04"))
// 	fmt.Println(myTime.Format("2006 Jan 02 03:04:05"))

// }

// ################################

package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

func main() {
	myTime := time.Now()
	fmt.Println(myTime)
	fmt.Println(myTime.Format("02 Jan 2006, 15:04"))
	fmt.Println(myTime.Format("02.01.06 3:04"))
	fmt.Println(myTime.Format("2006-02-01 15-04-05"))
	fmt.Println(myTime.Format("Jan, 02."))

	fmt.Println(str)
	fmt.Println((math.MaxFloat64))
	fmt.Println((math.MaxInt8))

}

func str() {
	fmt.Println(strconv.IntSize)
}

// 06 Nov 2025, 21:22
// 06.11.25 9:22
// 2025-06-11 21-23-50
// Nov, 06.
