package prime

func Factors(n int64) (out []int64) {
	var i int64 = 2
	for n > 1 {
		if n%i == 0 {
			n /= i
			out = append(out, i)
		} else {
			i++
		}
	}

	return
}
