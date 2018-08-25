package main

import (
	"fmt"
	"reflect"
)

func main() {
	personSalary := make(map[string]int)

	personSalary["stev"] = 10000
	personSalary["joe"] = 5000
	fmt.Println(personSalary)

	personSalary1 := map[string]int{
		"stev": 10000,
		"joe":  50000,
	}
	fmt.Println(personSalary1)
	fmt.Println(personSalary1["stev"])

	value, ok := personSalary["joe"]
	if ok == true {
		personSalary["joe"] = 10000
	}
	fmt.Println("Updated salary for Joe is:", value)

	for key, value := range personSalary {
		fmt.Printf("personSalary[%s] = %d \n", key, value)
	}

	newPersonSalary := personSalary
	fmt.Println("******************")
	fmt.Println(reflect.DeepEqual(personSalary, newPersonSalary))
	fmt.Println("******************")
	fmt.Println(personSalary)
	newPersonSalary["stev"] = 20000
	fmt.Println(personSalary)
	fmt.Println(newPersonSalary)
	fmt.Println(len(personSalary))
	delete(personSalary, "stev")
	fmt.Println(personSalary)
	fmt.Println(len(personSalary))

}
