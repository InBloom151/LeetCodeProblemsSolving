import java.util.Arrays;

public class solution {
    public static int longestSubstringLength(String s) {
        int[] lastOccurrence = new int[256];
        Arrays.fill(lastOccurrence, -1);
        int start = 0;
        int maxLength = 0;

        for (int i = 0; i < s.length(); i++) {
            char currentChar = s.charAt(i);
            int charCode = (int) currentChar;
            if (lastOccurrence[charCode] >= start) {
                start = lastOccurrence[charCode] + 1;
            }
            lastOccurrence[charCode] = i;
            maxLength = Math.max(maxLength, i - start + 1);
        }

        return maxLength;
    }

    public static void main(String[] args) {
        System.out.println(longestSubstringLength("abcabcbb"));
        System.out.println(longestSubstringLength("bbbbb"));
        System.out.println(longestSubstringLength("pwwkew"));
    }
}
