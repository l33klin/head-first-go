package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	URL  string
	Size int
}

func responseSize(url string, ch chan Page) {
	fmt.Println("Getting", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(len(body))
	ch <- Page{URL: url, Size: len(body)}
}

func main() {
	sizes := make(chan Page)
	urls := []string{
		"https://apple.com",
		"https://baidu.com",
		"https://taobao.com",
	}
	for _, url := range urls {
		go responseSize(url, sizes)
	}
	for i := 0; i < len(urls); i++ {
		fmt.Println(<-sizes)
	}
}
