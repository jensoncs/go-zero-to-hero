package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

func main() {
	fptr := flag.String("fpath", "test.txt", "file path to read from")
	flag.Parse()
	fmt.Println("value of file path is:", *fptr)

	data, err := ioutil.ReadFile(*fptr)
	if err != nil {
		fmt.Println("could not read")
		return
	}
	fmt.Println("contents of file", string(data))
}
