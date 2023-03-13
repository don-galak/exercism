package palindrome

import (
	"errors"
	"fmt"
)

type Product struct {
	value          int
	Factorizations [][2]int
}

func (p Product) pairExists(a, b int) bool {
	for _, f := range p.Factorizations {
		if a == f[0] || a == f[1] || b == f[0] || b == f[1] {
			return true
		}
	}
	return false
}

func (p *Product) removeFalsyProducts() {
	for i, f := range p.Factorizations {
		if p.value > f[0]*f[1] {
			p.Factorizations = append(p.Factorizations[:i], p.Factorizations[i+1:]...)
		}
	}
}

var errNoPalidromes = errors.New("no palindromes")
var errFMinGTFmax = errors.New("fmin > fmax")

func isPalindrome(num int) bool {
	toString := fmt.Sprintf("%d", num)
	orig := toString

	n := len(toString)
	runes := make([]rune, n)
	for _, rune := range toString {
		n--
		runes[n] = rune
	}
	return string(runes[n:]) == orig
}

func Products(fmin, fmax int) (Product, Product, error) {
	pmin := Product{}
	pmax := Product{}
	if fmin > fmax {
		return pmin, pmax, errFMinGTFmax
	}

	for i := fmin; i < fmax+1; i++ {
		for j := fmin; j < fmax+1; j++ {
			if isPalindrome(i*j) && pmin.value == 0 {
				pmin.value = i * j
				pmin.Factorizations = append(pmin.Factorizations, [2]int{i, j})
				break
			}
		}
		if pmax.value != 0 {
			break
		}
	}

	for i := fmax; i > fmin-1; i-- {
		for j := fmax; j > fmin-1; j-- {
			if isPalindrome(i*j) && i*j >= pmax.value && !pmax.pairExists(i, j) {
				pmax.value = i * j
				pmax.Factorizations = append(pmax.Factorizations, [2]int{j, i})
			}
		}
	}
	pmax.removeFalsyProducts()

	if pmin.value == 0 && pmax.value == 0 {
		return pmin, pmax, errNoPalidromes
	}

	return pmin, pmax, nil
}
