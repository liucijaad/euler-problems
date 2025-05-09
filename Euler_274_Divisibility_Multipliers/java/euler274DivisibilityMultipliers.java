public class euler274DivisibilityMultipliers {
    public static void main(String[] args) {
        int limit = Integer.parseInt(args[0]);
        long sumOfDivisibilityMultipliers = 0;

        for (int i = 1; i < limit; i++) {
            if (areCoprime(i, 10) && isPrime(i)) {
                int x = findDivisibilityMultiplier(i, limit);
                sumOfDivisibilityMultipliers += x;
            }
        }

        System.out.println("Answer: " + sumOfDivisibilityMultipliers);
    }

    /** Check if two numbers are coprimes. */
    private static boolean areCoprime(int a, int b) {
        return gcd(a, b) == 1;
    }

    /** Calculate the greatest common divisor (GCD) of two numbers. */
    private static int gcd(int a, int b) {
        while (b != 0) {
            int temp = b;
            b = a % b;
            a = temp;
        }
        return a;
    }

    /** Check if a number is prime. */
    private static boolean isPrime(int n) {
        if (n <= 1) {
            return false;
        }
        for (int i = 2; i <= Math.sqrt(n); i++) {
            if (n % i == 0) {
                return false;
            }
        }
        return true;
    }

    /** Find the divisibility multiplier for a prime number. */
    private static int findDivisibilityMultiplier(int prime, int limit) {
        for (int x = 1; x < prime; x++) {
            boolean found = true;

            for (int n = 1; n < limit; n++) {
                int f_n = getFunctionValue(n, x);

                if ((f_n % prime != 0 && n % prime == 0) ||
                        (f_n % prime == 0 && n % prime != 0)) {
                    found = false;
                    break;
                }
            }

            if (found) {
                return x;
            }
        }

        // Default value if no x is found (should not happen for primes).
        return -1;
    }

    /** Calculate the value of the given function f(n). */
    private static int getFunctionValue(int n, int x) {
        int lastDigit = n % 10;
        int remainingDigits = n / 10;
        return remainingDigits + (lastDigit * x);
    }
}