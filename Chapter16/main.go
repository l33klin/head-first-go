package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/guestbook", guestBookHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
