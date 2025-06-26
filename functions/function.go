package main

import "fmt"

// can also be written as func add(a, b int) int {}
func add(a int, b int) int {
	return a + b
}

// can also return multiple values, just put brackets in the return type and mention the type that it would be returning.
func getLanguages() (string, int, bool) {
	return "golang", 1, false
}

// func processIt(fn func(a int) int) {
// 	fn(5)
// }

func processIt() func(a int) int {
	return func(a int) int {
		return 2
	}
}

func main() {
	result := add(2, 3)
	fmt.Println(result)

	// lang1, lang2, lang3 := getLanguages() -- if you don't want to use a value like lang3, replace it by _(underscore)
	lang1, lang2, _ := getLanguages()

	fmt.Println(lang1, lang2)

	//  fn := func(a int) int {
	// 	return 2
	// }
	// this is an anonymous function
	fn := processIt()
	fn(6)

}
