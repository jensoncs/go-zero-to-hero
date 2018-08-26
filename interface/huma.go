package main

import "fmt"

type human interface {
	sayhellow() string
}

type men struct {
	greeting string
}

type women struct {
	greeting string
}

func (m men) sayhellow() string {
	return m.greeting
}

func (w women) sayhellow() string {
	return w.greeting
}

func main() {
	thomas := men{greeting: "Hey..."}

	polly := women{greeting: "Hi..."}

	printGreeting(thomas)
	printGreeting(polly)
}

func printGreeting(h human) {
	fmt.Printf("%T,%v", h, h)
	fmt.Println(h.sayhellow())
}
