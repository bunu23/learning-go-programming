package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	content := "hello from go"
	file, err := os.Create("./fromString.txt")
	checkError(err)
	defer file.Close()

	length, err := io.WriteString(file, content)
	checkError(err)
	fmt.Printf("wrote a file with %v characters \n", length)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
