package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Post struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var posts []Post
	err = json.NewDecoder(resp.Body).Decode(&posts)
	if err != nil {
		log.Fatal(err)
	}

	// Accede a los elementos del array de posts
	for _, post := range posts {
		fmt.Printf("UserID: %d, ID: %d, Title: %s\n", post.UserId, post.Id, post.Title)
	}
}
