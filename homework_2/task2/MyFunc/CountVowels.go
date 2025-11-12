package myfunc

import "strings"

func CountVowel(rawString string) int {
	countVowel := 0
	text := strings.ToLower(rawString)
	for _, v := range text {
		if v == 'а' || v == 'е' || v == 'ё' || v == 'и' || v == 'о' || v == 'у' || v == 'ы' || v == 'э' || v == 'ю' || v == 'я' {
			countVowel++
		}
	}
	return countVowel
}
