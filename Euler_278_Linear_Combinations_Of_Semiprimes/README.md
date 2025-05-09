# [Euler problem 278](https://projecteuler.net/problem=278)

## Task Description:

Given the values of integers 1 < a_1 < a_2 < ... < a_n, consider the linear combination
q_1*a_1 + q_2*a_2 + ... + q_n*a_n = b, using only integer values q_k >= 0. 

Note that for a given set of a_k, it may be that not all values of b are possible.
For instance, if a_1 = 5 and a_2 = 7, there are no q_1 >= 0 and q_2 >= 0 such that b could be
1, 2, 3, 4, 6, 8, 9, 11, 13, 16, 18 or 23.

In fact, 23 is the largest impossible value of b for a_1=5 and a_2=7. We therefore call f(5, 7) = 23. Similarly, it can be shown that f(6, 10, 15)=29 and f(14, 22, 77) = 195.

Find the Sum of f(pq, pr, qr), where p, q and r are prime numbers and p < q < r < 5000.

## Example Invocation
### Java
$ javac euler_278_linear_combinations_of_semiprimes.java

$ java euler_278_linear_combinations_of_semiprimes 5000 > stdout.txt

### Go
$ go run euler-278-linear-combinations-of-semiprimes.go 5000 > stdout.txt

## Example Standard Input
5000

## Example Standard Output
Answer: 1228215747273908452
