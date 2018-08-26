package main

import "fmt"

func main() {

	sendch := make(chan int)
	go sendData(sendch)
	fmt.Println(<-sendch)
}

func sendData(sendch chan<- int) {
	sendch <- 10
}
