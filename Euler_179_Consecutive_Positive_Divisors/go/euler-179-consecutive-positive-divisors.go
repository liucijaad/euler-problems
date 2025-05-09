package main

import (
	"fmt"
	"os"
	"strconv"
	"math"
)

// Function to count the number of positive divisors for a given number
func countDivisors(n int) int {
	count := 0
	for i := 1; i <= int(math.Sqrt(float64(n))); i++ {
		if n % i == 0 {
			if n / i == i {
				count++
			} else {
				count += 2
			}
		}
	}
	return count
}

// Function to check if two numbers have the same number of positive divisors
func hasSameDivisors(a, b int) bool {
	return countDivisors(a) == countDivisors(b)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run euler-179-consecutive-positive-divisors.go <limit>")
		return
	}

	limit, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid input. Please provide a valid integer limit.")
		return
	}

	count := 0
	for i := 1; i < limit; i++ {
		if hasSameDivisors(i, i + 1) {
			count++
		}
	}

	fmt.Println("Answer:", count)
}