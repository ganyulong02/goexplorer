package main

import (
	"fmt"
	"strconv"
)

/*
Variable Summary
1. Variable declaration
	var foo int
	var foo int = 42
	foo := 42
2. Can't redeclare variables, but can shadow them
3. All variables must be used
4. Visibility
	lower case first letter for package scope
	upper case first letter to export
	no private scope
5. Naming conventions
	Pascal or camelCase
		Capitalize acronyms(HTTP, URL)
	As short as reasonable
		longer names for longer lives
6. Type conversions
	destinationType(variable)
	use strconv package for strings
*/

// var m float32 = 32.

// var (
// 	counter1 int = 1
// 	counter2 int = 2
// )

func main() {
	var i int = 42
	var j string = strconv.Itoa(i)
	// k := 99
	// fmt.Println(i)
	fmt.Printf("%v, %T", j, j)
}
