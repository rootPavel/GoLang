package myfunc

import (
	"strings"
	"unicode"
)

func UnicWord(srtingWords string) map[string]int {
	listWords := ParsString(srtingWords)
	uniqueWords := make(map[string]int)
	for _, word := range listWords {
		uniqueWords[word]++
	}
	return uniqueWords
}

func ParsString(srtingWords string) []string {
	stringsLowWords := strings.ToLower(srtingWords)
	newString := ""
	for _, symbol := range stringsLowWords {
		if unicode.IsLetter(symbol) || unicode.IsDigit(symbol) {
			newString += string(symbol)
		} else {
			newString += string(" ")
		}
	}
	sliceString := strings.Fields(newString)
	return sliceString
}
