package main

import "fmt"

func main() {
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"

	fmt.Println(elements["Li"])
	name, ok := elements["Wow"]
	fmt.Println(name, ok)
}
