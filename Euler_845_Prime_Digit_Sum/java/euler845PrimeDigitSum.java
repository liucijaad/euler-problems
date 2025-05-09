import java.util.ArrayList;

public class euler845PrimeDigitSum {
    public static void main(String[] args) {

        ArrayList<Integer> primeNumbers = generatePrimesBelow170();
        int number = 1;
        long count = 0;
        int limit = Integer.parseInt(args[0]);

        while (count < limit) {
            int digitSum = calculateDigitSum(number);

            if (primeNumbers.contains(digitSum)) {
                count++;
            }

            number++;
        }

        System.out.println("Answer: " + (number - 1));
    }

    private static int calculateDigitSum(int n) {
        int sum = 0;
        while (n > 0) {
            sum += n % 10;
            n /= 10;
        }
        return sum;
    }

    private static ArrayList<Integer> generatePrimesBelow170() {
        ArrayList<Integer> primes = new ArrayList<>();
        for (int i = 2; i < 170; i++) {
            if (isPrime(i)) {
                primes.add(i);
            }
        }
        return primes;
    }

    /** Method to check if a number is prime. */
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
}