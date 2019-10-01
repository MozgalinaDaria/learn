package main

import "../base/console"
import "./task"

func main() {
	var a int
	var nominative, genitiveSingular, genitivePlural string

	a = console.ReadInt("Введите число: ")
	nominative = console.ReadString("Введите сущетвительное в и.п.: ")
	genitiveSingular = console.ReadString("Введите сущетвительное в р.п. ед.ч.: ")
	genitivePlural = console.ReadString("Введите сущетвительное в р.п. мн.ч.: ")

	console.Writeln(a, task.Declension(a, nominative, genitiveSingular, genitivePlural))
}
