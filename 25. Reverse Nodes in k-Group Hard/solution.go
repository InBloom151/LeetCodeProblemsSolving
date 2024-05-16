package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head1 := &ListNode{Val: 1}
	head1.Next = &ListNode{Val: 2}
	head1.Next.Next = &ListNode{Val: 3}
	head1.Next.Next.Next = &ListNode{Val: 4}
	head1.Next.Next.Next.Next = &ListNode{Val: 5}

	k1 := 2
	fmt.Printf("Input: [1,2,3,4,5], k = %d\n", k1)
	result1 := reverseKGroup(head1, k1)
	for result1 != nil {
		fmt.Printf("%d ", result1.Val)
		result1 = result1.Next
	}
	fmt.Println()

	head2 := &ListNode{Val: 1}
	head2.Next = &ListNode{Val: 2}
	head2.Next.Next = &ListNode{Val: 3}
	head2.Next.Next.Next = &ListNode{Val: 4}
	head2.Next.Next.Next.Next = &ListNode{Val: 5}

	k2 := 3
	fmt.Printf("Input: [1,2,3,4,5], k = %d\n", k2)
	result2 := reverseKGroup(head2, k2)
	for result2 != nil {
		fmt.Printf("%d ", result2.Val)
		result2 = result2.Next
	}
	fmt.Println()
}

func reverseLinkedList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	current := head
	for current != nil {
		nextNode := current.Next
		current.Next = prev
		prev = current
		current = nextNode
	}
	return prev
}

func countNodes(node *ListNode) int {
	count := 0
	for node != nil {
		count++
		node = node.Next
	}
	return count
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	prevGroupEnd := dummy

	nodesCount := countNodes(head)

	for nodesCount >= k {
		groupStart := prevGroupEnd.Next
		current := groupStart
		for i := 0; i < k-1; i++ {
			current = current.Next
		}

		groupEnd := current
		nextGroupStart := groupEnd.Next

		groupEnd.Next = nil
		prevGroupEnd.Next = reverseLinkedList(groupStart)
		groupStart.Next = nextGroupStart

		prevGroupEnd = groupStart
		nodesCount -= k
	}

	return dummy.Next
}
