package summultiples

func SumMultiples(limit int, divisors ...int) (sum int) {
	for i := 1; i < limit; i++ {
		for _, divisor := range divisors {
			if divisor == 0 {
				break
			}
			if i%divisor == 0 {
				sum += i
				break
			}
		}
	}
	return
}
