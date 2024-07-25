package main

import (
	"fmt"
	"strconv"
)

func calculate(valuel string, value2 string) float64 {

	float1, err1 := strconv.ParseFloat(valuel, 64)
	if err1 != nil {
		fmt.Println("Error parsing valuel:", err1)
		return 0
	}

	float2, err2 := strconv.ParseFloat(value2, 64)
	if err2 != nil {
		fmt.Println("Error parsing value2:", err2)
		return 0
	}

	return float1 + float2
}

func main() {

	valuel := "10.0"
	value2 := "5.5"
	result := calculate(valuel, value2)
	fmt.Printf("Result: %.6f\n", result)

	valuel = "100.0"
	value2 = "-110.5"
	result = calculate(valuel, value2)
	fmt.Printf("Result: %.6f\n", result)
}

/*
This defines a calculate function that takes two string inputs,
parses them into float64 values, sums them, and returns the result.
 If parsing fails, it prints an error message and returns 0. The main
  function demonstrates the usage of
calculate with example inputs and prints the results formatted to six decimal places.
*/
