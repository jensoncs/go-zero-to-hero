package main

import "fmt"

func main() {
	a := [3][2]string{
		{"lion", "tiger"},
		{"cat", "dog"},
		{"duck", "pigion"},
	}

	var b [3][2]string

	b[0][0] = "apple"
	b[0][1] = "samsung"
	b[1][0] = "lg"
	b[1][1] = "onida"
	b[2][0] = "haier"
	b[2][1] = "micromax"

	printarry(a)
	printarry(b)
}

func printarry(a [3][2]string) {
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
		fmt.Println("/n")
	}
}
