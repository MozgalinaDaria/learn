package main

import (
	"../base/console"
	"strings"
)

const CommandRegister = "register"
const CommandLogin = "login"
const CommandLogout = "logout"
const CommandExit = "exit"

type User struct {
	Password string
	Online   bool
}

var users = make(map[string]User)

func main() {
	var input []string
	var operation, login, password string

	for {
		input = strings.Split(console.ReadString("Введите операцию, логин и пароль через пробел: "), " ")
		operation = input[0]
		login = input[1]
		password = input[2]

		switch operation {
		case CommandRegister:
			console.Writeln(Register(login, password))
		case CommandLogin:
			console.Writeln(LogIn(login, password))
		case CommandLogout:
			console.Writeln(LogOut(login))
		case CommandExit:
			console.Writeln("You've made an exit")
			return
		}
	}
}

func Register(login, password string) string {
	if _, exists := users[login]; exists {
		return "fail: user already exists"
	}

	users[login] = User{Password: password, Online: false}
	return "success: new user added"
}

func LogIn(login, password string) string {
	if _, exists := users[login]; !exists {
		return "fail: no such user"
	}

	if users[login].Password != password {
		return "fail: incorrect password"
	}

	if users[login].Online {
		return "fail: already logged in"
	}

	user := users[login]
	user.Online = true
	users[login] = user
	return "success: user logged in"
}

func LogOut(login string) string {
	if _, exists := users[login]; !exists {
		return "fail: no such user"
	}

	if !users[login].Online {
		return "fail: already logged out"
	}

	user := users[login]
	user.Online = false
	users[login] = user
	return "success: user logged out"
}
