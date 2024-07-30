package main

import (
	"fmt"
)

func main() {
	theAnswer := 4
	var result string

	if theAnswer < 0 {
		result = "less than 0"
	} else if theAnswer == 0 {
		result = "equal to zero"
	} else {
		result = "greater than zero"
	}
	fmt.Println(result)

	if a := -3; a < 0 {
		result = "less than zero"
	} else if a == 0 {
		result = "equal to zero"
	} else {
		result = "greater than 0"
	}
	fmt.Println(result)
}
