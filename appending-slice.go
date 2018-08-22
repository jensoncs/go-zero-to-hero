package main

import "fmt"

func main() {

	veg := []string{"tomoato", "potato", "onion"}
	fruit := []string{"apple", "orange"}
	food := append(veg, fruit...)
	fmt.Println(food)
}
