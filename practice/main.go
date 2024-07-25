package main

const aConst int = 64

func main() {

	/*WORKING WITH DATES AND TIMES
	n := time.Now()
	fmt.Println("current date", n)

	t := time.Date(2001, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println("Hey its: ", t)
	fmt.Println(t.Format(time.ANSIC))

	parsedTime, _ := time.Parse(time.ANSIC, "Sun Nov 30 23:00:00 2000")
	fmt.Printf("type of parsed time is %T\n", parsedTime)

	*/

	/*
		USING THE MATH OPERATORSS
		i1, i2, i3 := 12, 44, 22
		intSum := i1 + i2 + i3
		fmt.Println("Integer sum: ", intSum)

		f1, f2, f3 := 12.32, 44.3, 22.3
		floatSum := f1 + f2 + f3
		fmt.Println("Integer sum: ", floatSum)

		floatSum = math.Round(floatSum*100) / 100
		fmt.Println("The sum is:", floatSum)

		circleRadius := 15.5
		circumference := circleRadius * 2 * math.Pi
		fmt.Printf("Circumference: %2f\n", circumference)

	*/

	/*
		GET INPUT FROM CONSOLE

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		input, _ := reader.ReadString('\n')
		fmt.Println("You entered:", input)

		fmt.Print("Enter a number: ")
		numInput, _ := reader.ReadString('\n')
		aFlot, err := strconv.ParseFloat(strings.TrimSpace(numInput), 64)
		if err != nil {
		 	fmt.Println(err)
		 } else {
		 	fmt.Println("Value of number: ", aFlot)
		 }

	*/
	// var aString string = "This is Go!"
	// fmt.Println(aString)
	// fmt.Printf("The varibles type is %T\n", aString)

	// var anInteger int = 42
	// fmt.Println(anInteger)

	// var defaultInt int
	// fmt.Println(defaultInt)

	// var anotherString = "This is another string"
	// fmt.Println(anotherString)
	// fmt.Printf("The varibles type is %T\n", anotherString)

	// var anotherInt = 44
	// fmt.Println(anotherInt)
	// fmt.Printf("The varibles type is %T\n", anotherInt)

	// myString := "this is also a string"
	// fmt.Println(myString)
	// fmt.Printf("The varibles type is %T\n", myString)

	// fmt.Println(aConst)
	// fmt.Printf("The varibles type is %T\n", aConst)

}
