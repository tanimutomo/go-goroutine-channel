package main

import (
	"fmt"
	"time"
)

func main() {
	go printer("by goroutine")
	printer("by main")
	fmt.Println("done")
}

func printer(v string) {
	for i := 0; i < 3; i++ {
		fmt.Println(v)
		time.Sleep(3 * time.Second)
	}
}
