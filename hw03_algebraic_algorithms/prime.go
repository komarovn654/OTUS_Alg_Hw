package hw03algebraicalgorithms

import (
	"errors"
	"fmt"
	"math"
)

var (
	ErrPrimeGreaterOne = errors.New("Prime numbers greater than one")
)

func BruteforcePrime(max int) (int, error) {
	if max <= 1 {
		return 0, ErrPrimeGreaterOne
	}
	primeCount := 1
	for i := 3; i < max; i++ {
		if isPrimeBF(i) {
			primeCount++
		}
	}
	return primeCount, nil

}

func isPrimeBF(v int) bool {
	for i := 2; i < v; i++ {
		if v%i == 0 {
			return false
		}
	}
	return true
}

func PrimeOpt(max int) (int, error) {
	if max <= 1 {
		return 0, ErrPrimeGreaterOne
	}
	primes := []int{2}
	primeCount := 1

	for i := 3; i < max; i++ {
		if isPrimeOpt(i, primes) {
			primes = append(primes, i)
			primeCount++
		}
	}
	return primeCount, nil

}

func isPrimeOpt(v int, primes []int) bool {
	sqrt := int(math.Sqrt(float64(v)))
	for i := 0; primes[i] <= sqrt; i++ {
		if v%primes[i] == 0 {
			return false
		}
	}
	return true
}

func PrimeEr(max int) int {
	primesCount := 0
	isntPrime := make([]bool, max+1)
	ptr := 2
	for ptr < max {
		for i := ptr + ptr; i < len(isntPrime); i += ptr {
			isntPrime[i] = true
		}
		for ptr++; isntPrime[ptr] && ptr < max; ptr++ {
		}
		primesCount++
	}
	return primesCount
}

// 0	 1	   2	 3	   4	 5	   6	 7	   8	 9	   10
// false false false false false false false false false false false
// false false false false true  false true  false true  false true
// false false false false true  false true  false true  true  true
// false false false false true  false true  false true  true  true
// false false false false true  false true  false true  true  true

func PrimeErOpt(max int) int {
	primesCount := 0
	isntPrime := make([]int32, max/32+1)
	ptr := 2
	for ptr < max {
		for i := ptr + ptr; i < max; i += ptr {
			isntPrime[i/32] |= 1 << i
		}
		for ptr++; isntPrime[ptr/32]&(1<<ptr) != 0 && ptr < max; ptr++ {
		}
		primesCount++
	}
	fmt.Println(primesCount)
	return primesCount
}
