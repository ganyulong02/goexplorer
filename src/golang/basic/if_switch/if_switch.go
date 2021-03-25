package main

import (
	"fmt"
	"math"
)

func main() {
	statePopulations := make(map[string]int)
	statePopulations = map[string]int{
		"California": 21345,
		"Texas":      12345,
		"New York":   11345,
		"Ohio":       37489,
	}
	if pop, ok := statePopulations["Texas"]; ok {
		fmt.Println(pop)
	}

	guess_num()

	avoid_floating_point_error()
}

func guess_num() {
	num := 50
	guess := 30
	if guess < 1 {
		fmt.Println("The guess must be greater than 1!")
	} else if guess > 100 {
		fmt.Println("The guess must be less than 100!")
	} else {
		if guess < num {
			fmt.Println("Too low")
		}
		if guess > num {
			fmt.Println("Too high")
		}
		if guess == num {
			fmt.Println("You got it!")
		}
	}

	if guess < 1 || guess > 100 {
		fmt.Println("The guess must be between 1 and 100")
	}
}

func avoid_floating_point_error() {
	myNum := 0.123
	if math.Abs(myNum/math.Pow(math.Sqrt(myNum), 2)-1) < 0.001 {
		fmt.Println("These are the same")
	} else {
		fmt.Println("These are different!")
	}
}

func switch_demo() {
	switch i := 2 + 3; i {
	case 1, 5, 10:
		fmt.Println("one, five, ten")
	case 2, 4, 6:
		fmt.Println("two, four, six")
	default:
		fmt.Println("another number")
	}
}

func tagless_switch_example() {
	i := 9
	switch {
	case i <= 10:
		fmt.Println("less than or equal to ten")
		fallthrough // you are taking control of the flow, it will literally fallthrough, be careful
	case i <= 20:
		fmt.Println("less than or equal to twenty")
	default:
		fmt.Println("greater than twenty")
	}
}

func type_switch_example() {
	// var i interface{} = 1.0
	var i interface{} = [3]int{} // array of size 3
	switch i.(type) {
	case int:
		fmt.Println("i is an integer")
		break
		// fmt.Println("This will print too")
	case float64:
		fmt.Println("i is a float64")
	case string:
		fmt.Println("i is string")
	default:
		fmt.Println("i is another type")
	}
}
