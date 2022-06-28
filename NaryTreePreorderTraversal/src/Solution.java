import java.util.ArrayList;
import java.util.List;

public class Solution {

    public static void main(String[] args) {
        
    }

    public static List<Integer> preorder(Node root) {
        List<Integer> list = new ArrayList<>();
        traverse(root, list);
        return list;
    }

    private static void traverse(Node root, List<Integer> list) {

        if (root == null)
            return;

        list.add(root.val);

        for (Node n : root.children) {
            traverse(n, list);
        }
    }
}
