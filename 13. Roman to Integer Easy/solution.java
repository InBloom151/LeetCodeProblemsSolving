import java.util.Map;
import java.util.HashMap;

class solution {
    public static void main(String[] args) {
        System.out.println(romanToInt("III"));
        System.out.println(romanToInt("LVIII"));
        System.out.println(romanToInt("MCMXCIV"));
    }

    public static int romanToInt(String s) {
        Map<String, Integer> ASSOCIATIONS = new HashMap<>();
        ASSOCIATIONS.put("I", 1);
        ASSOCIATIONS.put("V", 5);
        ASSOCIATIONS.put("X", 10);
        ASSOCIATIONS.put("L", 50);
        ASSOCIATIONS.put("C", 100);
        ASSOCIATIONS.put("D", 500);
        ASSOCIATIONS.put("M", 1000);
    
        int total = 0;
        int prevValue = 0;
    
        for (int numeral = s.length() - 1; numeral >= 0; numeral--) {
            int value = ASSOCIATIONS.get(String.valueOf(s.charAt(numeral)));
    
            total = value < prevValue ? total - value : total + value;
    
            prevValue = value;
        }
    
        return total;
    }
}