package myfunc

import "strconv"

func SearchCity(listCities []string, searchCity string) string {
	result := "Ваш город не найден в списке!"
	for i, city := range listCities {
		if city == searchCity {
			result = "Ваш город найден в списке под порядковым номером " + strconv.Itoa(i+1)
		}
	}
	return result
}
