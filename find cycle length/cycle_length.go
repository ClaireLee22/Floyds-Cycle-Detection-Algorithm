package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func findCycleLength(head *ListNode) int {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return calculateCycleLength((slow))
		}
	}
	return 0
}

func calculateCycleLength(slow *ListNode) int {
	current := slow
	cycleLength := 0
	for {
		current = current.Next
		cycleLength += 1
		if current == slow {
			break
		}
	}
	return cycleLength
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
	fmt.Println("cycle length:", findCycleLength(head))
}

/*
output: cycle length: 3
*/
