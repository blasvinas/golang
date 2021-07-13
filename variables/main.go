package main

import "fmt"

func main() {
	// Strings
	var s string
	s = "I like Go!"
	fmt.Println(s)

	//Runes
	var c rune
	c = 'A'
	fmt.Println(c) // Print 65 which is the Unicode code

	// Booleans
	isOk := true
	fmt.Println(isOk)

	// Numbers
	var x int = 5
	var y float64 = 3.15
	fmt.Println(x, y)
}
