package sieve

func Sieve(n int) []int {
	ns := make([]bool, n+1)
	primes := []int{}
	for i := 2; i <= n; i++ {
		marked := ns[i]
		if !marked {
			primes = append(primes, i)
			for j := 1; ; j++ {
				k := i * j
				if k > n {
					break
				}
				ns[k] = true
			}
		}
	}
	return primes
}
