package main

import "fmt"

func main() {
	var nums [5]int

	//array length
	fmt.Println(len(nums))

	nums[0] = 1
	fmt.Println(nums[0])

	fmt.Println(nums)

	temp := [3]int{1, 2, 3} //shorthand syntax to declare and initialize an array at the same time.
	fmt.Println(temp)

	//2d arrays
	data := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(data)
}
