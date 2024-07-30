package main

import (
	"fmt"
	"sort"
)

func main() {

	var number = [4]int{4, 3, 2, 3}
	fmt.Println(number)

	fmt.Println("number : ", len(number))

	// var colors [3]string
	// colors[0] = "blue"
	// colors[1] = "blue"
	// colors[2] = "blue"
	// fmt.Println(colors[0])
	// fmt.Println(colors)

	var colors = []string{"red", "green", "blue"}
	fmt.Println(colors)
	colors = append(colors, "purple")
	fmt.Println(colors)

	colors = append(colors[:len(colors)-1])
	fmt.Println(colors)

	numbers := make([]int, 4, 4)
	numbers[0] = 22
	numbers[1] = 22
	numbers[2] = 22
	numbers[2] = 22
	fmt.Println(numbers)

	numbers = append(numbers, 33)
	fmt.Println(numbers)
	sort.Ints(numbers)
	fmt.Println(numbers)

}
