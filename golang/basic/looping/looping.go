package main

import (
	"fmt"
)

func main() {
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// }

	// // i++ is not an expression, is a statement in Go
	// for i, j := 0, 0; i < 5; i, j = i+1, j+1 {
	// 	fmt.Println(i, j)
	// }

	// i := 0
	// for i < 5 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// j := 0
	// for {
	// 	fmt.Println(j)
	// 	if j == 5 {
	// 		break
	// 	}
	// }
Loop:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i * j)
			if i*j >= 3 {
				break Loop
			}
		}
	}

	/*
		Looping over collections
		arrays, slices, maps, strings, channels
		for k, v := range collection {}
	*/

	s := []int{1, 2, 3}
	for k, v := range s {
		fmt.Println(k, v)
	}

	statePopulations := make(map[string]int)
	statePopulations = map[string]int{
		"California": 21345,
		"Texas":      12345,
		"New York":   11345,
		"Ohio":       37489,
	}

	for _, v := range statePopulations {
		fmt.Println(v)
	}

}
