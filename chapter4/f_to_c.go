package main

import "fmt"

func main() {
	var f float64
	var c float64

	fmt.Print("Value in Fahrenheit: ")
	fmt.Scanf("%f", &f)

	c = (f - 32) * 5.0 / 9.0

	fmt.Println("Value in Celsius:", c)
}	
