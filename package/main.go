package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/hassanjawwad12/go-practice/package/auth"
	"github.com/hassanjawwad12/go-practice/package/user"
)

func main() {
	auth.Login("hassan", "secret")

	session := auth.Session()

	fmt.Println(session)
	user := user.User{
		Email: "user@gmail.com",
		Name:  "john doe",
	}
	color.Red(user.Email)
	color.Magenta(user.Name)
}
