package main

import (
	"../base/console"
	"strings"
)

const CommandRegister = "register"
const CommandLogin = "login"
const CommandLogout = "logout"
const CommandExit = "exit"

var logins = make(map[string]string)
var IsOnline = make(map[string]bool)

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
	if _, exists := logins[login]; exists {
		return "fail: user already exists"
	}

	logins[login] = password
	return "success: new user added"
}

func LogIn(login, password string) string {
	if _, exists := logins[login]; !exists {
		return "fail: no such user"
	}

	if logins[login] != password {
		return "fail: incorrect password"
	}

	if IsOnline[login] {
		return "fail: already logged in"
	}

	IsOnline[login] = true
	return "success: user logged in"
}

func LogOut(login string) string {
	if _, exists := logins[login]; !exists {
		return "fail: no such user"
	}

	if IsOnline[login] == false {
		return "fail: already logged out"
	}
	IsOnline[login] = false
	return "success: user logged out"
}
