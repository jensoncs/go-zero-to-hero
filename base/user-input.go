package main

import "fmt"

func main() {

	var sn int
	var ln int

	fmt.Print("Please enter a small number : ")
	fmt.Scan(&sn)
	fmt.Print("Please enter a large number : ")
	fmt.Scan(&ln)

	fmt.Println("The reminder is : ", ln/sn)
}
