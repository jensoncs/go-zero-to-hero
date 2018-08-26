package main

import (
	"fmt"
	"sync"
	"time"
)

func process(i int, wg *sync.WaitGroup) {

	fmt.Println("started", i)
	time.Sleep(2 * time.Second)
	fmt.Println("ended", i)
	wg.Done()
}
func main() {
	no := 3
	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		wg.Add(1)
		go process(i, &wg)
	}
	wg.Wait()
	fmt.Println("All operation finished")
}
