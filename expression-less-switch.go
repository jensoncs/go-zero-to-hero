package main

import "fmt"

func main() {
	num := 2
	switch {
	case num > 1 && num <= 4:
		fmt.Println("Hello")
	}
}
