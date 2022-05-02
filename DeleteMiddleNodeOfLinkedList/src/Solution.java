public class Solution {

    public static void main(String[] args) {
        ListNode l6 = new ListNode(6);
        ListNode l5 = new ListNode(5, l6);
        ListNode l4 = new ListNode(4, l5);
        ListNode l3 = new ListNode(3, l4);
        ListNode l2 = new ListNode(2, l3);
        ListNode l1 = new ListNode(1, l2);

        l1 = deleteMiddle(l1);

        while (l1 != null) {
            System.out.print(l1.val + " -> ");
            l1 = l1.next;
        }

        System.out.println("null");
    }

    public static ListNode deleteMiddle(ListNode head) {

        if (head == null || head.next == null)
            return null;

        ListNode l1 = head;
        ListNode slow = head;
        ListNode fast = head;

        while(fast != null && fast.next != null) {
            l1 = slow;
            slow = slow.next;
            fast = fast.next.next;
        }

        l1.next = slow.next;

        return head;
    }
}
