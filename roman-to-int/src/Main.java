import java.util.HashMap;
import java.util.Map;

public class Main {

    public static void main(String[] args) {

        String romanString = "MCMXCIV";

        Map<String, String> replacements = new HashMap<>();
        replacements.put("IV", "a");
        replacements.put("IX", "b");
        replacements.put("XL", "c");
        replacements.put("XC", "d");
        replacements.put("CD", "e");
        replacements.put("CM", "f");
        int integerNumber = 0;
        Map<Character, Integer> romanMap = new HashMap<>();
        romanMap.put('I', 1);
        romanMap.put('a', 4);
        romanMap.put('V', 5);
        romanMap.put('b', 9);
        romanMap.put('X', 10);
        romanMap.put('c', 40);
        romanMap.put('L', 50);
        romanMap.put('d', 90);
        romanMap.put('C', 100);
        romanMap.put('e', 400);
        romanMap.put('D', 500);
        romanMap.put('f', 900);
        romanMap.put('M', 1000);

        for (String k : replacements.keySet()) {
            romanString = romanString.replaceAll(k, replacements.get(k));
        }

        for( int i = 0; i < romanString.length(); i++ ){
            integerNumber += romanMap.get(romanString.charAt(i));
        }

        System.out.println(integerNumber);

    }
}
