public class euler278LinearCombinationsOfSemiprimes {

    public static void main(String[] args) {
        int limit = Integer.parseInt(args[0]);
        long sum = 0;

        boolean[] isPrime = generatePrimes(limit);

        for (int p = 2; p < limit; p++) {
            if (isPrime[p]) {
                for (int q = p + 1; q < limit; q++) {
                    if (isPrime[q]) {
                        for (int r = q + 1; r < limit; r++) {
                            if (isPrime[r]) {
                                long result = calculateLargestImpossibleValue(p, q, r);
                                sum += result;
                            }
                        }
                    }
                }
            }
        }

        System.out.println("Answer: " + sum);
    }

    private static boolean[] generatePrimes(int limit) {
        boolean[] isPrime = new boolean[limit];
        for (int i = 2; i < limit; i++) {
            isPrime[i] = true;
        }

        for (int i = 2; i * i < limit; i++) {
            if (isPrime[i]) {
                for (int j = i * i; j < limit; j += i) {
                    isPrime[j] = false;
                }
            }
        }

        return isPrime;
    }

    private static long calculateLargestImpossibleValue(int p, int q, int r) {
        return (long) 2 * p * q * r - p * q - p * r - q * r;
    }
}