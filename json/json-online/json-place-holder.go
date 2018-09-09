package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type posts struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	url := "https://jsonplaceholder.typicode.com/posts1/"
	posts := getPosts(url)
	for _, p := range posts {
		fmt.Println(p.toString())
	}
}
func getPosts(url string) []posts {
	posts := make([]posts, 0)
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error %s", err)
		os.Exit(1)
	}
	json.Unmarshal(body, &posts)
	return posts
}

func (p posts) toString() string {
	bytes, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error in json marsheling")
		os.Exit(1)
	}
	return string(bytes)
}
