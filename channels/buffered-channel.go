package main

import "fmt"

func main() {

	ch := make(chan string, 2)

	ch <- "Ram"
	ch <- "Hari"

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
