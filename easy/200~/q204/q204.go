package q204

func countPrimes(n int) int {
	if n < 2 {
		return 0
	} else if n == 2 {
		return 1
	}
	primes := []int{2}
	for i := 3; i <= n; i = i + 2 {
		if isPrime(primes, i) {
			primes = append(primes, i)
		}
	}
	return len(primes)
}

func isPrime(primes []int, num int) bool {
	mid := num/2 + 1
	for _, p := range primes {
		if p > mid {
			return true
		}
		if num%p == 0 {
			return false
		}
	}
	return true
}
