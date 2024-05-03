class solution {
    public static boolean isMatch(String s, String p) {
        int m = s.length();
        int n = p.length();
        boolean[][] memo = new boolean[m + 1][n + 1];
        memo[m][n] = true;

        for (int i = m; i >= 0; i--) {
            for (int j = n - 1; j >= 0; j--) {
                boolean firstMatch = i < m && (p.charAt(j) == s.charAt(i) || p.charAt(j) == '.');
                if (j + 1 < n && p.charAt(j + 1) == '*') {
                    memo[i][j] = memo[i][j + 2] || (firstMatch && memo[i + 1][j]);
                } else {
                    memo[i][j] = firstMatch && memo[i + 1][j + 1];
                }
            }
        }

        return memo[0][0];
    }

    public static void main(String[] args) {
        System.out.println(isMatch("aa", "a"));
        System.out.println(isMatch("aa", "a*"));
        System.out.println(isMatch("ab", ".*"));
    }
}