package main

import "fmt"

func main() {
	finger := 6
	switch finger {
	case 1, 2, 3, 4, 5:
		fmt.Println("hello")
	default:
		fmt.Println("Hello11")
	}
}
