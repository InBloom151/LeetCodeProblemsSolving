class Solution {
    public static void main(String[] args) {
        System.out.println(convert("PAYPALISHIRING", 3));
        System.out.println(convert("PAYPALISHIRING", 4));
        System.out.println(convert("A", 1));
    }

    public static String convert(String s, int numRows) {
        if (numRows >= s.length() || numRows == 1) {
            return s;
        }

        char[] result = new char[s.length()];
        int rowSize = 2 * numRows - 2;
        int index = 0;

        for (int i = 0; i < numRows; i++) {
            for (int j = i; j < s.length(); j += rowSize) {
                result[index++] = s.charAt(j);
                if (i != 0 && i != numRows - 1) {
                    int diagonal = j + rowSize - 2 * i;
                    if (diagonal < s.length()) {
                        result[index++] = s.charAt(diagonal);
                    }
                }
            }
        }
        return new String(result);
    }
}