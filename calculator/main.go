package main

import (
	"fmt"
	"log"
)

func main() {
	var operator string
	var num1, num2 float64

	fmt.Println("Enter first number:")
	_, err := fmt.Scan(&num1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Enter operator (+, -, *, /):")
	_, err = fmt.Scan(&operator)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Enter second number:")
	_, err = fmt.Scan(&num2)
	if err != nil {
		log.Fatal(err)
	}

	switch operator {
	case "+":
		fmt.Printf("Result: %.2f\n", num1+num2)
	case "-":
		fmt.Printf("Result: %.2f\n", num1-num2)
	case "*":
		fmt.Printf("Result: %.2f\n", num1*num2)
	case "/":
		if num2 != 0 {
			fmt.Printf("Result: %.2f\n", num1/num2)
		} else {
			fmt.Println("Error: Division by zero is not allowed.")
		}
	default:
		fmt.Println("Error: Invalid operator")
	}
}
