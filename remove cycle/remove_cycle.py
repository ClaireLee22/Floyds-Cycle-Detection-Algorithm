# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution(object):
    def detectCycle(self, head):
        """
        :type head: ListNode
        :rtype: ListNode
        """
        
        slow, fast = head, head
        while fast is not None and fast.next is not None:
            fast = fast.next.next
            slow = slow.next
            if fast == slow:
                slow = head
                while slow != fast:
                    slow = slow.next
                    fast = fast.next

                return slow
            
        return None
    
    def removeCycle(self, head):
        cycleStartNode = self.detectCycle(head)
        if cycleStartNode is not None:
            pt1 = cycleStartNode
            pt2 = cycleStartNode
            while pt2.next != pt1:
                pt2 = pt2.next
            
            pt2.next = None

        print("linked list after removing cycle", printLinkedList(head))
    

def printLinkedList(head):
    current = head
    nodes = []
    while current is not None:
        nodes.append(str(current.val))
        current = current.next
    return " -> ".join(nodes)


if __name__ == "__main__":
    node1 = ListNode(3)
    node2 = ListNode(2)
    node3 = ListNode(0)
    node4 = ListNode(-4)
    node1.next = node2
    node2.next = node3
    node3.next = node4
    node4.next = node2

    solution = Solution()
    head = node1
    solution.removeCycle(head)

    # output: ('linked list after removing cycle', '3 -> 2 -> 0 -> -4')