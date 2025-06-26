package main

import "fmt"

// creating a variadic function
func sum(nums ...int) int {
	result := 0
	for _, num := range nums {
		result += num
	}
	return result
}

// variadic functions are those functions where you can pass any number of parameters. An example for this can be "println"
func main() {
	fmt.Println(1, 2, 3, 4)

	temp := []int{1, 2, 3, 4, 5}

	fmt.Println(sum(1, 2, 3, 43, 4, 5, 5645))
	fmt.Println(sum(temp...)) //slice's value is now spread
}
