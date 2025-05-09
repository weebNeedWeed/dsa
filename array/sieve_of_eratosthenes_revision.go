package main

import "fmt"

func main() {
	printPrimes(12)
}

func printPrimes(n int) {
	primes := make([]bool, n+1)

	for i, _ := range primes {
		primes[i] = true
	}

	for division := 2; division*division <= n; division++ {
		if primes[division] {
			for j := 2 * division; j <= n; j += division {
				primes[j] = false
			}
		}
	}

	for i := 2; i <= n; i++ {
		if primes[i] {
			fmt.Println(i)
		}
	}
}
