package trailling_zeroes

func CountTrailingZeroesFact(n int) int {
	twos := 0
	fives := 0

	for i := n; i > 0; i-- {
		number := i
		fives += distribution(number, 5)
		twos += distribution(number, 2)
	}

	if twos > fives {
		return fives
	} else {
		return twos
	}

	return twos
}

func distribution(number int, factor int) int {
	count := 0
	for number%factor == 0 {
		count++
		number /= factor
	}
	return count
}
