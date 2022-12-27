package main

import (
	"fmt"
)

// var result uint

func main() {

	greetUser()

	x, y, ops := userInput()

	result := getresult(x, y, ops)

	print("Result of the operation:", result, "\n")
}

func greetUser() {
	fmt.Println("Welcome to the calculator")
	fmt.Println("Please enter the values and the operation you want to perform")
}

func userInput() (uint, uint, string) {
	var variable_x uint
	var variable_y uint
	var operation string

	fmt.Println("Enter the first variable")
	fmt.Scan(&variable_x)

	fmt.Println("Enter the second variable")
	fmt.Scan(&variable_y)

	fmt.Println("Enter the operation")
	fmt.Scan(&operation)

	return variable_x, variable_y, operation
}

func getresult(x uint, y uint, ops string) uint {
	var result uint

	if ops == "+" {
		result = x + y
	} else if ops == "-" {
		result = x - y
	} else if ops == "*" {
		result = x * y
	} else if ops == "/" {
		result = x / y
	} else {
		fmt.Println("Operation entered is incorrect")
	}

	return result
}
