package main

import "fmt"

const age = 30

func main() {
	const name = "Kartikey" //cannot be reassigned now

	// name = "hello" will throw an error now since it is already assigned.

	fmt.Println(name, age)

	//constants grouping
	const (
		port = 5000
		host = "localhost"
	)

	fmt.Println(port, host)
}
