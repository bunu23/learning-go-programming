package main

import (
	"fmt"
)

func main() {
	colors := []string{"red", "blue"}
	fmt.Println(colors)

	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}
	for i := range colors {
		fmt.Println(colors[i])
	}

	for _, color := range colors {
		fmt.Println(color)
	}

	value := 11
	for value < 13 {
		fmt.Println("Value:", value)
		value++
	}

	sum := 1
	for sum < 1000 {
		sum += sum
		fmt.Println("sum:", sum)
		if sum < 100 {
			goto theEnd
		}
	}

theEnd:
	fmt.Println("end of program")
}
