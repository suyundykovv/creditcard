package main

func lunacheckS(a string) bool {
	sum := 0
	ndigits := len(a)
	parity := ndigits % 2

	for i := 0; i < ndigits; i++ {
		digit := int(a[i] - '0')
		if i%2 == parity {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
	}
	return sum%10 == 0
}
