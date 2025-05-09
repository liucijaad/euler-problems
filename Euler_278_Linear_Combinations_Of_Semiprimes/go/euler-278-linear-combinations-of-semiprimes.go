package main

import (
    "fmt"
    "os"
    "strconv"
)

func generatePrimes(limit int)[] bool {
    isPrime := make([] bool, limit)
    for i := 2;i < limit;i++{
        isPrime[i] = true
    }

    for i := 2;i * i < limit;i++{
        if isPrime[i] {
            for j := i * i;
            j < limit;
            j += i {
                isPrime[j] = false
            }
        }
    }

    return isPrime
}

func calculateLargestImpossibleValue(p, q, r int64) int64 {
    return 2 * p * q * r - p * q - p * r - q * r
}

func main() {
    if len(os.Args) != 2 {
        fmt.Println(
            "Usage: go run euler-278-linear-combinations-of-semiprimes.go <limit>")
        return
    }

    limit, err := strconv.Atoi(os.Args[1])
    if err != nil {
        fmt.Println(
            "Invalid input. Please provide a valid integer limit.")
        return
    }

    sum := int64(0)
    isPrime := generatePrimes(limit)

    for p := 2;
    p < limit;
    p++{
        if isPrime[p] {
            for q := p + 1;
            q < limit;
            q++{
                if isPrime[q] {
                    for r := q + 1;
                    r < limit;
                    r++{
                        if isPrime[r] {
                            result :=
                                calculateLargestImpossibleValue(
                                    int64(p), int64(q), int64(r))
                            sum += result
                        }
                    }
                }
            }
        }
    }

    fmt.Println("Answer:", sum)
}
