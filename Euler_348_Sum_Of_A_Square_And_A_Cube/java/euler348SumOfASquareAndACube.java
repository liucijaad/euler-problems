public class euler348SumOfASquareAndACube {
    public static void main(String[] args) {
        int count = 0;
        long number = 1;
        long sumOfPalindromes = 0;
        int limit = Integer.parseInt(args[0]);

        while (count < limit) {
            if (isPalindromic(number) && countExpressions(number) == 4) {
                sumOfPalindromes += number;
                count++;
            }
            number++;
        }

        System.out.println("Answer: " + sumOfPalindromes);
    }

    private static boolean isPalindromic(long n) {
        String str = Long.toString(n);
        return str.equals(new StringBuilder(str).reverse().toString());
    }

    private static int countExpressions(long n) {
        int count = 0;

        for (long a = 2; a * a < n; a++) {
            for (long b = 2; a * a + b * b * b <= n; b++) {
                if (a * a + b * b * b == n) {
                    count++;
                }
            }
        }

        return count;
    }
}