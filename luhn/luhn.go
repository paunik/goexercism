package luhn

// Valid checks if input valid Luhn checksum
func Valid(input string) bool {
	dCnt := 0
	sum := 0

	for i := len(input) - 1; 0 <= i; i-- {
		d := input[i]
		switch {
		case d == ' ':
			continue
		case d >= '0' && d <= '9':
			di := int(d - '0')

			if dCnt%2 == 1 {
				di *= 2

				if di > 9 {
					di -= 9
				}
			}
			sum += di
		default:
			return false
		}

		dCnt++
	}

	return dCnt > 1 && sum%10 == 0
}
