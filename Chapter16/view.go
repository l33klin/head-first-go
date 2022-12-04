package main

import (
	"bufio"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type Guestbook struct {
	SignatureCount int
	Signatures     []string
}

func guestBookHandler(w http.ResponseWriter, r *http.Request) {
	signatures := getStrings("signatures.txt")
	html, err := template.ParseFiles("view.html")
	check(err)
	guestbook := Guestbook{
		SignatureCount: len(signatures),
		Signatures:     signatures,
	}
	err = html.Execute(w, guestbook)
	check(err)
}

func newHandler(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("new.html")
	err = html.Execute(w, nil)
	check(err)
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	signature := r.FormValue("signature")
	if signature != "" {
		// write to file
		options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
		file, err := os.OpenFile("signatures.txt", options, os.FileMode(0600))
		check(err)
		_, err = fmt.Fprintln(file, signature)
		check(err)
		err = file.Close()
		check(err)
	}
	http.Redirect(w, r, "/guestbook", http.StatusFound)
}

func getStrings(fileName string) []string {
	var lines []string
	file, err := os.Open(fileName)
	if os.IsNotExist(err) {
		return nil
	}
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	check(scanner.Err())
	return lines
}
