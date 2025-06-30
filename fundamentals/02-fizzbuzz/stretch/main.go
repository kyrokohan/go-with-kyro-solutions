package main

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Declare ANSI color codes
const (
	ResetColor  = "\033[0m"
	GreenColor  = "\033[32m"
	YellowColor = "\033[33m"
)

// Declare all flag variables
var (
	divisors [2]int
	max      int
	format   string
	color    bool
)

// Capture the flags -- Go will run this automatically
func init() {
	flag.IntVar(&max, "max", 100, "Provide an upperbound")
	flag.IntVar(&divisors[0], "fizz", 3, "Provide the divisor for \"Fizz\"")
	flag.IntVar(&divisors[1], "buzz", 5, "Provide the divisor for \"Buzz\"")
	flag.StringVar(&format, "format", "plain", "Provide the output format (plain, csv, or json)")
	flag.BoolVar(&color, "color", false, "Include ANSI color codes!")
}

// Helper function to exit and report errors
func exitWithError(err error) {
	fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	os.Exit(1)
}

// Ensure format flag is within what we allow
func validateFormat(fmt string) error {
	for _, v := range [3]string{"plain", "csv", "json"} {
		if fmt == v {
			return nil
		}
	}
	return errors.New("invalid format flag set!")
}

// Run the main FizzBuzz logic and return the results
func fizzBuzz() []string {
	// Initialize a slice of length `max`.
	// This is fine with small values, but to allow larger
	// values, a streaming approach is more appropriate
	slice := make([]string, max)

	for i := 0; i < max; i++ {
		// Two booleans to capture if Fizz, Buzz, or FizzBuzz
		divByFirst := (i+1)%divisors[0] == 0
		divBySecond := (i+1)%divisors[1] == 0

		if divByFirst && divBySecond {
			slice[i] = "FizzBuzz"
		} else if divByFirst {
			slice[i] = "Fizz"
		} else if divBySecond {
			slice[i] = "Buzz"
		} else {
			slice[i] = strconv.Itoa(i + 1)
		}
	}

	return slice
}

func printPlain(list []string) {
	// I'm setting up color only for plain printing.
	for _, v := range list {
		if color {
			switch v {
			case "Fizz":
				fmt.Println(GreenColor + v + ResetColor)
			case "Buzz":
				fmt.Println(YellowColor + v + ResetColor)
			case "FizzBuzz":
				fmt.Println(GreenColor + "Fizz" + YellowColor + "Buzz" + ResetColor)
			default:
				fmt.Println(v)
			}
		} else {
			fmt.Println(v)
		}
	}
}

func printCSV(list []string) {
	w := csv.NewWriter(os.Stdout)
	_ = w.Write(list)
	w.Flush()
	if err := w.Error(); err != nil {
		exitWithError(err)
	}
}

func printJSON(list []string) {
	b, err := json.Marshal(list)
	if err != nil {
		exitWithError(err)
	}
	fmt.Println(string(b))
}

func main() {
	// Good practice to do this here
	flag.Parse()

	// Catch and throw an error if either of the divisors are 0 or less
	if divisors[0] <= 0 || divisors[1] <= 0 {
		exitWithError(errors.New("divisor(s) cannot be less than or equal to 0!"))
	}

	// Also make sure max is set to a positive integer
	if max <= 0 {
		exitWithError(errors.New("--max must be larger than 0!"))
	}

	// Turn format flag into lowercase so the flag becomes case insensitive
	format = strings.ToLower(format)

	if err := validateFormat(format); err != nil {
		exitWithError(err)
	}

	result := fizzBuzz()

	// Based on the format, print accordingly
	switch format {
	case "plain":
		printPlain(result)
	case "csv":
		printCSV(result)
	case "json":
		printJSON(result)
	}

}
