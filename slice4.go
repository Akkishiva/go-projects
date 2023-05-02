package main

import "fmt"

func main() {

	// Create a slice of integers
	mySlice := []int{1, 2, 3, 4, 5}

	// Create an empty slice of integers
	emptySlice :=make( []int,6,20)

	// Append the empty slice to the existing slice
	mySlice = append(mySlice, emptySlice...)

	// Print the capacity of the slice
	fmt.Println(cap(mySlice))
}