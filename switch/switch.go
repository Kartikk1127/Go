package main

import (
	"fmt"
	"time"
)

func main() {

	//simple switch -- no need to write the break keyword unlike other languages. Also, default case is optional.
	i := 3

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("other")
	}

	// multiple condition switch

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println("work day")
	}

	//type switch
	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("Its an integer")
		case string:
			fmt.Println("this is a string")
		case bool:
			fmt.Println("this is a boolean")
		default:
			fmt.Println("this is any type", t)
		}
	}

	whoAmI(3.5)
}
