package main

/*
constant
Immutable, but inner scope(func level) constant can shadow the outer scope(package level)
Replaced by compiler at compile time, value must be calculable at compile time
Named like variables
	PascalCase for exported constants
	camelCase for internal constants
Typed constants work like immutable variables, can interoperate only with same type
Untyped constants work like literals, can interoperate only with similar type
Enumerated constants
	Special symbol iota allows related constants to be created easily
	Iota starts at 0 in each const block and increments by one
	Watch out of constant values that match zero values for variables
Enumerated expressions
	Operations that can be determined at compile time are allowed
		Arithmetic	Bitwise operation	Bitshifting
*/
import (
	"fmt"
)

const (
	a = iota
	b
	c
)

const (
	a2 = iota
)

const (
	_ = iota
	catSpec
	dogSpec
	snakeSpec
)

const (
	_  = iota             // ignore first value by assigning to blank identifier
	KB = 1 << (10 * iota) // 2 ^ 10
	MB                    // 2 ^ 100
	GB
	TB
	PB
	EB
	ZB
	YB
)

const (
	isAdmin = 1 << iota
	isHeadquarters
	canSeeFinancials

	canSeeAfrica
	canSeeAsia
	canSeeEurope
)

func main() {
	const innerConst int = 1
	// const OuterConst int = 2
	fmt.Printf("%v, %T\n", innerConst, innerConst)
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", a2)

	var specType int
	fmt.Printf("%v\n", specType == catSpec)

	// file size in GB
	fileSize := 4000000000.
	fmt.Printf("%.2fGB", fileSize/GB)

	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("%b\n", roles)
	fmt.Printf("Is Admin? %v\n", isAdmin&roles == isAdmin)
	fmt.Printf("Is HQ? %v\n", isHeadquarters&roles == isHeadquarters)

}
