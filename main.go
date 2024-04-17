package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Post struct {
	UserId string `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type Posts struct {
	Post []Post `json:"post"`
}

func main() {
	resp, error := http.Get("https://jsonplaceholder.typicode.com/posts")
	if error != nil {
		panic(error)
	}
	defer resp.Body.Close()

	data := Posts{}
	err := json.NewDecoder(resp.Body).Decode(&data)
	if err == nil {
		log.Println(err)
		return
	}
	fmt.Println(data.Post)
}
