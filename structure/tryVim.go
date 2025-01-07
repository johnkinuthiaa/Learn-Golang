package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("hii...im using neo vim")
	fmt.Println(time.Now())
	user := login{
		username: "john",
		password: "Lightyagami@123",
	}
	fmt.Println(user.loginUser())
	fmt.Println(user.registerUser())
}

type login struct {
	username string
	password string
}

func (login login) loginUser() (result string) {
	return fmt.Sprintf("login successfull \nusername: %s \npassword: %s ", login.username, login.password)
}
func (register login) registerUser() (result string) {
	return fmt.Sprintf("register user : %s", register.password)

}
