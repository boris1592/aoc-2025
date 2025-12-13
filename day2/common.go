package day2

func intLen(n int) (l int) {
	for n > 0 {
		l++
		n /= 10
	}

	return
}
