package main

import "fmt"

// Method - 1
// func printSlice(items []int) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// Now the generic example
// Method - 2
// func printSlice[T any](items []T) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// If you want to pass values of specific data types then
// func printSlice[T int | string](items []T) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// or you can also specify the type as comparable, means the same.
type stack[T comparable] struct {
	elements []string
}

func main() {

	myStack := stack[string]{
		elements: []string{"john", "doe"},
	}
	fmt.Println(myStack)

	// nums := []int{1, 2, 3}

	// names := []string{"golang", "typescipt"} // since printSlice method expects an integer slice hence can't pass this.
	// printSlice(nums)
}
