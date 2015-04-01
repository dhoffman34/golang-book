package main

import "fmt"

func main() {
	for i := 1; i <= 100; i += 1 {
		fmt.Println(i)
		switch 0 {
		case i % 15: fmt.Println("FizzBuzz")
		case i % 3: fmt.Println("Fizz")
		case i % 5: fmt.Println("Buzz")
		}
	}
}
