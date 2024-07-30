package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var result string
	switch dow := rand.Intn(7) + 1; dow {
	case 1:
		result = "hey its monday"
	case 2:
		result = "Tuesday"
	default:
		result = "its Weekend"
	}

	fmt.Println(result)
}
