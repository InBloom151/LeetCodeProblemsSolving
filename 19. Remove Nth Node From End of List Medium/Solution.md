### This approach has a time complexity of `O(n)`

### 1. Python

**Runtime:** `28 ms` faster than `94.77%` submissions  
**Memory usage:** `16.6 MB` less than `45.51%` submissions  

``` python
def remove_nth_from_end(head: Optional[ListNode], n: int) -> Optional[ListNode]:
    dummy = ListNode(0)
    dummy.next = head
    first = dummy
    second = dummy

    for _ in range(n + 1):
        first = first.next

    while first is not None:
        first = first.next
        second = second.next

    second.next = second.next.next

    return dummy.next
```

### 2. TypeScript

**Runtime:** `51 ms` faster than `93.53%` submissions  
**Memory usage:** `51.5 MB` less than `71.28%` submissions  

``` typescript
function removeNthFromEnd(head: ListNode | null, n: number): ListNode | null {
    let first: ListNode | null = head;
    let second: ListNode | null = head;

    for (let i = 0; i < n; i++) {
        if (first === null) {
            return head;
        }
        first = first.next;
    }

    if (first === null) {
        return head !== null ? head.next : null;
    }

    while (first.next !== null) {
        first = first.next;
        second = second!.next;
    }

    if (second !== null && second.next !== null) {
        second.next = second.next.next;
    }

    return head;
}
```

### 3. GO

**Runtime:** `4 ms` faster than `97.42%` submissions  
**Memory usage:** `2.8 MB` less than `80.88%` submissions  

``` go
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Val: 0}
	dummy.Next = head
	first := dummy
	second := dummy

	for i := 0; i < n+1; i++ {
		first = first.Next
	}

	for first != nil {
		first = first.Next
		second = second.Next
	}

	second.Next = second.Next.Next

	return dummy.Next
}
```