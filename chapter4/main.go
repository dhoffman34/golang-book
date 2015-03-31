package main

import "fmt"

func main() {
	var x string = "Hello World"
	fmt.Println(x)
	x = "Hi there"
	fmt.Println(x)

	y := "Hmm..."
	fmt.Println(y)

	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2

	fmt.Println(output)
}	
