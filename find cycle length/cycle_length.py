# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution(object):
    def findCycleLength(self, head):
        """
        :type head: ListNode
        :rtype: int
        """
    
        slow, fast = head, head
        while fast is not None and fast.next is not None:
            slow = slow.next
            fast = fast.next.next
            if slow == fast:
                return self.calculateCycleLength(slow)
            
        return 0 

    def calculateCycleLength(self, slow):
        current = slow
        cycleLength = 0
        while True:
            current = current.next
            cycleLength += 1
            if current == slow:
                break
        return cycleLength  
        
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
    print("cycle length", solution.findCycleLength(head))

"""
output: ('cycle length', 3)

"""