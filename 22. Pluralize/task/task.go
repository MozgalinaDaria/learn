package task

func Declension(a int, nominative, genitiveSingular, genitivePlural string) string {
	var reminderTen = a % 10
	var reminderHundred = a % 100

	if reminderTen == 1 && reminderHundred != 11 {
		return nominative
	}

	if reminderTen >= 2 && reminderTen <= 4  && reminderHundred < 12 || reminderTen >= 2 && reminderTen <= 4  && reminderHundred > 19 {
		return genitiveSingular
	}

	return genitivePlural
}
