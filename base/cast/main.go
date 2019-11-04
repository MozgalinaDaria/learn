package cast

func CastRuneToInt(r []rune) int{
	var number, value int
	length := len(r) - 1
	countForTen := 1

	for i := length; i >= 0; i-- {
		number = int (r [i] - '0')
		value += number * countForTen
		countForTen *= 10
	}

	return value
}
