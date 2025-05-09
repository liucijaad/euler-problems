# [Euler Problem 80](https://projecteuler.net/problem=80)

## Task Description:

It is well known that if the square root of a natural number is not an integer, 
then it is irrational. The decimal expansion of such square roots is infinite without any 
repeating pattern at all.

The square root of two is 1.41421356237309504880..., and the digital sum of the first one 
hundred decimal digits is 475.

For the first one hundred natural numbers, find the total of the digital sums of the first 
one hundred decimal digits for all the irrational square roots.

## Example Invocation
### Java
$ javac euler80SquareRootDigitalExpansion.java

$ java euler80SquareRootDigitalExpansion 100 > stdout.txt

### Go
$ go run euler-80-square-root-digital-expansion.go 100 > stdout.txt

## Example Standard Input
100

## Example Standard Output
Answer: 40886

## Explanation
To get 100 digits of square root of i, we multiply i by 10^200. Then we extract first 100 digits and sum them up.
Repeat this process while i < input. Return sum of root 100 digit sums.
