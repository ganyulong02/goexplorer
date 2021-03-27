package main

import (
	"fmt"
)


/*
	Parameters
	* Comma delimited list of variables and types
		func foo(bar string, baz int)
	* Parameters of same type list type once
		func foo(bar, baz int)
	* When pointers are passed in, the function can change the value in the caller
		This is always true for data of slices and maps
	* Use variadic parameters to send list of same types in
		Must be last parameter
		Received as a slice
		func foo(bar string, baz ...int)

	Return values
	* Single return values just list type
	* Multiple return value list types surrounded by parentheses
		The (result type, error) paradigm is a very common idiom
	* Can use named return values
	* Can return addresses of local variables
	Automatically promoted from local memory (stack) to shared memory (heap)

	Methods
	Function that executes in context of a type (struct)
	Format
		func (g greeter) greet() {
			...
		}
	Receiver can be value or pointer
		Value receiver gets copy of type
		Pointer receiver gets pointer to type
 */


func main() {
	greeting := "Hello"
	name := "Jason"
	sayMessage(&greeting, &name)
	fmt.Println(name)

	// anonymous functions
	func() {
		fmt.Println("Anonymous func")
	}() // '()' means immediate invoke the func

	s := sum(1, 2, 3, 4, 5)
	fmt.Println("The sum is: ", *s)
	//d, err := _divide(5.0, 0.0)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(d)

	// Method
	g := greeter{
		greeting: "Hello",
		name: "Go",
	}
	g.greet()

	// declare a func as variable
	var divide func(float64, float64) (float64, error)
	divide = func(a,b float64) (float64, error) {
		if b == 0.0 {
			return 0.0, fmt.Errorf("Cannot divide by zero")
		} else {
			return a / b, nil
		}
	}
	d, err := divide(6.0, 0.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)

}

type greeter struct {
	greeting string
	name string
}

func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
}

func sayMessage(greeting, name *string) {
	fmt.Println(*greeting, *name)
	*name = "Ted"
	fmt.Println(*name)
}

func sum(values ...int) *int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return &result // local stack -> shared heap
}

func _divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}


