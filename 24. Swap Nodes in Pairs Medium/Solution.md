### This approach has a time complexity of `O(n)`

### 1. Python

**Runtime:** `23 ms` faster than `98.78%` submissions  
**Memory usage:** `16.5 MB` less than `89.89%` submissions  

``` python
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

def swapPairs(head: ListNode) -> ListNode:
    dummy = ListNode(0)
    dummy.next = head
    prev = dummy

    while prev.next and prev.next.next:
        first, second = prev.next, prev.next.next
        prev.next, first.next, second.next = second, second.next, first

        prev = first

    return dummy.next
```

### 2. TypeScript

**Runtime:** `55 ms` faster than `68.83%` submissions  
**Memory usage:** `51.3 MB` less than `72.42%` submissions  

``` typescript
class ListNode {
    val: number;
    next: ListNode | null;

    constructor(val: number = 0, next: ListNode | null = null) {
        this.val = val;
        this.next = next;
    }
}

function swapPairs(head: ListNode): ListNode {
    let dummy: ListNode = new ListNode(0);
    dummy.next = head;
    let prev = dummy;

    while (prev.next !== null && prev.next?.next !== null) {
        let [first, second] = [prev.next, prev.next.next];
        [prev.next, first.next, second.next] = [second, second.next, first];

        prev = first;
    }

    return dummy.next;
}
```

### 3. GO

**Runtime:** `1 ms` faster than `82.58%` submissions  
**Memory usage:** `2.4 MB` less than `80.12%` submissions  

``` go
type ListNode struct {
	Val  int
	Next *ListNode
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
```