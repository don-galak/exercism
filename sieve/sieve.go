package sieve

func Sieve(n int) []int {
	marks := make([]bool, n+1)

	primes := make([]int, n/2)
	k := 0
	for i := 2; i <= n; i++ {
		if marks[i] {
			continue
		}
		primes[k] = i
		k++

		for j := i; j <= n; j += i {
			marks[j] = true
		}
	}
	return primes[:k]
}
