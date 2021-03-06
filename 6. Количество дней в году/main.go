package main

import "../base/console"

func main() {
	year := console.ReadInt("Введите год")

	if isLeap(year) {
		console.Writeln("Год ", year, " - високосный, в нем  365 дней")
		return
	}

	console.Writeln("Год ", year, " - невисокосный, в нем 366 дней")
}

func isLeap(year int) bool {
	var isDividedBy = make(map[int]bool)

	isDividedBy[4] = year%4 == 0
	isDividedBy[100] = year%100 == 0
	isDividedBy[400] = year%400 == 0

	return isDividedBy[400] || !isDividedBy[100] && isDividedBy[4]
}
