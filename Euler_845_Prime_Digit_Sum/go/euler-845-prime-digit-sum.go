package main

import (
	"fmt"
	"os"
	"strconv"
)

func calculateDigitSum(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

func generatePrimesBelow170() []int {
	var primes []int
	for i := 2; i < 170; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

// Function to check if a number is prime.
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func contains(arr []int, num int) bool {
	for _, v := range arr {
		if v == num {
			return true
		}
	}
	return false
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run euler-845-prime-digit-sum.go <limit>")
		return
	}

	limit, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid input. Please provide a valid integer limit.")
		return
	}

	primeNumbers := generatePrimesBelow170()
	number := 1
	count := 0

	for count < limit {
		digitSum := calculateDigitSum(number)

		if contains(primeNumbers, digitSum) {
			count++
		}

		number++
	}

	fmt.Println("Answer:", number-1)
}
