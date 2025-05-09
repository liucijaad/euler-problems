package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	size, _ := strconv.ParseInt(os.Args[1], 10, 64)
	fmt.Printf("Answer: %d\n", CheckBlocks(int(size)))
}

// possibleArr[i] - an array to store number of possible combinations of tile of length i
// i.e. possibleArr[1] = tile with length of 1 and this has only 1 possible combination (all gray)
// Iterate through all of sizes up to needed size to find the amount of combinations.
func CheckBlocks(size int) int64 {
	possibleArr := [51]int64{1, 1, 1}
	for i := 3; i <= size; i++ {
		sum := possibleArr[i-1] + 1
		for j := 3; j < i; j++ {
			sum = sum + possibleArr[i-j-1]
		}
		possibleArr[i] = sum
	}
	return possibleArr[size]
}
