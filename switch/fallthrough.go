package main

import "fmt"

func main() {

	finger := 2

	switch finger {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
		fallthrough
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("default")
	}
}
