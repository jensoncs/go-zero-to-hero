package main

import "fmt"

func main() {

	for no, i := 10, 1; i <= 10 && no <= 19; i, no = i+1, no+1 {
		fmt.Println(no, i, no+1)
	}
}
