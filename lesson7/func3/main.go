package main

import "fmt"

func Say(animal string) string {
	if animal == "dog" {
		return "wuf"
	} else if animal == "cat" {
		return "Myau"
	} else {
		return "nah"
	}
}

func PrintVoice(animal string, how func(string) string) {
	fmt.Println(how(animal))
}

func main() {
	PrintVoice("dog", Say)
	PrintVoice("cat", Say)
}
