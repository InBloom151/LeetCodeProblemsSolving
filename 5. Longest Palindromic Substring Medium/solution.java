public class Solution {
    public static void main(String[] args) {
        System.out.println(palindromSub("babad"));
        System.out.println(palindromSub("cbbd"));
    }

    public static String preprocess(String s) {
        StringBuilder processed = new StringBuilder("#");
        for (int i = 0; i < s.length(); i++) {
            processed.append(s.charAt(i)).append("#");
        }
        return processed.toString();
    }

    public static String palindromeSub(String s) {
        if (s.length() < 1) {
            return "";
        }

        String processedS = preprocess(s);
        int n = processedS.length();
        int[] P = new int[n];
        int C = 0, R = 0;

        for (int i = 1; i < n - 1; i++) {
            int mirror = 2 * C - i;
            if (i < R) {
                P[i] = Math.min(R - i, P[mirror]);
            }

            while (i + (1 + P[i]) < n && i - (1 + P[i]) >= 0 && processedS.charAt(i + (1 + P[i])) == processedS.charAt(i - (1 + P[i]))) {
                P[i]++;
            }

            if (i + P[i] > R) {
                C = i;
                R = i + P[i];
            }
        }

        int maxLen = 0;
        int centerIndex = 0;
        for (int i = 0; i < n; i++) {
            if (P[i] > maxLen) {
                maxLen = P[i];
                centerIndex = i;
            }
        }

        int startIndex = (centerIndex - maxLen) / 2;
        return s.substring(startIndex, startIndex + maxLen);
    }
}