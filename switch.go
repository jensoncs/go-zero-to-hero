package main

import "fmt"

func main() {
	finger := 5

	switch finger {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("No")
	}
}
