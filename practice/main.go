package main

import (
	"fmt"
)

const aConst int = 64

func main() {
	var aString string = "This is Go!"
	fmt.Println(aString)
	fmt.Printf("The varibles type is %T\n", aString)

	var anInteger int = 42
	fmt.Println(anInteger)

	var defaultInt int
	fmt.Println(defaultInt)

	var anotherString = "This is another string"
	fmt.Println(anotherString)
	fmt.Printf("The varibles type is %T\n", anotherString)

	var anotherInt = 44
	fmt.Println(anotherInt)
	fmt.Printf("The varibles type is %T\n", anotherInt)

	myString := "this is also a string"
	fmt.Println(myString)
	fmt.Printf("The varibles type is %T\n", myString)

	fmt.Println(aConst)
	fmt.Printf("The varibles type is %T\n", aConst)

}
