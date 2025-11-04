package main

import "fmt"

func main() {
	var a, b, c, d, e int
	var mean float64
	fmt.Print("Введите первое число: ")
	fmt.Scan(&a)
	fmt.Print("Введите второе число: ")
	fmt.Scan(&b)
	fmt.Print("Введите третье число: ")
	fmt.Scan(&c)
	fmt.Print("Введите четвертое число: ")
	fmt.Scan(&d)
	fmt.Print("Введите пятое число: ")
	fmt.Scan(&e)

	fmt.Printf("Вы ввели %d, %d, %d, %d, %d\n", a, b, c, d, e)

	if a < b {
		a, b = b, a
	}
	if a < c {
		a, c = c, a
	}
	if a < d {
		a, d = d, a
	}
	if a < e {
		a, e = e, a
	}

	if b < c {
		b, c = c, b
	}
	if b < d {
		b, d = d, b
	}
	if b < e {
		b, e = e, b
	}

	if c < d {
		c, d = d, c
	}
	if c < e {
		c, e = e, c
	}

	if d < e {
		d, e = e, d
	}

	mean = float64(a+b+c+d+e) / 5
	fmt.Printf("Ваши числа отсортированы %d, %d, %d, %d, %d\n", a, b, c, d, e)
	fmt.Printf("Самое большое число %d\n", a)
	fmt.Printf("Самое маленькое число %d\n", e)
	fmt.Printf("Среднее арифметическое %.2f\n", mean)
}
