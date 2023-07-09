package main

import "fmt"

func main() {
	// Define Map
	//emails := make(map[string]string)

	// Assign key value
	//emails["Ryan"] = "ryan@email.com"
	//emails["John"] = "john@email.com"

	// Declare map and add key value
	emails := map[string]string{"John": "john@gmail.com", "Ryan": "ryan@gmail.com"}

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Ryan"])

	// Delete from Map
	delete(emails, "John")
}
