import java.util.ArrayList;
import java.util.List;
import java.util.Optional;

public class Solution {

    static boolean doBreak = false;
    static int size = 1;
    static int counter = 0;

    public static void main(String[] args) {
        ListNode l7 = new ListNode(7);
        ListNode l6 = new ListNode(6, l7);
        ListNode l5 = new ListNode(5, l6);
        ListNode l4 = new ListNode(4, l5);
        ListNode l3 = new ListNode(3, l4);
        ListNode l2 = new ListNode(2, l3);
        ListNode l1 = new ListNode(1, l2);

        reorderList(l1);

        ListNode head = l1;
        for (int i = 0; i < 7; i++) {
            System.out.print(head.val + " -> ");
            head = head.next;
        }
        System.out.println("null");
    }

    private static void reorderList(ListNode head) {
        reorderBySplitting(head);
    }

    private static void reorderBySplitting(ListNode head) {
        if (head == null || head.next == null) return;

        // head of first half
        ListNode l1 = head;

        // head of second half
        ListNode slow = head;

        // tail of second half
        ListNode fast = head;

        // tail of first half
        ListNode prev = null;

        // splitting the Singly Linked List into two Linked List
        while (fast != null && fast.next != null) {
            prev = slow;
            slow = slow.next;
            fast = fast.next.next;
        }

        // marking the end of for first list
        prev.next = null;

        ListNode l2 = reverse(slow);

        merge(l1, l2);
    }

    private static void merge(ListNode l1, ListNode l2) {
        while (l1 != null) {
            ListNode l1_next = l1.next;
            ListNode l2_next = l2.next;

            l1.next = l2;

            if (l1_next == null)
                break;

            l2.next = l1_next;
            l1 = l1_next;
            l2 = l2_next;
        }
    }

    private static ListNode reverse(ListNode head) {
        ListNode prev = null;
        ListNode currentNode = head;

        while (currentNode != null) {
            ListNode nextNode = currentNode.next;
            currentNode.next = prev;
            prev = currentNode;
            currentNode = nextNode;
        }

        return prev;
    }

    // solution using arraylist
    private static void reorder(ListNode head) {
        List<ListNode> list = new ArrayList<>();
        ListNode current = head;
        while(current != null) {
            list.add(current);
            current = current.next;
        }

        int reverseCounter = list.size() - 1;
        int loopLimit = (list.size() / 2) + (list.size()%2) - 1;

        for(int i = 0; i < loopLimit; i++) {
            ListNode temp = list.get(i).next;
            list.get(i).next = list.get(reverseCounter);
            list.get(reverseCounter).next = temp;
            reverseCounter--;
        }

        list.get(list.size() / 2).next = null;
    }

    // solution using recursion
    private static ListNode reorder(ListNode head, ListNode other) {
        if (other == null)
            return null;
        size++;
        ListNode sentinal = reorder(head, other.next);

        if (doBreak)
            return null;

        int loopLimit = (size / 2) + (size%2) - 1;
        if (counter == loopLimit) {
            other.next = null;
            doBreak = true;
            return null;
        }

        counter++;

        if (sentinal == null) {
            other.next = head.next;
            head.next = other;
        }
        else {
            other.next = sentinal.next;
            sentinal.next = other;
        }
        return other.next;
    }
}
