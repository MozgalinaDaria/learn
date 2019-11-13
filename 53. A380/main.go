package main

import (
	"../base/console"
	"strconv"
)

func main() {
	row, seat := ParseInput(console.ReadString("Введите номер посадочного места (1А): "))
	if row <= 2 {
		console.Writeln(Premium(seat))
		return
	}

	if row > 2 && row <= 20 {
		console.Writeln(Business(seat))
		return
	}

	console.Writeln(Econom(seat))
	return
}

func ParseInput(input string) (int, string) {
	length := len(input)
	num, _ := strconv.Atoi(input[:length-1])
	seat := input[length-1:]
	return num, seat
}

func Premium(seat string) string {
	if seat == "A" || seat == "D" {
		return "window"
	}
	return "aisle"
}

func Business(seat string) string {
	if seat == "A" || seat == "F" {
		return "window"
	}
	return "aisle"
}

func Econom(seat string) string {
	if seat == "A" || seat == "K" {
		return "window"
	}

	if seat == "B" || seat == "E" || seat == "F" || seat == "J" {
		return "neither"
	}
	return "aisle"
}
