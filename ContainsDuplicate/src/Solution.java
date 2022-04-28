import java.util.HashMap;
import java.util.TreeMap;

public class Solution {

    public static void main(String[] args) {
        int[] nums = new int[]{1,2,3,4};
        System.out.println(containsDuplicate(nums));
    }

    public static boolean containsDuplicate(int[] nums) {
        TreeMap<Integer, Integer> map = new TreeMap<>();

        for( int num : nums) {
            if (map.containsKey(num))
                return true;
            map.put(num, 0);
        }

        return false;
    }
}
