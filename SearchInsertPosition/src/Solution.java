public class Solution {

    public static void main(String[] args) {

        int[] nums = new int[]{1,3,5,6};

        System.out.println(searchInsert(nums, 2));

    }

    public static int searchInsert(int[] nums, int target) {
        int lo = 0;
        int hi = nums.length-1;
        int mid = 0;
        while (lo <= hi) {
            mid = lo + (hi - lo) / 2;
            if (target < nums[mid])
                hi = mid - 1;
            else if (target > nums[mid])
                lo = mid + 1;
            else {
                return mid;
            }
            if (lo > hi) {
                return lo;
            }
        }
        return -1;
    }
}
