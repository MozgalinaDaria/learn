package main

import (
	"../base/console"
)

func main() {
	var key string
	var value int
	var numberOfLogins int

	logins := make(map[string]int)

	numberOfLogins = console.ReadInt("Введите общее количество логинов: ")

	for i := 0; i < numberOfLogins; i++ {
		logins[console.ReadString("Введите логин: ")]++
	}

	console.Writeln("Логины спамеров: ")

	for key, value = range logins {
		if value > 1 {
			console.Write(key, " ")
		}
	}
}
