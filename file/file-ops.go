package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println("could not read")
		return
	}
	fmt.Println("contents of file", string(data))
}
