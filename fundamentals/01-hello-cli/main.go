package main

import (
	"flag"
	"fmt"
	"os"
)

// nameFlag stores the --name/-n argument
var nameFlag string

// Go will automatically run this function
func init() {
	// Set up the --name and -n and store them in our nameFlag variable
	flag.StringVar(&nameFlag, "name", "World", "Provide a name for personalized greetings!")
	flag.StringVar(&nameFlag, "n", "World", "Shorthand of --name")

	// Change the default --help/-h behavior
	flag.Usage = func() {
		fmt.Fprintf(
			flag.CommandLine.Output(),
			"\nHello-CLI - Prints A Friendly Greeting\n\n"+
				"Flags:\n")

		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr)
	}
}

func main() {
	// Parse the flags
	// We'll parse the flags here so init() stays side-effect free
	flag.Parse()

	// Finally print out the greeting message
	fmt.Printf("Hello, \033[36m%s\033[0m!\n", nameFlag)
}
