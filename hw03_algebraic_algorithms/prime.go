package hw03alg

import (
	"math"
)

// Calculate primes count in max range. Brutforce method. O(N^2).
func PrimeBruteforce(max int64) int64 {
	if max <= 1 {
		return 0
	}

	primeCount := int64(1)
	for i := int64(3); i <= max; i++ {
		if isPrimeBF(i) {
			primeCount++
		}
	}

	return primeCount
}

func isPrimeBF(v int64) bool {
	for i := int64(2); i < v; i++ {
		if v%i == 0 {
			return false
		}
	}

	return true
}

// Calculate primes count in max range. Optimized brutforce method. O(N * Sqrt(N) / Ln (N)).
func PrimeBFOpt(max int64) int64 {
	if max <= 1 {
		return 0
	}
	primes := []int64{2}
	primeCount := int64(1)

	for i := int64(3); i <= max; i++ {
		if isPrimeOpt(i, primes) {
			primes = append(primes, i)
			primeCount++
		}
	}

	return primeCount
}

func isPrimeOpt(v int64, primes []int64) bool {
	sqrt := int64(math.Sqrt(float64(v)))
	for i := int64(0); primes[i] <= sqrt; i++ {
		if v%primes[i] == 0 {
			return false
		}
	}

	return true
}

// Calculate primes count in max range. Sieve of Eratosthenes method. O(N Log Log N).
func PrimeErat(max int64) int64 {
	if max <= 1 {
		return 0
	}
	if max == 2 {
		return 1
	}

	primesCount := int64(0)
	isntPrime := make([]bool, max+2)
	ptr := int64(2)
	for ptr <= max {
		for i := int64(ptr + ptr); i <= max; i += ptr {
			isntPrime[i] = true
		}
		for ptr++; isntPrime[ptr] && ptr < max; ptr++ {
		}
		if !isntPrime[ptr] {
			primesCount++
		}
	}

	return primesCount
}

// Calculate primes count in max range. Memory optimized Sieve of Eratosthenes method. O(N Log Log N).
func PrimeEratMemOpt(max int64) int64 {
	if max <= 1 {
		return 0
	}
	if max == 2 {
		return 1
	}

	primesCount := int64(0)
	isntPrime := make([]int32, max/32+2)
	ptr := int64(2)
	for ptr <= max {
		for i := ptr + ptr; i <= max; i += ptr {
			isntPrime[i/32] |= calcOffset(i)
		}
		for ptr++; isntPrime[ptr/32]&calcOffset(ptr) != 0 && ptr < max; ptr++ {
		}
		if isntPrime[ptr/32]&calcOffset(ptr) == 0 {
			primesCount++
		}
	}

	return primesCount
}

func calcOffset(i int64) int32 {
	return 1 << (i - 32*(i/32))
}

// Calculate primes count in max range. Optimized Sieve of Eratosthenes method. O(N).
func PrimeEratOpt(max int64) int64 {
	if max <= 1 {
		return 0
	}
	if max == 2 {
		return 1
	}

	li := max // last index
	lp := make([]int64, max+1)
	primes := make([]int64, max+1)
	pPtr := li
	for i := int64(2); i <= max; i++ {
		if lp[i] == 0 {
			lp[i] = i
			primes[pPtr] = i
			pPtr--
		}
		for p := li; primes[p] <= lp[i] && primes[p]*i <= li && primes[p] != 0; p-- {
			lp[i*primes[p]] = primes[p]
		}
	}

	return li - pPtr
}
