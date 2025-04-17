package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Veer")
	fmt.Println("Simple Calculator")
	fmt.Println("=================")
	fmt.Println("Available operations: +, -, *, /, %")
	fmt.Println("Type 'exit' to quit")
	fmt.Println()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		input := scanner.Text()
		if strings.ToLower(input) == "exit" {
			fmt.Println("Goodbye!")
			break
		}

		result, err := calculate(input)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			continue
		}

		fmt.Printf("Result: %g\n", result)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Reading standard input: %v\n", err)
	}
}

func calculate(input string) (float64, error) {
	// Parse the input string
	input = strings.TrimSpace(input)
	
	// Check for operators
	var operator string
	var parts []string
	
	if strings.Contains(input, "+") {
		operator = "+"
		parts = strings.Split(input, "+")
	} else if strings.Contains(input, "-") {
		operator = "-"
		parts = strings.Split(input, "-")
	} else if strings.Contains(input, "*") {
		operator = "*"
		parts = strings.Split(input, "*")
	} else if strings.Contains(input, "/") {
		operator = "/"
		parts = strings.Split(input, "/")
	} else if strings.Contains(input, "%") {
		operator = "%"
		parts = strings.Split(input, "%")
	} else {
		return 0, fmt.Errorf("no valid operator found")
	}

	// Check if we have two numbers
	if len(parts) != 2 {
		return 0, fmt.Errorf("invalid expression format")
	}

	// Parse the numbers
	num1, err := strconv.ParseFloat(strings.TrimSpace(parts[0]), 64)
	if err != nil {
		return 0, fmt.Errorf("invalid first number: %v", err)
	}

	num2, err := strconv.ParseFloat(strings.TrimSpace(parts[1]), 64)
	if err != nil {
		return 0, fmt.Errorf("invalid second number: %v", err)
	}

	// Perform the calculation
	switch operator {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		return num1 / num2, nil
	case "%":
		if num2 == 0 {
			return 0, fmt.Errorf("modulo by zero")
		}
		return float64(int(num1) % int(num2)), nil
	default:
		return 0, fmt.Errorf("unknown operator: %s", operator)
	}
}