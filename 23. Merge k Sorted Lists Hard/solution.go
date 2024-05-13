package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}

	if l1 != nil {
		current.Next = l1
	} else {
		current.Next = l2
	}

	return dummy.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	for len(lists) > 1 {
		mergedLists := []*ListNode{}
		for i := 0; i < len(lists); i += 2 {
			if i+1 < len(lists) {
				mergedLists = append(mergedLists, mergeTwoLists(lists[i], lists[i+1]))
			} else {
				mergedLists = append(mergedLists, lists[i])
			}
		}
		lists = mergedLists
	}

	return lists[0]
}

func main() {
	lists := [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}

	linkedLists := []*ListNode{}
	for _, lst := range lists {
		head := &ListNode{Val: lst[0]}
		current := head
		for _, val := range lst[1:] {
			current.Next = &ListNode{Val: val}
			current = current.Next
		}
		linkedLists = append(linkedLists, head)
	}

	result := mergeKLists(linkedLists)

	output := []int{}
	for result != nil {
		output = append(output, result.Val)
		result = result.Next
	}
	fmt.Println(output)
}
