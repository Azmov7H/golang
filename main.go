package main

import (
	"fmt"
	"strconv"
)

func main() {
	var userName string = "Ali Naghi"
	var age int = 30
	fmt.Println("Hello, " + userName + "!")
	fmt.Println("Your age is: " + strconv.Itoa(age))
	calculateArea()
	stutement()
	}

func calculateArea()  {
	a := 10
	b := 5
	area := a * b
	plus := a + b
	munus := a - b
	division := a / b
	// Convert integers to strings for printing
	fmt.Println("The sum of a and b is: " + strconv.Itoa(plus))
	fmt.Println("The area of the rectangle is: " + strconv.Itoa(area))
	fmt.Println("The difference of a and b is: " + strconv.Itoa(munus))
	fmt.Println("The division of a and b is: " + strconv.Itoa(division))
	//Boolean
	isGreater := a > b
	notGreater := a < b
	fmt.Println("Is a less than b? " + strconv.FormatBool(notGreater))
	fmt.Println("Is a greater than b? " + strconv.FormatBool(isGreater))

}


func stutement()  {
	// if statement
	a := 10
	b := 5
	if a > b {
		fmt.Println("a is greater than b")
	} else if a < b {
		fmt.Println("a is less than b")
	} else {
		fmt.Println("a is equal to b")
	}
}
