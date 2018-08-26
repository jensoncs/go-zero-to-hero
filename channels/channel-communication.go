package main

import "fmt"
import "time"

func hello(done chan bool) {
	fmt.Println("Go routine is going to sleep")
	time.Sleep(4 * time.Second)
	fmt.Println("Hello world goroutines")
	fmt.Println("Route awake and writing to data")
	done <- true
}

func main() {
	done := make(chan bool)
	fmt.Println("Making call to hello go routine")
	go hello(done)
	value := <-done
	fmt.Println(value)
	fmt.Println("main fn stopped")
}
