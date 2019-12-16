package main

import "../base/console"

func main() {
	number := 0

	for line := 0; line <= 50; line++ {
		for j := 0; j < 5; j++ {
			console.Write("|", number, " ", string(number), "|")
			number++
		}
		console.Writeln()
	}
	console.Writeln(number, " ", string(number))
}
