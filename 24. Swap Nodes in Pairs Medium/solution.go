package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head1 := createLinkedList([]int{1, 2, 3, 4})
	fmt.Println(linkedListToList(swapPairs(head1)))

	head2 := createLinkedList([]int{})
	fmt.Println(linkedListToList(swapPairs(head2)))

	head3 := createLinkedList([]int{1})
	fmt.Println(linkedListToList(swapPairs(head3)))
}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	prev := dummy

	for prev.Next != nil && prev.Next.Next != nil {
		first := prev.Next
		second := prev.Next.Next

		prev.Next = second
		first.Next = second.Next
		second.Next = first

		prev = first
	}

	return dummy.Next
}

func createLinkedList(lst []int) *ListNode {
	dummy := &ListNode{}
	current := dummy
	for _, val := range lst {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}
	return dummy.Next
}

func linkedListToList(head *ListNode) []int {
	lst := []int{}
	current := head
	for current != nil {
		lst = append(lst, current.Val)
		current = current.Next
	}
	return lst
}
