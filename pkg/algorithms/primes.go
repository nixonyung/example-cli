package algorithms

import "math"

func Primes(
	ub int, // upper bound, inclusive
) []int {
	// (ref.) [Sieve of Eratosthenes](https://www.wikiwand.com/en/Sieve_of_Eratosthenes)

	// init
	isPrimes := make([]bool, ub+1)
	isPrimes[0] = false
	isPrimes[1] = false
	for i := 2; i <= ub; i++ {
		isPrimes[i] = true
	}

	// algo
	ubSqrt := int(math.Sqrt(float64(ub)))
	for i := 2; i <= ubSqrt; i++ {
		if isPrimes[i] {
			for j := i * i; j <= ub; j += i {
				isPrimes[j] = false
			}
		}
	}

	result := []int{}
	for i, isPrime := range isPrimes {
		if isPrime {
			result = append(result, i)
		}
	}
	return result
}
