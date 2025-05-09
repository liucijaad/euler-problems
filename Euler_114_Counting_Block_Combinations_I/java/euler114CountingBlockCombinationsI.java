public class euler114CountingBlockCombinationsI {
    public static void main(String[] args) {
        int size = Integer.parseInt(args[0]);
        System.out.println("Answer: " + checkBlocks(size));
        
    }

    static long checkBlocks(int size) {
        long[] possibleArr = new long[size + 1];
        possibleArr[0] = 1;
        possibleArr[1] = 1;
        possibleArr[2] = 1;
        for(int i=3; i<=size; i++) {
            long sum = possibleArr[i-1] + 1;
            for(int j=3; j<i; j++) {
                sum = sum + possibleArr[i - j - 1];
            }
            possibleArr[i] = sum;
        }
        return possibleArr[size];
    }
}