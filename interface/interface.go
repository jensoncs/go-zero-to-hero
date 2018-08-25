package main

import "fmt"

type VowelsFinder interface {
	FindVowels() []rune
}

type MyStirng string

func (ms MyStirng) FindVowels() []rune {
	var vowels []rune
	for _, runes := range ms {
		if (runes == 'a') || (runes == 'e') || (runes == 'i') || (runes == 'o') || (runes == 'u') {
			vowels = append(vowels, runes)
		}
	}
	return vowels
}

func main() {
	name := MyStirng("Jenson")
	var v VowelsFinder
	v = name
	fmt.Println(string(v.FindVowels()))
}
