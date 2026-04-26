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
}
