package main

import (
	"fmt"
	"os"
)

//type PathError struct {
//	op   string
//	path string
//	err  error
//}

//
//func (e *PathError) Error() string {
//	return e.op + "" + e.path + "" + e.err
//}

func main() {
	f, err := os.Open("ABC.txt")
	if err, ok := err.(*os.PathError); ok {
		fmt.Println("file of path", err.Path, "could not find")
		return
	}
	fmt.Println("Opened file successfully", f)
}
