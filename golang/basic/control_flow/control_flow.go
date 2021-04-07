package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/*
	Defer, Panic, Recover

	Defer
	Used to delay execution of a statement until function exits
	Defer move the defered function after the main but before the main exit
	Defer runs in LIFO order, the function gets lastly defered, will be firstly called
	Arguments evaluated at the time defer is executed, not at the time of called function execution
	Defer is normally used in closing resources

	Panic
	Occur when program cannot continue at all
		Don't use when file can't be opened, unless it is critical
		Use for unrecoverable events - cannot obtain TCP port for web server
	Function will stop executing
		Deferred functions will still fire
	If nothing handels panic, program will exit

	Recover
		Used to recover from panics
		Only useful in deferred functions
		Current function will not attemp to continue, but higher functions in call stack will.
*/

func main() {
	// fmt.Println("start")
	// defer fmt.Println("middle")
	// panic("something bad happened")
	// fmt.Println("end")
	// simple()
	// panicing()
	fmt.Println("start")
	panicker()
	fmt.Println("end")

}

func panicker() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error: ", err)
			panic(err) // rethrow the panic
		}
	}()
	panic("something bad happened")
	fmt.Println("done panicking")
}

func simple() {
	a := "start"
	defer fmt.Println(a) // defer will eagerlly evaluate, so will print start not end
	a = "end"
}

func httpCall() {
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	robots, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}

func panicing() {
	// Go doesn't have exception
	// a, b := 1, 0
	// ans := a / b
	// fmt.Println(ans)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Go!"))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	}
}
