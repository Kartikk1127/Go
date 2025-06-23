package main

import (
	"fmt"
	"slices"
)

// slices are dynamic  arrays
// used when the number of elements are not known
// most used construct in go
// useful methods
func main() {

	// declaring a slice
	// uninitialized slice is nil(null in Java)
	var temp []int
	fmt.Println(temp == nil)

	fmt.Println(len(temp))

	var nums = make([]int, 2) // this method "make" is provided by go that is usually used to declare and initialize slices.
	fmt.Println(nums)
	fmt.Println(nums == nil) // now this is not nil

	// maximum number of elements can fit
	fmt.Println(cap(nums)) // "cap" is the inbuilt method provided by go.

	var data = make([]int, 2, 5) // "make" takes either 2 or 3 arguments. 2nd argument is the length, 3rd is the capacity.
	fmt.Println(cap(data))

	// adding elements in slice
	data = append(data, 1)
	data = append(data, 2)
	data = append(data, 3)
	data = append(data, 4)
	data = append(data, 5)
	data = append(data, 6)
	data = append(data, 7)
	data = append(data, 8)
	data = append(data, 9)
	fmt.Println(data)
	fmt.Println(cap(data)) // capacity is doubled whenever the number of elements inside a slice exceed the current capacity.

	// in the above examples, you can see that the slices created have 0 as their first few elements. This is because when "make" method was used and was provided the length, the number of 0 was inserted till the length was reached internally, hence whenever using "make" method, we keep the length as 0 while the capacity can be anything that we want.
	var dataList = make([]int, 0, 5)
	dataList = append(dataList, 2)
	fmt.Println(dataList) // now you won't see any 0 elements since we kept the length to 0, but made sure to mention the capacity.
	fmt.Println(len(dataList))

	// another way to create a slice
	info := []int{}
	info = append(info, 1)
	info = append(info, 2)
	fmt.Println(info)
	fmt.Println(cap(info))
	fmt.Println(len(info))

	// copying a slice
	var slc = make([]int, 0, 5)

	slc = append(slc, 2)
	var slc2 = make([]int, len(slc))

	fmt.Println(slc)
	fmt.Println(slc2)
	// copy function
	copy(slc2, slc)
	fmt.Println(slc)
	fmt.Println(slc2)

	// slice operator
	var slcOperator = []int{1, 2, 3}
	fmt.Println(slcOperator[0:2]) // takes a range, index1 is inclusive, index2 is exclusive
	fmt.Println(slcOperator[:2])  // if first index not given, it will consider everything from the start(inclusive) to the 2nd index(exclusive)
	fmt.Println(slcOperator[1:])  // if second index not given, it will consider everything from the first index(inclusive) to the end(inclusive)
	fmt.Println(slcOperator[1:2]) // you can also give a mid range, index1 is inclusive, index2 is exclusive

	// slice package
	var nums1 = []int{1, 2}
	var nums2 = []int{1, 2}

	fmt.Println(slices.Equal(nums1, nums2)) // returns if equal or not

	// 2d slices
	var nums2D = [][]int{{1, 2}, {3, 4}}
	fmt.Println(nums2D)
}
