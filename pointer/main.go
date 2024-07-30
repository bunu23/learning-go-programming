package main

import (
	"fmt"
)

func main() {
	anInt := 42
	var p = &anInt
	fmt.Println("value od p:", *p)

	value := 42.12
	pointer1 := &value
	fmt.Println("value:", *pointer1)

	*pointer1 = *pointer1 / 2
	fmt.Println("pointer1: ", *pointer1)
	fmt.Println("value:", value)
}
