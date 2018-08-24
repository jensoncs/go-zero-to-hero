package main

import (
	"fmt"
)

func reverse(x, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Println(reverse("jenson", "hello"))
}
