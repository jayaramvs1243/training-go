package main

import (
	"errors"
	"fmt"
)

/*
	4. Write a program that checks if a given number is even or odd
	5. Check whether a given year is a leap year
	6. Implement a simple calculator using `if-else` statements
*/

func main() {

	// Write a program that checks if a given number is even or odd
	inputNumber := 54
	if isEven(inputNumber) {
		fmt.Printf("%d is even.\n", inputNumber)
	} else {
		fmt.Printf("%d is odd.\n", inputNumber)
	}

	// Check whether a given year is a leap year
	year := 2005
	if isLeapYear(year) {
		fmt.Printf("%d is a leap year.\n", year)
	} else {
		fmt.Printf("%d is not a leap year.\n", year)
	}

	fmt.Println("***SIMPLE CALCULATOR***")

	fmt.Println("Enter the first number: ")
	var input1 float64
	fmt.Scan(&input1)

	fmt.Println("Enter the second number: ")
	var input2 float64
	fmt.Scan(&input2)

	fmt.Println("Input an operator(+, -, * and /): ")
	var operator string
	fmt.Scan(&operator)

	result, err := calculate(input1, input2, operator)
	if err != nil {
		fmt.Printf(err.Error())
	} else {
		fmt.Printf("Result is: %f\n", result)
	}

	result, err = calculateUsingSwitch(input1, input2, operator)
	if err != nil {
		fmt.Printf(err.Error())
	} else {
		fmt.Printf("Result (using switch) is: %f\n", result)
	}

}

func isEven(inputNumber int) bool {
	return inputNumber%2 == 0
}

func isLeapYear(year int) bool {
	return year%4 == 0
}

func calculate(input1 float64, input2 float64, operator string) (float64, error) {
	if operator == "+" {
		return input1 + input2, nil
	} else if operator == "-" {
		return input1 - input2, nil
	} else if operator == "*" {
		return input1 * input2, nil
	} else if operator == "/" {
		if input2 == 0 {
			return 0, errors.New("Division by zero is not allowed")
		} else {
			return input1 / input2, nil
		}
	} else {
		return 0, errors.New("Invalid operator")
	}
}

func calculateUsingSwitch(input1 float64, input2 float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return input1 + input2, nil
	case "-":
		return input1 - input2, nil
	case "*":
		return input1 * input2, nil
	case "/":
		if input2 == 0 {
			return 0, errors.New("Division by zero is not allowed")
		} else {
			return input1 / input2, nil
		}
	default:
		return 0, errors.New("Invalid operator")
	}
}
