package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	target := rand.Intn(100) + 1
	chances := 10
	for chances > 0 {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Please input a number(should less than 100): ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		guessNum, err := strconv.Atoi(strings.TrimSpace(input))
		if err != nil {
			log.Fatal(err)
		}
		if guessNum > target {
			fmt.Println("The number your input is too big!")
		} else if guessNum < target {
			fmt.Println("The number your input is too little!")
		} else {
			fmt.Println("Cool! You got the number!")
			break
		}
		chances--
		if chances <= 0 {
			fmt.Printf("You've out of chance, the target is: %v.\n", target)
			break
		}
		fmt.Printf("Your have %d chances to guess!", chances)
	}
}
