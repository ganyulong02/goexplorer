package main

import (
	"log"
	"os"
	// "_" before the matchers package path is called blank identifier.
	// The blank identifier allows the compiler to accept the import and call any init
	// functions that can be found in the different code files within that package.
	_ "go-in-action/chapter2/sample/matchers"
	"go-in-action/chapter2/sample/search"
)

// init is called prior to main.
func init() {
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)
}

// main is the entry point for the program.
func main() {
	// Perform the search for the specified term.
	search.Run("president")
}
