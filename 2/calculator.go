package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var operator string

	fmt.Println(" Simple Calculator")
	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)

	fmt.Print("Enter operator (+, -, *, /, %): ")
	fmt.Scan(&operator)

	fmt.Print("Enter second number : ")
	fmt.Scan(&num2)

	switch operator{
	case "+":
		fmt.Printf("Result: %.2f\n", add(num1, num2))
	case "-":
		fmt.Printf("Result:%.2f\n", subtract(num1, num2))
	case "*":
		fmt.Printf("Result:%.2f\n", multiply(num1, num2))
	case "/":
		if num2 == 0 {
			fmt.Println("Error: Division by zero")
		} else {
			fmt.Printf("Result: %.2f\n", divide(num1, num2))
		}
	case "%":
		fmt.Printf("Result: %.0f\n", float64(int(num1)%int(num2)))
	default:
		fmt.Println("Invalid operator.")
	}
}

func add(a, b float64) float64 {
	return a + b
}

func subtract(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}

func divide(a, b float64) float64 {
	return a / b
}
