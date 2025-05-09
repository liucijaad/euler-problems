package main

import (
	"fmt"
	"math"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := strconv.ParseInt(os.Args[1], 10, 64)
	fmt.Printf("Answer: %d\n", SumIrrRoots(FindIrrRoots(input)))
}

// Find all irregular roots from 1 up to numRange.
func FindIrrRoots(numRange int64) []string {
	var roots []string
	for i := range numRange {
		if math.Pow(float64(int64(math.Sqrt(float64(i)))), 2) != float64(i) {
			shift := big.NewInt(10)
			shift.Exp(shift, big.NewInt(100*2), nil)
			sqrt := big.NewInt(int64(i))
			sqrt.Mul(sqrt, shift)
			root := sqrt.Sqrt(sqrt).String()[:100]
			roots = append(roots, root)
		}
	}
	return roots
}

// Sum digit sums of all irregular roots.
func SumIrrRoots(roots []string) *big.Int {
	sum := big.NewInt(0)
	for i := range roots {
		dSum := DigitSum(roots[i])
		sum = sum.Add(sum, dSum)
	}
	return sum
}

// Sum all digits of the factorial number.
func DigitSum(number string) *big.Int {
	acc := big.NewInt(0)
	dList := strings.Split(number, "")
	for i := range len(number) {
		digit, _ := strconv.ParseInt(dList[i], 10, 64)
		acc.Add(acc, big.NewInt(digit))
	}
	return acc
}
