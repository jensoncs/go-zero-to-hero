package main

import "fmt"

func main() {

	helloworld := "Helloworld"
	printChar(helloworld)
}

func printChar(s string) {
	runes := []rune(s)
	fmt.Println(runes)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c\n", runes[i])
	}
}
