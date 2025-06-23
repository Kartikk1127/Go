package main

import (
	"fmt"
	"maps"
)

func main() {

	// creating maps
	m := make(map[string]string)

	// setting an element
	m["name"] = "golang"
	m["area"] = "backend"

	// get an element
	fmt.Println(m["name"])
	fmt.Println(m["area"])
	fmt.Println(m["data"]) // if a key doesn't exist in the map, it doesn't throw an error instead returns an empty value(NOT NIL). If the type was Integer, it would have returned 0 --  basically depends on the type of the value inside the map

	m2 := make(map[string]int)

	m2["age"] = 30
	fmt.Println(m2["age"])
	fmt.Println(m2["date"]) // now this will return 0 since the key doesn't exist but the data type for value was Integer.

	fmt.Println(len(m)) // to get the length of the map

	//delete an element from the map
	delete(m, "area")
	fmt.Println(m)

	//clear the map
	clear(m)
	fmt.Println(m)

	//second way to create a map
	temp := map[string]int{"price": 40, "phones": 3}
	fmt.Println(temp)

	// check whether an element exist or not
	v, ok := temp["phones"] // here "ok" is just an expression used a lot in golang, you can use any other literal as well. However, "ok" denotes a boolean.

	fmt.Println(v) // here "v" denotes the value of the key mentioned above
	if ok {
		fmt.Println("all ok")
	} else {
		fmt.Println("not ok")
	}

	// check whether 2 maps are equal or not
	temp2 := map[string]int{"price": 40, "phones": 8}
	fmt.Println(maps.Equal(temp, temp2)) // returns if two maps are equal or not
}
