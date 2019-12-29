package leap

func divByNum(target int, num int) bool {
	return target % num == 0
}

// IsLeapYear should have a comment documenting it.
func IsLeapYear(year int) bool {
	return divByNum(year, 4) && (divByNum(year, 100) == false || divByNum(year, 400))
}
