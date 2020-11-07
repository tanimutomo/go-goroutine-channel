package main

import "fmt"

func main() {
	msgs := make(chan string)
	go func() { msgs <- "hello" }()
	msg := <-msgs
	fmt.Println(msg)
}
