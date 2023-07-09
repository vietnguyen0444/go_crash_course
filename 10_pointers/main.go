package main

import "fmt"

func main() {
	a := 5
	b := &a // link to address of variable a

	fmt.Println(a, b)     // Output: 5 0xc00001c098
	fmt.Printf("%T\n", b) // Output: *int

	// Use * top read val from address of variable
	fmt.Println(*b)  // Output: 5
	fmt.Println(*&a) // Output: 5

	// Change val with pointer
	*b = 10
	fmt.Println(a) // Output: 10
}
