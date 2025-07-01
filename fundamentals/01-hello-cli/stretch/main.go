package main

import (
	"flag"
	"fmt"
	"os"
)

const version = "0.1.0"

// nameFlag stores the --name/-n argument
var (
	nameFlag    string
	versionFlag bool
)

// Go will automatically run this function
func init() {
	// Set up the --name and -n and store them in our nameFlag variable
	flag.StringVar(&nameFlag, "name", "World", "Provide a name for personalized greetings!")
	flag.StringVar(&nameFlag, "n", "World", "Shorthand of --name")
	flag.BoolVar(&versionFlag, "version", false, "Print the current version of the app.")

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

func greet(name string, noColor bool) string {
	if noColor {
		return fmt.Sprintf("Hello, %s!", name)
	}
	return fmt.Sprintf("Hello, \033[36m%s\033[0m!", name)
}

func main() {
	// Parse the flags
	// We'll parse the flags here so init() stays side-effect free
	flag.Parse()

	// Runs on if --version flag is given
	if versionFlag {
		fmt.Printf("hello-cli v%s\n", version)
		return
	}

	// Respect the NO_COLOR environment variable
	fmt.Println(greet(nameFlag, os.Getenv("NO_COLOR") != ""))

}
