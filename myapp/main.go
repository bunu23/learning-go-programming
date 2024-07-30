package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	jsonString := `{"name": "John Doe", "age": 30, "email": "john.doe@example.com"}`

	var person Person

	err := json.Unmarshal([]byte(jsonString), &person)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Parsed JSON: %+v\n", person)
}
