package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)

	d := "data"

	go send(d, c1)
	go relay(c1, c2)
	go receive(c2)

	time.Sleep(1 * time.Second)
}

func send(d string, c chan<- string) {
	fmt.Printf("send: %s to c1\n", d)
	c <- d
}

func relay(c1 <-chan string, c2 chan<- string) {
	d := <-c1
	fmt.Printf("receive: %s from c1\n", d)
	c2 <- d
	fmt.Printf("send: %s to c2\n", d)
}

func receive(c2 <-chan string) {
	d := <-c2
	fmt.Printf("receive: %s from c2\n", d)
}
