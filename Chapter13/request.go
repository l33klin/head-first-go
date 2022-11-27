package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	sizes := make(chan int)
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

func responseSize(url string, ch chan int) {
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
	ch <- len(body)
}
