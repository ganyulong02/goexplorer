package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan bool)
	go func() {
		time.Sleep(1 * time.Second)
		<-c
	}()
	start := time.Now()
	c <- true // blocks until other goroutine receives
	fmt.Printf("send took %v\n", time.Since(start))
}

func deadlockexample() {
	c := make(chan bool)
	// d := make(chan bool, 1) buffered channel
	<-c
	c <- true
}
