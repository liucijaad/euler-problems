package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
)

func main() {
	number, _ := strconv.ParseFloat(os.Args[1], 64)
	// Check if number is a positive integer or 0.
	if !isIntegral(number) || number < 0 {
		panic("Error: Input must be a positive integer or 0.")
	} else if number == 0 { // If 0 == 1 return the answer.
		fmt.Println("Answer: 1")
	} else {
		fmt.Printf("Answer: %d\n", DigitSum(FactorialOf(big.NewInt(int64(
			number-1)), big.NewInt(int64(number)))))
	}
}

// Check if number is an integer.
func isIntegral(number float64) bool {
	return number == float64(int(number))
}

// Calculate the factorial of the given number.
func FactorialOf(number *big.Int, acc *big.Int) *big.Int {
	if number.Cmp(big.NewInt(1)) <= 0 {
		return acc
	} else {
		acc := acc.Mul(acc, number)
		return FactorialOf(number.Sub(number, big.NewInt(1)), acc)
	}
}

// Sum all digits of the factorial number.
func DigitSum(number *big.Int) *big.Int {
	// Accumulator for the sum of the digits.
	// Multiplicator and digit used to store values in loop for calculations.
	acc, multip, digit := big.NewInt(0), new(big.Int), new(big.Int)
	// Loop to get each digit out of the number using modulus operator.
	for i := big.NewInt(1); number.Cmp(i) >= 0; i.Mul(i, big.NewInt(10)) {
		digit.Set(number).Mod(digit, multip.Mul(i, big.NewInt(10))).Div(
			digit, i)
		acc.Add(acc, digit)
		number.Sub(number, digit.Mul(digit, i))
	}
	return acc
}
