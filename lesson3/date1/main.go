package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		01 - месяц (январь)
		02 - число (2 января)
		03 - час (2 января 3 часа)
		04 - минута (2 января 3 часа 4 минуты)
		05 - секунда (2 января 3 часа 4 минуты 5 секунд)
		06 - год (2 января 3 часа 4 минуты 2006 года)
		any languages: fmt.Printf("%h:%m:%s")
	*/
	template := "2006-01-02"
	myDate := "2025-11-06"

	myTime, err := time.Parse(template, myDate)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(myTime)
	fmt.Printf("%T", myTime)

}
