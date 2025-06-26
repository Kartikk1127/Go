package main

import "fmt"

// the variable passed here is by value i.e. it's just a copy
func changeNum(num int) {
	num = 5
	fmt.Println("In changeNum", num)
}

// by reference[* here is the pointer]
func changeNumByReference(num *int) {
	*num = 5 //dereferencing
	fmt.Println("In changeNumByReference", *num)
}

func main() {
	num := 1
	// changeNum(num)
	// fmt.Println("After changeNum in main", num) // the value of num in main won't change since we are passing it to changeNum by value

	//what happens if we send the reference?
	fmt.Println("Memory address", &num) //& helps in fetching the memory address
	changeNumByReference(&num)

	fmt.Println("After changeNumByReference", num)
}
