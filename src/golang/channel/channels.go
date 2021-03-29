package main

import (
	"fmt"
	"sync"
	"time"
)



/*
Channel basic
	Create a channel with make command
		make(chan int)
	Send message into channel
		ch<- val
	Receive message from channel
		val := <-ch
	Can have multiple senders and receivers
Restricting data flow
	Channel can be cast into send-only or receive-only versions
		Send-only: chan<- int
		Receive-only: <-chan int
Buffered channels
	Channels block sender side till receiver is available
	Block receiver side till message is available
	Can decouple sender and receiver with buffered channels
		make(chan int, 50)
	Use buffered channels when send and receiver have asymmetric loading
For...range loops with channels
	Use to monitor channel and process messages as they arrive
	Loops exists when channel is closed
Select statements
	Allow goroutine to monitor several channels at once
		Blocks if all channels block
		If multiple channels receive value simultaneously, behavior is undefined
*/

var wg = sync.WaitGroup{}

const (
	logInfo    = "INFO"
	logWarning = "WARNING"
	logError   = "ERROR"
)

type logEntry struct {
	time     time.Time
	severity string
	message  string
}

var logCh = make(chan logEntry, 50)
var doneCh = make(chan struct{}) // signal channel, struct with zero fields require no memory allocation

func logger() {
	for {
		select {
		case entry := <-logCh:
			fmt.Printf("%v - [%v]%v\n",
				entry.time.Format("2006-01--2T15:04:05"),
				entry.severity,
				entry.message)
		case <-doneCh:
			break
		// default:
			// some operation, non-blocking
		}
	}
}

func main() {
	go logger()
	//defer func(){
	//	close(logCh)
	//}()
	logCh <- logEntry{time.Now(), logInfo, "App is starting"}

	logCh <- logEntry{time.Now(), logInfo, "App is shutting down"}
	time.Sleep(100 * time.Millisecond)
	doneCh <- struct{}{}
	//exampleOne()
	//exampleTwo()
	//channelWithDirection()
	//channelWithBuffer()
}

func exampleOne() {
	ch := make(chan int)
	wg.Add(2)
	go func() {
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}()
	go func() {
		ch <- 42
		wg.Done()
	}()
	wg.Wait()
}

func exampleTwo() {
	ch := make(chan int)
	for j := 0; j < 5; j++ {
		wg.Add(2)
		go func() {
			i := <-ch
			fmt.Println(i)
			wg.Done()
		}()
		go func() {
			ch <- 42
			wg.Done()
		}()
	}
	wg.Wait()
}

func channelWithDirection() {
	ch := make(chan int)
	wg.Add(2)
	// receive only goroutine, send data into it
	go func(ch <-chan int) {
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}(ch)
	// send only goroutine, send data out of it
	go func(ch chan<- int) {
		ch <- 42
		wg.Done()
	}(ch)
	wg.Wait()
}

func channelWithBuffer() {
	ch := make(chan int, 50)
	wg.Add(2)

	// send only goroutine, send data out of it
	go func(ch chan<- int) {
		ch <- 42
		ch <- 27
		wg.Done()
	}(ch)

	// receive only goroutine, send data into it
	go func(ch <-chan int) {
		i := <-ch
		fmt.Println(i)
		i = <-ch
		fmt.Println(i)
		wg.Done()
	}(ch)
	wg.Wait()
}

func commaOkSyntaxChannel() {
	ch := make(chan int, 50)
	wg.Add(2)
	// receive only goroutine, send data into it
	go func(ch <-chan int) {
		for {
			if i, ok := <-ch; ok {
				fmt.Println(i)
			} else {
				break
			}
		}
		wg.Done()
	}(ch)
	// send only goroutine, send data out of it
	go func(ch chan<- int) {
		ch <- 42
		ch <- 27
		close(ch)
		wg.Done()
	}(ch)
	wg.Wait()
}

func forRangeChannel() {
	ch := make(chan int, 50)
	wg.Add(2)
	// receive only goroutine, send data into it
	go func(ch <-chan int) {
		for i := range ch {
			fmt.Println(i)
		}
		wg.Done()
	}(ch)
	// send only goroutine, send data out of it
	go func(ch chan<- int) {
		ch <- 42
		ch <- 27
		close(ch)
		wg.Done()
	}(ch)
	wg.Wait()
}
