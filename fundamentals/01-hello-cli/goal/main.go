package main

import (
	"flag"
	"fmt"
)

// nameFlag stores the --name/-n argument
var nameFlag string

// Go will automatically run this function
func init() {
	// Set up the --name flag
	flag.StringVar(&nameFlag, "name", "World", "Provide a name for personalized greetings!")
}

func main() {
	// Parse the flags
	// We'll parse the flags here so init() stays side-effect free
	flag.Parse()

	fmt.Printf("Hello, %s!\n", nameFlag)

}
