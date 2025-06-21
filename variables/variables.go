package main

import "fmt"

func main() {
	//declaring a string variable

	var name string = "golang"
	fmt.Println(name)

	/* OR
	var name = "golang" //here the type is automatically inferred
	fmt.Println(name) */

	//declaring a boolean

	var isAdult = true
	fmt.Println(isAdult)

	//declaring a number

	var age int = 30
	fmt.Println(age)

	//shorthand syntax used when you want to declare and assign a variable at the same time.
	data := "golang"
	fmt.Println(data)
}
