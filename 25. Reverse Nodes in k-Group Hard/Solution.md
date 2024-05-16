### This approach has a time complexity of `O(n)`

### 1. Python

**Runtime:** `40 ms` faster than `72.20%` submissions  
**Memory usage:** `17.5 MB` less than `73.45%` submissions  

``` python
from typing import Optional

class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

def reverseKGroup(head: Optional[ListNode], k: int) -> Optional[ListNode]:
    def reverseLinkedList(head: Optional[ListNode]) -> Optional[ListNode]:
        prev = None
        current = head
        while current:
            next_node = current.next
            current.next = prev
            prev = current
            current = next_node
        return prev
    
    def countNodes(node: Optional[ListNode]) -> int:
        count = 0
        while node:
            count += 1
            node = node.next
        return count
    
    dummy = ListNode(0)
    dummy.next = head
    prev_group_end = dummy
    
    nodes_count = countNodes(head)
    
    while nodes_count >= k:
        group_start = prev_group_end.next
        current = group_start
        for _ in range(k - 1):
            current = current.next
        
        group_end = current
        next_group_start = group_end.next
        
        group_end.next = None
        prev_group_end.next = reverseLinkedList(group_start)
        group_start.next = next_group_start
        
        prev_group_end = group_start
        nodes_count -= k
    
    return dummy.next
```

### 2. TypeScript

**Runtime:** `71 ms` faster than `73.09%` submissions  
**Memory usage:** `54.3 MB` less than `85.02%` submissions  

``` typescript
class ListNode {
    val: number;
    next: ListNode | null;

    constructor(val: number, next: ListNode | null = null) {
        this.val = val;
        this.next = next;
    }
}

function reverseKGroup(head: ListNode | null, k: number): ListNode | null {
    function reverseLinkedList(head: ListNode | null): ListNode | null {
        let prev: ListNode | null = null;
        let current: ListNode | null = head;
        while (current !== null) {
            let nextNode: ListNode | null = current.next;
            current.next = prev;
            prev = current;
            current = nextNode;
        }
        return prev;
    }

    function countNodes(node: ListNode | null): number {
        let count: number = 0;
        while (node !== null) {
            count++;
            node = node.next;
        }
        return count;
    }

    let dummy: ListNode = new ListNode(0);
    dummy.next = head;
    let prevGroupEnd: ListNode = dummy;

    let nodesCount: number = countNodes(head);

    while (nodesCount >= k) {
        let groupStart: ListNode | null = prevGroupEnd.next;
        let current: ListNode | null = groupStart;
        for (let i = 0; i < (k - 1); i++) {
            current = current!.next;
        }

        let groupEnd: ListNode | null = current;
        let nextGroupStart: ListNode | null = groupEnd!.next;

        groupEnd!.next = null;
        prevGroupEnd.next = reverseLinkedList(groupStart);
        groupStart.next = nextGroupStart;

        prevGroupEnd = groupStart;
        nodesCount -= k;
    }

    return dummy.next;
}
```

### 3. GO

**Runtime:** `4 ms` faster than `71.67%` submissions  
**Memory usage:** `3.7 MB` less than `81.18%` submissions  

``` go
type ListNode struct {
	Val  int
	Next *ListNode
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
```