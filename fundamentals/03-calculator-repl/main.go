package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func eval(line string) {
	var operatorIndex int = -1
	var operator string

	// Figure out what and where the operator is
	for i, ch := range line {
		if ch == '+' || ch == '-' || ch == '*' || ch == '/' {
			operatorIndex = i
			operator = string(ch)
			break
		}
	}

	// If no operator is found, there's nothing more to do
	if operatorIndex == -1 {
		fmt.Println("error: invalid expression")
		return
	}

	// Extract both sides of the expression
	lhs, e1 := strconv.Atoi(strings.TrimSpace(line[:operatorIndex]))
	rhs, e2 := strconv.Atoi(strings.TrimSpace(line[operatorIndex+1:]))

	// Catch non-integer tokens
	if e1 != nil || e2 != nil {
		fmt.Println("error: invalid expression")
		return
	}

	// Catch division by zero issue
	if operator == "/" && rhs == 0 {
		fmt.Println("error: division by zero")
		return
	}

	// Perform the operation
	var result int
	switch operator {
	case "+":
		result = lhs + rhs
	case "-":
		result = lhs - rhs
	case "*":
		result = lhs * rhs
	case "/":
		result = lhs / rhs
	}

	fmt.Println(result)
}

func repl(scanner *bufio.Scanner) {
	for {
		fmt.Printf("> ")

		// If there's an error or interrupt signal, we'll break out
		if !scanner.Scan() {
			break
		}

		// Remove leading and trailing spaces
		line := strings.TrimSpace(scanner.Text())

		if line == "exit" || line == "quit" {
			fmt.Println("Bye!")
			break
		}

		// Use eval() to evaluate the line
		eval(line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "input error:", err)
	}
}

func main() {
	// Define a new scanner to read from standard input
	scanner := bufio.NewScanner(os.Stdin)

	repl(scanner)
}
