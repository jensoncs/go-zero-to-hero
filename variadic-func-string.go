package main

import "fmt"

func main() {
	welcome := []string{"Hi", "World"}

	change(welcome...)
	fmt.Println(welcome)
}

func change(s ...string) {
	s[0] = "Go"
}
