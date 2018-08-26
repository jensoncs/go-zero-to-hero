package main

import "fmt"

func producer(cha chan int) {
	for i := 0; i < 10; i++ {
		cha <- i
	}
	close(cha)
}

func main() {
	ch := make(chan int)
	go producer(ch)
	for {
		v, ok := <-ch //check the channel status,
		if ok == false {
			break
		}
		fmt.Printf("is ok %v\n", v)
	}
}
