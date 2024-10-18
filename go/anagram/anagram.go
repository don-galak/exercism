package anagram

import "strings"

// first 26 prime numbers, with n equal to the number of letters in latin alphabet
// will it work for other alphabets?
var primes = [26]uint{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101}

// * INFO: https://en.wikipedia.org/wiki/Fundamental_theorem_of_arithmetic

// TODO - Collisions are possible with unicode strings
func hashString(in string) uint {
	var hash uint = 1
	for _, v := range in {
		hash *= primes[v%26]
	}
	return hash
}

// Detect returns a list of anagrams
func Detect(in string, cands []string) []string {
	in = strings.ToLower(in)
	inHash := hashString(in)
	results := cands[:0] // Reuse allocated space
	for _, cand := range cands {
		if len(cand) == len(in) {
			candLower := strings.ToLower(cand)
			if candLower != in && hashString(candLower) == inHash {
				results = append(results, cand)
			}
		}
	}
	return results
}
