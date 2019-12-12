package main

import "../base/console"

func main() {
	var sum int

	for pens := 0; pens < 30; pens++ {
		for pencils := 0; pencils < 30; pencils++ {
			for rubber := 30 - pencils - pens; rubber >= 0; rubber-- {
				sum = (pens * 10) + (pencils * 5) + (rubber * 2)
				if sum == 100 {
					console.Writeln("Можно приобрести ручек, карандашей и ластиков в количестве ", pens, ", ",
						pencils, ", ", rubber, "шт. соответственно")
				}
			}
		}

	}
}
