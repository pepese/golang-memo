package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan struct{})
	/*
		var input chan<- struct{}
		var output <-chan struct{}
		input = c
		output = c
	*/

	go func() {
		fmt.Println("Starting goruoutine.")
		defer fmt.Println("Ending goruoutine.")
		time.Sleep(1 * time.Second)
		// input <- struct{}{}
		c <- struct{}{}
	}()

	//<-output
	<-c
	fmt.Println("Main goroutine ends.")
}
