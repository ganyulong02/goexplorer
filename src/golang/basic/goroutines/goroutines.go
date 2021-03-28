package main

import (
	"fmt"
	"runtime"
	// "time"
	"sync"
)

/*
Summary
Creating goroutines
	Use go keyword in front of function call
	When using anonymous functions, pass data as local variables
Synchronization
	Use sync.WaitGroup to wait for groups of goroutines to complete
	Use sync.Mutex and sync.RWMutex to protect data access
Parallelism
	By default, Go will use CPU threads equal to available cores
	Change with runtime.GOMAXPROCS
	More threads can increase performance, but too many can slow it down
Best practices:
Don't create goroutines in libraries
	Let consumer control concurrency
When creating a goroutine, know how it will end
	Avoids subtle memory leaks
Check for race conditions at compile time
	go run -race src/main.go
*/


var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{} // ReadWrite Mutex

func main() {
	//go sayHello()
	//time.Sleep(100 * time.Millisecond)
	//var msg = "Hello GoRoutines"
	//wg.Add(1)
	//go func(msg string) {
	//	fmt.Println(msg)
	//	wg.Done()
	//}(msg)
	//msg = "GoodBye"
	////time.Sleep(100 * time.Millisecond)
	//wg.Wait()
	runtime.GOMAXPROCS(16)
	//fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1))  // 8 -> 8 cores CPU
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go sayHello()
		m.Lock()
		go increment()
	}
	wg.Wait()
}

func sayHello() {
	fmt.Printf("Hello #%v\n", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	m.Unlock()
	wg.Done()
}
