package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("HelloWorld")
}

func main() {
	go hello()
	time.Sleep(1 * time.Second)
	fmt.Println("main func!")
}
