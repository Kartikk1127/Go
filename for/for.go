package main

import "fmt"

// for -> only construct in go for looping
func main() {

	// implementing while loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	// implementing while loop infinite times
	// for {
	// 	println(1)
	// }

	// classic for loop
	for j := 0; j < 3; j++ {
		// break -- breaks out of for loop
		// continue -- skips the current iteration
		if j == 2 {
			continue
		}
		fmt.Println(j)
	}

	//range -- below means 0(inclusive) and 3(exclusive)
	for i := range 3 {
		fmt.Println(i)
	}
}
