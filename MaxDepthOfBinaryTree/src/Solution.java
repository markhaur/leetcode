/**
 * Following tree structure is constructed in the main method
 *                      3
 *                    /   \
 *                   9     20
 *                        /   \
 *                      15      7
 */
public class Solution {

    public static void main(String[] args) {

        TreeNode root = new TreeNode(3);
        TreeNode node9 = new TreeNode(9);
        TreeNode node20 = new TreeNode(20);
        TreeNode node15 = new TreeNode(15);
        TreeNode node7 = new TreeNode(7);

        root.left = node9;
        root.right = node20;
        node20.left = node15;
        node20.right = node7;

        System.out.println("Max Depth from Root node: " + maxDepth(root));        // max depth would be 3
        System.out.println("Max Depth from Node 20  : " + maxDepth(node20));        // max depth would be 2
        System.out.println("Max Depth from Node 9   : " + maxDepth(node9));        // max depth would be 1

    }

    public static int maxDepth(TreeNode root) {

        if (root == null)
            return 0;

        if (root.left == null && root.right == null)
            return 1;

        int leftSize = maxDepth(root.left);
        int rightSize = maxDepth(root.right);

        if (leftSize < rightSize) {
            return rightSize + 1;
        }

        return leftSize + 1;
    }
}
