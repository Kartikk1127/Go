package auth

import "fmt"

// as of now this method can be used within this package only
// methods starting with small case are scoped within their own package.
// func loginWithCredentials(username string, password string) {
// 	fmt.Println("Logging in with ", username)
// }

// now this can be scoped outside as well since it starts with capital letter
func LoginWithCredentials(username string, password string) {
	fmt.Println("Logging in with ", username)
}
