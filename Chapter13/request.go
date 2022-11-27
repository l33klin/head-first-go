package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	ch := make(chan int)
	go responseSize("https://apple.com", ch)
	go responseSize("https://baidu.com", ch)
	go responseSize("https://taobao.com", ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
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
