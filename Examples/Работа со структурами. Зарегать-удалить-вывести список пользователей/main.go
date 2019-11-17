package main

import (
	"../../base/console"
	"strings"
)

const CommandRegister = "register"
const CommandDelete = "delete"
const CommandList = "list"

type User struct {
	Name     string
	Surname  string
	Password string
}

var users = make(map[string]User)

func main() {
	var operation, login, password string

	for {
		operation = console.ReadString("Введите операцию: ")

		switch operation {
		case CommandRegister:
			input := strings.Split(console.ReadString("Введите логин, пароль: "), " ")
			login = input[0]
			password = input[1]
			console.Writeln(Register(login, password))
		case CommandDelete:
			login = console.ReadString("Введите логин, данные по которому нужно удалить: ")
			console.Writeln(Delete(login))
		case CommandList:
			List(users)
		}
	}
}

func Register(login, password string) string {
	if _, exists := users[login]; exists {
		return "fail: user already exists"
	}
	input := (strings.Split(console.ReadString("Введите фамилию и имя пользователя: "), " "))
	users[login] = User{Surname: input[0], Name: input[1], Password: password}

	return "success: new user added"
}

func Delete(login string) string {
	if _, exists := users[login]; !exists {
		return "can't delete it, because it doesn't exists"
	}
	delete(users, login)
	return "the user has been deleted"
}

func List(users map[string]User) {
	console.Writeln(users)
	return
}
