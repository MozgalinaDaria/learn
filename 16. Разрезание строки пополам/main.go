package main

import "../base/console"

func main() {

	s := console.ReadString("Напишите предложение")
	length := len(s)

	console.Writeln("Половина строки: ", s[0:length/2])
	console.Writeln("Вторая половина строки: ", s[length/2:length])
}
