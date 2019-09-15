package main

import "../base/console"

func main() {
	var s string
	var length int

	console.Writeln("Напишите предложение")
	s = console.ReadString()
	length = len(s)

	console.Writeln("Половина строки: ", s[0:length/2])
	console.Writeln("Вторая половина строки: ", s[length/2:length])
}
