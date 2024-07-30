package main

import (
	"fmt"
)

func main() {
	doSomething()
}
func doSomething() {
	fmt.Println("playing tennis")
	sum := addValues(5, 9)
	fmt.Println("sum: ", sum)

	totalSum, totalCount := addAll(3, 2, 3, 4)
	fmt.Println("sum of multiple number: ", totalSum)
	fmt.Println("count: ", totalCount)
}
func addValues(value1 int, value2 int) int {
	return value1 + value2
}

func addAll(value ...int) (int, int) {
	total := 0
	for _, v := range value {
		total += v
	}
	return total, len(value)
}
