package main

import "../dexes/console"
import "../dexes/rand"

func main () {
	console.Write("Рандомное трехзначное число = ", rand.RandomInt(10, 99) * 10)
}
