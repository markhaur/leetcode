public class Solution {

    public static void main(String[] args) {
        String input = "PAYPALISHIRING";

        System.out.println(convert(input, 3));
    }

    public static String convert(String s, int numRows) {

        if (numRows <= 1 || s == null || s.length() < numRows)
            return s;

        int size = s.length();
        char[] sChars = s.toCharArray();
        char[] result = new char[size];

        int resultPointer = 0;
        int number = (numRows - 1) * 2;

        int first = 0;
        int second = 0;

        for (int i = 0; i < numRows; i++) {

            first = number - i;
            second = number + i;

            result[resultPointer++] = sChars[i];

            while ((first < size || second < size) && resultPointer < size) {

                if (first == second) {
                    result[resultPointer++] = sChars[first];
                }
                else {
                    if (first != i && first < size) {
                        result[resultPointer++] = sChars[first];
                    }
                    if (second != i && second < size) {
                        result[resultPointer++] = sChars[second];
                    }
                }

                if (numRows - i != 1)
                    first += number;
                second += number;
            }
        }

        return new String(result);
    }
}
