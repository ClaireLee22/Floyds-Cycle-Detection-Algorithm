package main

import (
	"fmt"
	"strings"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			slow = head

			for slow != fast {
				slow = slow.Next
				fast = fast.Next
			}
			return slow
		}
	}
	return nil
}

func removeCycle(head *ListNode) {
	cycleStartNode := detectCycle(head)
	if cycleStartNode != nil {
		pt1, pt2 := cycleStartNode, cycleStartNode
		for pt2.Next != pt1 {
			pt2 = pt2.Next
		}
		pt2.Next = nil
	}
	fmt.Println("linked list after removing cycle", printLinkedList(head))
}

func printLinkedList(head *ListNode) string {
	current := head
	nodes := []string{}
	for current != nil {
		nodes = append(nodes, fmt.Sprintf("%d", current.Val))
		current = current.Next
	}
	return strings.Join(nodes, " -> ")
}

func main() {
	node1 := &ListNode{Val: 3}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 0}
	node4 := &ListNode{Val: -4}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node2

	head := node1
	removeCycle(head)
}

// output: linked list after removing cycle 3 -> 2 -> 0 -> -4
