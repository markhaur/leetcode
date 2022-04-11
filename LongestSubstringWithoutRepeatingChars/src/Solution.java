public class Solution {

    public static void main(String[] args) {
        String s = "abcabcbb";

        int result = lengthOfLongestSubstring(s);

        System.out.println(result);

    }

    public static int lengthOfLongestSubstring(String s) {
        int size = s.length();

        if (size == 0)
            return 0;

        int sIndex = 0;
        int resultCount = 0;
        char[] sChar = s.toCharArray();

        for (int i = 0; i < size; i++) {
            for(int z = sIndex; z < i; z++) {
                if (sChar[i] == sChar[z]) {
                    sIndex = z+1;
                }
            }
            if (resultCount <= ( i - sIndex))
                resultCount = i - sIndex + 1;
        }

        return resultCount;
    }
}
