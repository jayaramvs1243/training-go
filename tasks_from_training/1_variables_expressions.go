package main

import (
	"fmt"
)

/*
	1. Declare an `int`, `string`, and `bool` variable and print their values
	2. Swap two integer variables without using a temporary variable
	3. Write a program to calculate the area of a rectangle using variables
*/

// Package level variables
// Declaration & initialization
var intVariable int

var strVariable string = "string"
var boolVariable bool = true
var (
	firstName string = "Jayaram"
	lastName  string = "V S"
)

// Declaration & initialization with type inference
const placeOfBirth = "Kanyakumari"
const (
	gender      = "male"
	yearOfBirth = 1994
)

func main() {
	intVariable = 0

	println(intVariable)
	println(strVariable)
	println(boolVariable)

	println(firstName)
	println(lastName)

	fmt.Println("Printing constants:", placeOfBirth, gender, yearOfBirth)

	intVariable = 1
	fmt.Printf("Value of 'intVariable' changed to %d\n", intVariable)

	// Declare and initialize a new variable in one step
	// Only works inside functions; not at the package level
	testVariable := "test"
	fmt.Println(testVariable)

	// Complex types
	// complex64 & complex128
	var complexVariable complex128 = 3.5 + 4.2i
	fmt.Printf("Complex variable: %v. Real: %v. Image: %v.\n", complexVariable, real(complexVariable), imag(complexVariable))

	// Swap two integer variables without using a temporary variable
	a := 8
	b := 3
	fmt.Println("(a, b) before swapping:", a, b)

	a = a + b
	b = a - b
	a = a - b
	fmt.Println("(a, b) after swapping:", a, b)

	// Write a program to calculate the area of a rectangle using variables
	length := 10.0
	breadth := 5.0
	areaOfRectangle := areaOfRectangle(length, breadth)
	fmt.Printf("Area of rectangle of length %f and breadth %f is %f", length, breadth, areaOfRectangle)

}

func areaOfRectangle(length float64, breadth float64) float64 {
	return (length * breadth)
}
