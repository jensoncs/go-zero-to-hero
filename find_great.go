package main

import "fmt"

func main() {
	fmt.Println(findGreatest(1, 2, 3, 4, 5, 67, 8))
}

func findGreatest(arg ...int) int {
	value := 0
	for _, v := range arg {
		if v > value {
			value = v
		}
	}
	return value
}
