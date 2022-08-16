import java.util.Arrays;

public class Main {

    public static void main(String[] args) {
        String[] strings = {"flower", "flow", "Aflight"};
        System.out.println(longestCommonPrefix(strings));
    }

    public static String longestCommonPrefix(String[] strs) {
        char[] prefix = null;
        for(int i = 0; i< strs.length; i++) {
            prefix = union(prefix, strs[i]);
            if (prefix == null)
                return "";
        }
        return new String(prefix);
    }

    private static char[] union(char[] prefix, String other) {
        if (prefix == null) {
            return other.toCharArray();
        }
        int index;
        int size = other.length() < prefix.length ? other.length() : prefix.length;
        for (index = 0; index < size; index++) {
            if (prefix[index] != other.charAt(index)){
                break;
            }
        }
        if (index == 0)
            return null;
        char[] result = new char[index];
        for(int i = 0; i < index; i++)
            result[i] = prefix[i];
        return result;
    }
}
