package main

import "fmt"
import "time"

func main() {
	ch := make(chan int, 2)
	go write(ch)

	time.Sleep(2 * time.Second)
	for v := range ch {
		fmt.Printf("Read value v from ch %v\n", v)
		time.Sleep(2 * time.Second)
	}
}

func write(ch chan int) {
	for i := 0; i <= 5; i++ {
		ch <- i
		fmt.Printf("Succeff fully wrote i %v to channel\n", i)
	}
	close(ch)
}
