package main

import "fmt"

func main() {
	x := 15
	y := 10

	// If Else statement
	if x <= y {
		fmt.Printf("%d is less than or equal to %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	// Else If statement
	color := "red"

	if color == "red" {
		fmt.Println("color is red")
	} else if color == "green" {
		fmt.Println("color is green")
	} else {
		fmt.Println("color is Not green or red")
	}

	// Switch
	switch color {
	case "red":
		fmt.Println("color is red")
	case "green":
		fmt.Println("color is green")
	default:
		fmt.Println("color is Not green or red")
	}
}
