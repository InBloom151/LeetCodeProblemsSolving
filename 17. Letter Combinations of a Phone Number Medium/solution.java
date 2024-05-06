import java.util.ArrayList;
import java.util.List;

class solution {
    public static void main(String[] args) {
        System.out.println(letterCombinations("23"));
        System.out.println(letterCombinations(""));
        System.out.println(letterCombinations("2"));
    }

    private static final String[] PHONE_NUMS = {
        "0", "1", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"
    };
    
    public static List<String> letterCombinations(String digits) {
        List<String> result = new ArrayList<>();
        if (digits == null || digits.length() == 0) return result;
        
        backtrack(result, digits, 0, new StringBuilder());
        return result;
    }
    
    public static void backtrack(List<String> result, String digits, int index, StringBuilder combination) {
        if (index == digits.length()) {
            result.add(combination.toString());
            return;
        }
        
        int digit = digits.charAt(index) - '0';
        String letters = PHONE_NUMS[digit];
        for (char letter : letters.toCharArray()) {
            combination.append(letter);
            backtrack(result, digits, index + 1, combination);
            combination.deleteCharAt(combination.length() - 1);
        }
    }
} 
