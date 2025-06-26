package main

import (
	"fmt"
)

// range keyword is  used for iterating over data structures.
func main() {
	nums := []int{6, 7, 8}

	// normal iteration without range keyword
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}
	fmt.Println("end")

	var sum = 0
	// iteration with range keyword where "it" denotes the index and "num" denotes the value at the index
	for it, num := range nums {
		fmt.Println(it)
		sum += num
	}
	fmt.Println(sum)
	fmt.Println("end")

	//iteration on a map using range
	m := map[string]string{"fname": "kartik", "lname": "srivastava"}
	for key, value := range m {
		fmt.Println(key, value)
	}
	fmt.Println("end")

	//range can also be used on a string where "i" denotes the index and "c" denotes the unicode
	for i, c := range "golang" {
		fmt.Println(i, string(c))
	}
}
