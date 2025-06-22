package main

import "fmt"

func main() {

	// age := 20

	// if age >= 18 {
	// 	fmt.Println("Person is an adult")
	// } else if age >= 12 {
	// 	fmt.Println("Person is a teenager")
	// } else {
	// 	fmt.Println("Person is a kid")
	// }

	// logical operator example
	var role = "admin"
	var hasPermissions = false

	if role == "admin" && hasPermissions {
		fmt.Println("this man got access to everything")
	}

	//we can declare a variable inside if block as well
	if age := 15; age > 18 {
		fmt.Println("person is an adult")
	} else {
		fmt.Println("Person is not an adult")
	}

	// go does not have ternary operators.
}
