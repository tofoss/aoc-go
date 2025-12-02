package math

import "sort"

func FactorsOf(n int) []int {
	factors := []int{}

	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			factors = append(factors, i)
			if i != n/i {
				factors = append(factors, n/i)
			}
		}
	}

	sort.Ints(factors)
	return factors
}
