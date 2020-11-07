package main

import (
	"fmt"
)

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, opened := <-jobs
			if opened {
				fmt.Println("received job", j)
				continue
			}
			fmt.Println("received all jobs")
			done <- true
			return
		}
	}()

	for j := 1; j <= 5; j++ {
		jobs <- j
		fmt.Println("send job", j)
	}

	close(jobs)
	fmt.Println("sent all jobs")

	<-done
}
