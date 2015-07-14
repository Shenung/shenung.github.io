package main

import "fmt"

func Ping(p chan string) {
	for {
		msg := <-p
		fmt.Println(msg)
		c <- "ping"
	}
}

func pong(c chan string) {
	for {
		msg := <-p
		fmt.Println(msg)
		c <- "pong"
	}
}

func main() {
	c = make(chan string)
	c <- "start"
	go ping(c)
	go pong(c)

	fmt.Scanln()
}
