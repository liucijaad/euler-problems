public class euler179ConsecutivePositiveDivisors {
    public static void main(String[] args) {
        int count = 0;
        int limit = Integer.parseInt(args[0]);

        for (int i = 1; i < limit; i++) {
            if (hasSameDivisors(i, i + 1)) {
                count++;
            }
        }

        System.out.println("Answer: " + count);
    }

    /** Count the number of positive divisors for a given number. */
    private static int countDivisors(int n) {
        int count = 0;
        for (int i = 1; i <= Math.sqrt(n); i++) {
            if (n % i == 0) {
                if (n / i == i) {
                    count++;
                } else {
                    count += 2;
                }
            }
        }
        return count;
    }

    /** Check if two numbers have the same number of positive divisors. */
    private static boolean hasSameDivisors(int a, int b) {
        return countDivisors(a) == countDivisors(b);
    }
}