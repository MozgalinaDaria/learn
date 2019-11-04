package main

import "../base/console"

func main() {
	var n int

	n = console.ReadInt("Введите степень n: ")
	if n == 1 {
		console.Write("1 верблюд, 2 барана, 3 таракана")
		return
	}

	if n == 2 {
		console.Write("3 верблюда, 4 барана, 5 тараканов")
		return
	}

	console.Write("Решений нет")

}
