class solution {
    public static void main(String[] args) {
        System.out.println(intToRoman(3749));
        System.out.println(intToRoman(58));
        System.out.println(intToRoman(1994));
    }

    public static String intToRoman(int num) {
        int[] nums = {1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1};
        String[] romans = {"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"};

        StringBuilder result = new StringBuilder();

        // Iterate through associations
        for (int i = 0; i < nums.length; i++) {
            int k = nums[i];
            String v = romans[i];
            while (num >= k) {
                result.append(v);
                num -= k;
            }
        }

        return result.toString();
    }
}