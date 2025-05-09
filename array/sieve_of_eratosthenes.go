package main

import "fmt"

func main() {
	printPrimes(100)
}

func printPrimes(n int) {
	primes := make([]bool, n+1, n+1)

	for i := 1; i <= n; i++ {
		primes[i] = true
	}

	for divisor := 2; divisor*divisor <= n; divisor++ {
		if primes[divisor] {
			for i := divisor * 2; i <= n; i += divisor {
				primes[i] = false
			}
		}
	}

	for i := 1; i <= n; i++ {
		if primes[i] {
			fmt.Println(i)
		}
	}
}
