package main

import (
	"errors"
	"fmt"
	"math"
)

type person struct {
	name string
	age  int
}

func inc(x *int) {
	*x++
}

func main() {
	i := 7
	inc(&i)
	fmt.Println(i)
	fmt.Println(&i)

	p := person{name: "Jason", age: 23}
	fmt.Println(p)
	fmt.Println(p.age)

	fmt.Println("Hello World!")
	x := 5
	y := 7
	sum := x + y
	fmt.Println(sum)

	var arr [5]int
	arr[2] = 7

	arr1 := [5]int{5, 4, 3, 2, 1}
	arr2 := []int{1, 2, 3}
	arr2 = append(arr2, 4)
	fmt.Println(arr)
	fmt.Println(arr1)
	fmt.Println(arr2)

	vertices := make(map[string]int)
	vertices["triangle"] = 2
	vertices["square"] = 3
	vertices["dodecagon"] = 12
	delete(vertices, "square")

	fmt.Println(vertices)

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	str := []string{"a", "b", "c"}
	m := make(map[string]string)
	m["a"] = "alpha"
	m["b"] = "beta"

	for index, value := range str {
		fmt.Println("index", index, "value", value)
	}
	for key, value := range m {
		fmt.Println("key", key, "value", value)
	}

	res := sum1(2, 3)
	fmt.Println(res)

	result, err := sqrt(16)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

// the sum of two integers
func sum1(x int, y int) int {
	return x + y
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for negative numbers")
	}
	return math.Sqrt(x), nil
}
