package main

import (
	"../base/console"
	"../base/rand"
)

const TheSame = 0
const LessThanRandom = 1
const MoreThanRandom = 2

func main() {
	randomNumber := rand.RandomInt(1, 1000)
	usersNumber := 0

	console.Writeln("Компьютер загадал число, попробуйте угадать его")

	for randomNumber != usersNumber{
		usersNumber = console.ReadInt("Какое число загадал компьютер? ")

		switch compare(randomNumber, usersNumber) {
		case TheSame:
			console.Writeln("Победа! Вы угадали число")
			return
		case LessThanRandom:
			console.Writeln("Ваше число меньше загаданного. Попробуйте угадать снова")
			continue
		case MoreThanRandom:
			console.Writeln("Ваше число больше загаданного. Попробуйте угадать снова")
			continue
		}
	}
}

func compare(num_1, num_2 int) int {
	if num_1 > num_2 {
		return 1
	}
	if num_1 < num_2 {
		return 2
	}
	return 0
}
