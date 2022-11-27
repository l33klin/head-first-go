package main

import (
	"log"
	"net/http"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/guestbook", guestBookHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
