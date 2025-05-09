import java.math.BigInteger;

public class euler20FactorialDigitSum {
    public static void main(String[] args) throws Exception {
        BigInteger factorial = BigInteger.valueOf(Integer.parseInt(args[0]));
        // If input < 0, throw an error.
        if (factorial.compareTo(BigInteger.ZERO) == -1) {
            throw new Exception("Input must be a positive integer or 0.");
        }
        // If 0 == 1 skip the calculations.
        else if (factorial.equals(BigInteger.ZERO)) {
            System.out.println("Answer: 1");
        } else {
            factorial = factorialOf(factorial.subtract(BigInteger.ONE), factorial);
            System.out.println("Answer: " + digitSum(factorial));
        }
    }

    /**
     * Recursive method to calculate the factorial of the given number.
     * 
     * @param number number, for which the factorial is being calculated
     * @param acc    accumulator for recursion
     * @return factorial of the number
     */
    static BigInteger factorialOf(BigInteger number, BigInteger acc) {
        if (number.equals(BigInteger.ZERO)) {
            return acc;
        } else {
            acc = acc.multiply(number);
            return factorialOf(number.subtract(BigInteger.ONE), acc);
        }
    }

    /**
     * Calculate the sum of the digits of the number.
     */
    static BigInteger digitSum(BigInteger number) {
        // Accumulator for the sum of the digits.
        BigInteger acc = BigInteger.ZERO;
        // Loop to get each digit out of the number using modulus operator
        for (BigInteger i = BigInteger.ONE; number.compareTo(BigInteger.ZERO) == 1;
                i = i.multiply(BigInteger.valueOf(10))) {
            BigInteger digit = number.mod(i.multiply(BigInteger.valueOf(10))).divide(i);
            acc = acc.add(digit);
            // Substract the digit we found from number to ensure next digits
            // extracted are correct and to keep loop from running forever.
            number = number.subtract(digit.multiply(i));
        }
        return acc;
    }
}
