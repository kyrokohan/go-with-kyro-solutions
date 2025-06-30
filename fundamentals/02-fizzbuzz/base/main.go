package main

import (
	"flag"
	"fmt"
)

// Declare all flag variables
var (
	max int
)

// Capture the --max flag -- Go will run this automatically
func init() {
	flag.IntVar(&max, "max", 100, "Provide an upperbound")
}

func main() {
	// Good practice to do this here
	flag.Parse()

	// Print the values
	for i := 1; i <= max; i++ {
		// Two booleans to capture if Fizz, Buzz, or FizzBuzz
		divByThree := i%3 == 0
		divByFive := i%5 == 0

		if divByThree && divByFive {
			fmt.Println("FizzBuzz")
		} else if divByThree {
			fmt.Println("Fizz")
		} else if divByFive {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

}
