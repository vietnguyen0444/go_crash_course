package main

import "fmt"

func main() {
	/// MAIN TYPES
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// unit unit8 unit16 unit32 unit64
	// byte - alias for unit8
	// rune - alias for unit32
	// float32 float64
	// complex64 complex128
	///

	// Using var
	// Note: when crate variables, we need to use it
	//var name string = "Ryan"
	var age int32 = 22
	var isCool = true
	var size float32 = 1.4

	// Shorthand
	//name := "Ryan"
	//email := "ryan@gmail.com"

	// Create variables in 1 line
	name, email := "Ryan", "ryan@gmail.com"

	// REFFERENCE -> https://pkg.go.dev/fmt
	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T\n", size)
}
