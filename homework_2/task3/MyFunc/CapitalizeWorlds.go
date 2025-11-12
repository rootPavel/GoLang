package myfunc

import (
	"strings"
	"unicode"
)

func CapitalizeWorlds(rawString string) string {
	stringLower := strings.ToLower(rawString)
	wordsArray := strings.Fields(stringLower)

	for i, word := range wordsArray {
		runeArray := []rune(word)
		runeArray[0] = unicode.ToUpper(runeArray[0])
		wordsArray[i] = string(runeArray)
	}

	return strings.Join(wordsArray, " ")
}
