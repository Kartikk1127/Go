package main

import (
	"fmt"

	"github.com/Kartikk1127/Go/auth"
	"github.com/Kartikk1127/Go/user"
	"github.com/fatih/color"
)

func main() {
	auth.LoginWithCredentials("kartik", "secret")
	session := auth.GetSession()
	fmt.Println("Session: ", session)

	user := user.User{
		Email: "user@gmail.com",
		Name:  "Kartik",
	}

	// fmt.Println("User details:", user.Email, user.Name)
	color.Red(user.Email)
}
