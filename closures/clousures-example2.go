package main

import "fmt"

func main() {
	a := appendstr()
	b := appendstr()

	fmt.Println(a("world"), b("Everyone"))
	fmt.Println(a("Golang"), b("!"))
}

func appendstr() func(string) string {
	t := "Hello"
	c := func(b string) string {
		t = t + "" + b
		return t
	}
	return c
}
