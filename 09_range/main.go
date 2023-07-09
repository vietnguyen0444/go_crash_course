package main

import "fmt"

func main() {
	ids := []int{33, 67, 23, 22, 5}

	// Loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}

	fmt.Println("Sum: ", sum)

	// Range with map
	emails := map[string]string{"John": "john@gmail.com", "Ryan": "ryan@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range emails {
		fmt.Println("Name: ", k)
	}
}
