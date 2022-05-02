public class Solution {

    public static void main(String[] args) {
        ListNode l6 = new ListNode(6);
        ListNode l5 = new ListNode(5, l6);
        ListNode l4 = new ListNode(4, l5);
        ListNode l3 = new ListNode(3, l4);
        ListNode l2 = new ListNode(2, l3);
        ListNode l1 = new ListNode(1, l2);

        l1 = removeElements(l1, 3);

        while (l1 != null) {
            System.out.print(l1.val + " -> ");
            l1 = l1.next;
        }

        System.out.println("null");
    }

    public static ListNode removeElements(ListNode head, int val) {
        if (head == null) return null;

        ListNode prev = head;
        ListNode current = head;

        while (current != null) {
            if (current.val == val) {
                if (current == head) {
                    current = prev = head = head.next;
                }
                else {
                    prev.next = current.next;
                    current = current.next;
                }
            }
            else {
                prev = current;
                current = current.next;
            }
        }
        return head;
    }
}
