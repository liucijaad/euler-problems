package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

// Function to check if two numbers are coprime.
func areCoprime(a, b int) bool {
	return gcd(a, b) == 1
}

// Function to calculate the greatest common divisor (GCD) of two numbers.
func gcd(a, b int) int {
	for b != 0 {
		temp := b
		b = a % b
		a = temp
	}
	return a
}

// Function to check if a number is prime.
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Function to find the divisibility multiplier for a prime number.
func findDivisibilityMultiplier(prime, limit int) int {
	for x := 1; x < prime; x++ {
		found := true

		for n := 1; n < limit; n++ {
			fN := getFunctionValue(n, x)

			if (fN%prime != 0 && n%prime == 0) || (fN%prime == 0 && n%prime != 0) {
				found = false
				break
			}
		}

		if found {
			return x
		}
	}

	// Default value if no x is found (should not happen for primes).
	return -1
}

// Function to calculate the value of the given function f(n).
func getFunctionValue(n, x int) int {
	lastDigit := n % 10
	remainingDigits := n / 10
	return remainingDigits + (lastDigit * x)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run euler-274-divisibility-multipliers.go <limit>")
		return
	}

	limit, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid input. Please provide a valid integer limit.")
		return
	}

	sumOfDivisibilityMultipliers := 0

	for i := 1; i < limit; i++ {
		if areCoprime(i, 10) && isPrime(i) {
			x := findDivisibilityMultiplier(i, limit)
			sumOfDivisibilityMultipliers += x
		}
	}

	fmt.Println("Answer:", sumOfDivisibilityMultipliers)
}
