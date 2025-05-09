# [Euler problem 20](https://projecteuler.net/problem=20)

## Task Description:

n! means n x (n - 1) x ... x 3 x 2 x 1.

For example, 10! = 10 x 9 x ... x 3 x 2 x 1 = 3628800,
and the sum of the digits of the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.

Find the sum of the digits in the number 100!.

## Example Invocation
### Java
$ javac euler20FactorialDigitSum.java

$ java euler20FactorialDigitSum 10 > stdout.txt

### Go
$ go run euler-20-factorial-digit-sum.go 10 > stdout.txt

## Example Standard Input
10

## Example Standard Output
27

## Explanation
Calculate the factorial of the number. Extract each digit while adding them up in accumulator. Return accumulator as the answer.
