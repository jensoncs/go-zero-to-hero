package main

import (
	"fmt"
	"github.com/gobuffalo/packr"
)

func main() {
	box := packr.NewBox("./")
	data := box.String("test.txt")
	fmt.Println("contents of file", data)
}
