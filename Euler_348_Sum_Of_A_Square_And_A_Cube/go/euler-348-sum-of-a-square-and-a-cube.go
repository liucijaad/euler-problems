package main

import (
	"fmt"
	"os"
	"strconv"
)

func isPalindromic(n int64) bool {
	str := strconv.FormatInt(n, 10)
	return str == reverseString(str)
}

// From https://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func countExpressions(n int64) int {
	count := 0

	for a := int64(2); a*a < n; a++ {
		for b := int64(2); a*a+b*b*b <= n; b++ {
			if a*a+b*b*b == n {
				count++
			}
		}
	}

	return count
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run euler-348-sum-of-a-square-and-a-cube.go <limit>")
		return
	}

	limit, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid input. Please provide a valid integer limit.")
		return
	}

	count := 0
	number := int64(1)
	sumOfPalindromes := int64(0)

	for count < limit {
		if isPalindromic(number) && countExpressions(number) == 4 {
			sumOfPalindromes += number
			count++
		}
		number++
	}

	fmt.Println("Answer:", sumOfPalindromes)
}
