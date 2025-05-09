import java.math.BigDecimal;
import java.math.MathContext;
import java.util.ArrayList;
import java.util.List;

public class euler80SquareRootDigitalExpansion {
    public static void main(String[] args) {
        List<BigDecimal> roots = findIrrRoot(Integer.parseInt(args[0]));
        System.out.println("Answer: " + sumIrrRootDigits(roots));
    }

    /** Find all irrational roots and store first 100 decimals in list. */
    static List<BigDecimal> findIrrRoot(int range) {
        // List to store irrational root digit values and precision for calculations.
        List<BigDecimal> roots = new ArrayList<BigDecimal>();
        MathContext precision = new MathContext(105);
        for (int i = 1; i <= range; i++) {
            if (Math.pow((int) (Math.sqrt(i)), 2) != i) {
                // Decimal value of irrational root. Modulus 1 to get 0.
                // followed by digits
                String decValue = BigDecimal.valueOf(i)
                        .multiply(BigDecimal.TEN.pow(100 * 2))
                        .sqrt(precision)
                        .toString();
                // Add substring that leaves out 0 of decValue.
                roots.add(new BigDecimal(decValue.substring(0, 100)));
            }
        }
        return roots;
    }

    /** Calculate sum of irrational root digits. */
    static BigDecimal sumIrrRootDigits(List<BigDecimal> roots) {
        BigDecimal sum = BigDecimal.ZERO;
        for (int i = 0; i < roots.size(); i++) {
            BigDecimal dSum = digitSum(roots.get(i));
            sum = sum.add(dSum);
        }
        return sum;
    }

    /** Sum digits of given number. */
    static BigDecimal digitSum(BigDecimal number) {
        // Accumulator for the sum of the digits.
        BigDecimal acc = BigDecimal.ZERO;
        // Loop to get each digit out of the number using modulus operator.
        for (BigDecimal i = BigDecimal.ONE; number.compareTo(BigDecimal.ZERO) == 1; i = i
                .multiply(BigDecimal.TEN)) {
            BigDecimal digit = number.remainder(i.multiply(BigDecimal.TEN)).divide(i);
            acc = acc.add(digit);
            // Substract the digit we found from number to ensure next digits
            // extracted are correct and to keep loop from running forever.
            number = number.subtract(digit.multiply(i));
        }
        return acc;
    }
}
