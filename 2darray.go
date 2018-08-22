package main

import "fmt"

func main() {
	a := [2][2]int{}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			a[i][j] = i + j
		}
	}
	fmt.Println(a)
}
