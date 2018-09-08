package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//Tentacles struct
type Tentacles struct {
	Name        string `json: "name"`
	Description string `json: "description"`
}

func (t Tentacles) toString() string {
	bytes, err := json.Marshal(t)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	//fmt.Println(bytes)
	return string(bytes)
}

func getTentacles() []Tentacles {
	tentacles := make([]Tentacles, 3)
	row, err := ioutil.ReadFile("./tentacle.json")

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	json.Unmarshal(row, &tentacles)
	return tentacles
}

func main() {
	tentacles := getTentacles()
	fmt.Println(tentacles)
	for _, te := range tentacles {
		fmt.Println(te.toString())
		fmt.Println(te.Name)
	}
}
