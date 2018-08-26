package main

import (
	"fmt"
	"time"
)

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Println(i)
	}
}

func alphebets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c\n", i)
	}
}

func main() {
	go numbers()
	go alphebets()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("Main Over")
}
