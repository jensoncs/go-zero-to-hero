package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("ABC.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(f.Name(), "Opened successfully")
}
