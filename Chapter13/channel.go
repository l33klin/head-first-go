package main

import (
	"fmt"
	"time"
)

func main() {
	mc := make(chan string)
	go send(mc)
	reportNap("receiving goroutine", 5)
	fmt.Println(<-mc)
	fmt.Println(<-mc)
}

func reportNap(name string, delay int) {
	for i := 0; i < delay; i++ {
		fmt.Println(name, "sleeping")
		time.Sleep(time.Second)
	}
	fmt.Println(name, "wakes up!")
}

func send(mc chan string) {
	reportNap("sending goroutine", 2)
	fmt.Println("***sending value***")
	mc <- "a"
	fmt.Println("***sending value***")
	mc <- "b"
}
