package main

import "fmt"

func main() {
	// Arrays
	//var fruitArr [2]string

	// Assign values
	//fruitArr[0] = "Apple"
	//fruitArr[1] = "Orange"

	// Declare and assign
	//fruitArr := [2]string{"Apple", "Orange"}

	//fmt.Println(fruitArr)
	//fmt.Println(fruitArr[1])

	fruitArr := []string{"Apple", "Orange", "Grape"}

	// size of the array
	fmt.Println(len(fruitArr))
	// fruitArr[1:2] -> start in index 1 and stop in index 2 of array
	fmt.Println(fruitArr[1:2])
}
