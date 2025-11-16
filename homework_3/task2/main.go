/*
Создайте слайс строк, содержащий названия городов.
Реализуйте функции для добавления нового города, удаления города по имени и поиска города в списке.
Продемонстрируйте работу этих функций на примере.
*/

package main

import (
	"fmt"
	myfunc "task2/MyFunc"
)

func main() {
	sliceCities := []string{
		"Москва",
		"Санкт-Петербург",
		"Казань",
		"Краснодар",
		"Екатеринбург",
	}

	fmt.Printf("Slice городов:\n%s\n\n", sliceCities)

	sliceCities = myfunc.AddCity(sliceCities, "Архангельск")
	fmt.Printf("Реализация функции добавления города Архангельск:\n%s\n\n", sliceCities)

	sliceCities = myfunc.DeleteCity(sliceCities, "Краснодар")
	fmt.Printf("Реализация функции удаления города Краснодар:\n%s\n\n", sliceCities)

	fmt.Printf("Реализация функции поиска города Казань:\n%s\n\n", myfunc.SearchCity(sliceCities, "Казань"))

	fmt.Printf("Реализация функции поиска города Краснодар:\n%s\n\n", myfunc.SearchCity(sliceCities, "Краснодар"))
}
