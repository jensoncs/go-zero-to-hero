package main

import "fmt"

func main() {
	hexBytes := []byte{
		0x43, 0x61, 67, 97,
	}

	runes := []rune{
		0x43, 0x61, 67, 97,
	}
	fmt.Println(string(hexBytes))
	fmt.Println(string(runes))
}
