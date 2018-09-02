package main

import "fmt"

func main() {

	defer fmt.Println("deffered calling func main")
	firstName := "Sam"
	//	lastName := nil
	fullName(&firstName, nil)
	fmt.Println("Returned normally form main")
}

func fullName(firstName *string, lastName *string) {
	defer fmt.Println("deffered calling fullname")
	if firstName == nil {
		panic("first name cant be nil")
	}
	if lastName == nil {
		panic("last name cant be nil")
	}

	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returend nortmally from full name")
}
